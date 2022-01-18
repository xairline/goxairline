package main

import (
	"os"
	"path/filepath"
	"time"
	"xairline/goxairline/internal/xplane/config"
	datarefext "xairline/goxairline/internal/xplane/datarefExt"
	"xairline/goxairline/internal/xplane/shared"

	"github.com/gin-gonic/gin"
	"github.com/nakabonne/tstorage"
	"github.com/xairline/goplane/extra"
	"github.com/xairline/goplane/extra/logging"
	"github.com/xairline/goplane/xplm/dataAccess"
	"github.com/xairline/goplane/xplm/plugins"
	"github.com/xairline/goplane/xplm/processing"
	"github.com/xairline/goplane/xplm/utilities"
)

const PollFeq = 20

var Plugin *extra.XPlanePlugin
var Storage tstorage.Storage
var tracking bool
var datarefList []datarefext.DataRefExt

func main() {
}

func init() {
	Plugin = extra.NewPlugin("X Airline", "com.github.xairline.goxairline", "Native plugin for x airline")
	Plugin.SetPluginStateCallback(onPluginStateChanged)
	plugins.EnableFeature("XPLM_USE_NATIVE_PATHS", true)
	logging.MinLevel = logging.Info_Level

	// setup storage
	var storageErr error
	storageDuration, _ := time.ParseDuration("1h")
	Storage, storageErr = tstorage.NewStorage(
		tstorage.WithDataPath(os.Getenv("HOME")+"/.xairline/data"),
		tstorage.WithPartitionDuration(storageDuration),
		tstorage.WithTimestampPrecision(tstorage.Milliseconds),
	)
	if storageErr != nil {
		logging.Errorf("Failed initialize TS storage: %+v", storageErr)
	}
	logging.Infof("Initialized TS storage: %s", os.Getenv("HOME")+"/.xairline/data")

	tracking = false
	logging.Infof("Set tracking to: %v", tracking)
}

func onPluginStateChanged(state extra.PluginState, plugin *extra.XPlanePlugin) {
	switch state {
	case extra.PluginStart:
		onPluginStart()
	case extra.PluginStop:
		onPluginStop()
	case extra.PluginEnable:
		onPluginEnable()
	case extra.PluginDisable:
		onPluginDisable()
	}
}

func onPluginStart() {
	logging.Info("Plugin started")

	// get plugin path
	systemPath := utilities.GetSystemPath()
	pluginPath := filepath.Join(systemPath, "Resources", "plugins", "xairline")
	logging.Infof("Plugin path: %s", pluginPath)

	logger := shared.Logger{
		Infof:  logging.Infof,
		Errorf: logging.Errorf,
	}

	// get config from file
	config := config.NewConfig(filepath.Join(pluginPath, "config.yaml"), &logger)
	// create dataref listeners
	for _, dataref := range config.DatarefConfig {
		datarefList = append(datarefList, *datarefext.NewDataRefExt(
			dataref.Name,
			dataref.DatarefStr,
			int8(dataref.Precision),
			dataref.IsBytesArray,
			dataAccess.FindDataRef,
			dataAccess.GetDataRefTypes,
			&logger,
		))
	}

	// running data processing pipeline in background
	go xplane.GlobalDatarefStore.ProcessFromGlobalDatarefStore(&logger)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		logging.Info("ping")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	go r.Run(":8080")

	processing.RegisterFlightLoopCallback(flightLoop, 1/PollFeq, nil)
}

func flightLoop(elapsedSinceLastCall, elapsedTimeSinceLastFlightLoop float32, counter int, ref interface{}) float32 {
	logging.Debugf("Flight loop:%f", elapsedSinceLastCall)
	return 1 / PollFeq
	datarefElement := map[string]interface{}{}
	for _, dataref := range datarefList {
		datarefElement[dataref.GetName()] = dataref.GetCurrentValue()
	}
	xplane.GlobalDatarefStore = append(xplane.GlobalDatarefStore, datarefElement)
	if len(xplane.GlobalDatarefStore)%1000 == 0 {
		logging.Infof("%v", xplane.GlobalDatarefStore[len(xplane.GlobalDatarefStore)-1])
	}
	return -1 //every frame
}

func onPluginStop() {
	defer Storage.Close()
	logging.Info("Plugin stopped")
}

func onPluginEnable() {
	logging.Info("Plugin enabled")
}

func onPluginDisable() {
	logging.Info("Plugin disabled")
}
