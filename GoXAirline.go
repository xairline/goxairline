package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xairline/goplane/extra"
	"github.com/xairline/goplane/extra/logging"
	"github.com/xairline/goplane/xplm/menus"
	"github.com/xairline/goplane/xplm/processing"
)

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

	menuId := menus.FindPluginsMenu()
	menuContainerId := menus.AppendMenuItem(menuId, "TestPlugin Menu", nil, false)
	myMenuId = menus.CreateMenu("TestPlugin Menu", menuId, menuContainerId, menuHandler, nil)
	menus.AppendMenuItem(myMenuId, "TestPlugin Menu sub 1", "TestPlugin Menu sub 1", false)
	menus.AppendMenuItem(myMenuId, "TestPlugin Menu sub 2", "TestPlugin Menu sub 2", false)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		logging.Info("Plugin started")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	go r.Run(":8080")

	processing.RegisterFlightLoopCallback(flightLoop, 0.03, nil)
}

func menuHandler(menuRef, itemRef interface{}) {
	logging.Infof("clicked: %+v", itemRef)
}

func flightLoop(elapsedSinceLastCall, elapsedTimeSinceLastFlightLoop float32, counter int, ref interface{}) float32 {
	logging.Infof("Flight loop:%f", elapsedSinceLastCall)
	return 0.03
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
