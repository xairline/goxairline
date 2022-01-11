package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xairline/goplane/extra"
	"github.com/xairline/goplane/extra/logging"
	"github.com/xairline/goplane/xplm/menus"
	"github.com/xairline/goplane/xplm/processing"
)

const POLL_FEQ = 20

func main() {
}

var plugin *extra.XPlanePlugin

func init() {
	plugin = extra.NewPlugin("X Airline", "com.github.xairline.goxairline", "Native plugin for x airline")
	plugin.SetPluginStateCallback(onPluginStateChanged)
	logging.MinLevel = logging.Info_Level
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

var myMenuId menus.MenuID

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
	menus.DestroyMenu(myMenuId)
	logging.Info("Plugin stopped")
}

func onPluginEnable() {
	logging.Info("Plugin enabled")
}

func onPluginDisable() {
	logging.Info("Plugin disabled")
}
