package main

import (
	"core_x/handler"
	"core_x/manager"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	bootstrap := NewBootstrap()

	bootstrap.Modules(
		manager.NewActionMangerFile(),
		manager.NewBlockMangerFile(),
		manager.NewTokenManagerFile(),
		manager.NewHandlerManager(),
		manager.NewHookManager(),
		manager.NewModuleManagerFile(),
		handler.NewCoreHandler(),
	)

	app := &cli.App{
		Name:     "Core-X",
		Usage:    "Core of project X",
		Flags:    bootstrap.Flags(),
		Commands: bootstrap.Command(),
		Action:   bootstrap.Action,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
