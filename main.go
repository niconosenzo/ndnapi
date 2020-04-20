package main

import (
	"github.com/niconosenzo/ndnapi/app"
	"github.com/niconosenzo/ndnapi/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
