package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
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
	func(ctx *cli.Context) error {
		cmd := cli.Command{
			Name:  "goocker",
			Usage: "goocker command",
		}

		return logrus.Fatal(cmd)
	}
}
