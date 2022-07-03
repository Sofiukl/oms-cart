package main

import "github.com/sofiukl/oms-cart/core"

func main() {
	app := core.App{}
	app.Initialize()
	app.Run(":" + app.Config.ServerPort)
}
