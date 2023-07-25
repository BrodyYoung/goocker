package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "goocker"
	app.Usage = "docker-like container"
	app.Commands = []cli.Command{
		runCommand,
		stopCommand,
		commitCommand,
		//startCommand,
		//restartCommand,
		//buildCommand,
		//tagCommand,
		//imagesCommand,
		//rmiCommand,
		//searchCommand,

		pullCommand,
		pushCommand,

		execCommand,
		psCommand,
		rmCommand,
		initCommand,
		logCommand,
	}
	app.Before = func(ctx *cli.Context) error {
		logrus.SetOutput(os.Stdout)
		logrus.SetFormatter(&logrus.JSONFormatter{})
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
