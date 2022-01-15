package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nakabonne/tstorage"
	"github.com/xairline/goplane/extra"
	"github.com/xairline/goplane/extra/logging"
	"github.com/xairline/goplane/xplm/processing"
)

const POLL_FEQ = 20

var Plugin *extra.XPlanePlugin
var Storage tstorage.Storage
var tracking bool

func main() {
}

func init() {
	Plugin = extra.NewPlugin("X Airline", "com.github.xairline.goxairline", "Native plugin for x airline")
	Plugin.SetPluginStateCallback(onPluginStateChanged)
	logging.MinLevel = logging.Info_Level

	// setup storage
	var storageErr error
	storageDuration, _ := time.ParseDuration("1h")
	Storage, storageErr = tstorage.NewStorage(
		tstorage.WithDataPath(os.Getenv("HOME")+"/.xiarline/data"),
		tstorage.WithPartitionDuration(storageDuration),
		tstorage.WithTimestampPrecision(tstorage.Milliseconds),
	)
	if storageErr != nil {
		logging.Errorf("Failed initialize TS storage: %+v", storageErr)
	}
	logging.Infof("Initialized TS storage: %s", os.Getenv("HOME")+"/.xiarline/data")

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

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		logging.Info("ping")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	go r.Run(":8080")

	processing.RegisterFlightLoopCallback(flightLoop, 1/POLL_FEQ, nil)
}

func flightLoop(elapsedSinceLastCall, elapsedTimeSinceLastFlightLoop float32, counter int, ref interface{}) float32 {
	logging.Debugf("Flight loop:%f", elapsedSinceLastCall)
	return 1 / POLL_FEQ
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
