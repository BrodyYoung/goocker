package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"goocker/container"
	"goocker/images"
)

//docker run -p 3306:3306  --restart=always  --name mysql -v /root/mysql/conf:/etc/mysql/conf.d -v /root/mysql/logs:/logs -v /root/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=Coolshark1588 -d mysql:5.7

var runCommand = cli.Command{
	Name: "运行容器Run Container",
	Flags: []cli.Flag{
		//cli.BoolFlag{} ：有该参数则为true，没有为false。例如 -it   -d
		//cli.StringFlag{}  要获取参数后面紧跟的值。 例如 --name redis-client
		//cli.StringSliceFlag{}  获取参数后面的多个值  例如：--p 8000:80  8080:8080 8081:8081
		cli.BoolFlag{
			Name:  "d",
			Usage: "守护式后台运行",
		},
		cli.BoolFlag{
			Name:  "it",
			Usage: "交互式终端运行interface terminal",
		},
		cli.StringSliceFlag{
			Name:  "p",
			Usage: "端口映射",
		},
		cli.StringFlag{
			Name:  "name",
			Usage: "给容器起别名",
		},
		cli.StringFlag{
			Name:  "net",
			Usage: "指定网络模式",
		},
		cli.StringFlag{
			Name:  "e",
			Usage: "指定容器运行环境env",
		},
		cli.StringFlag{
			Name:  "v",
			Usage: "容器卷挂载",
		},
	},
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return fmt.Errorf("args is nil")
		}
		d := ctx.Bool("d")
		it := ctx.Bool("it")
		if d && it {
			return fmt.Errorf("不能同时有it和d参数")

		}
		p := ctx.StringSlice("p")
		name := ctx.String("name")
		imagesName := ctx.Args().Get(0)
		v := ctx.String("v")
		e := ctx.String("e")

		var cmdArray []string
		for _, args := range ctx.Args().Tail() {
			cmdArray = append(cmdArray, args)
		}
		Run(cmdArray, p, name, d, it, imagesName, v, e)
		return nil
	},
}

var commitCommand = cli.Command{
	Name:  "commit",
	Usage: "提交",


	Action: func(ctx *cli.Context) error {
		imageName := ctx.Args().Get(0)
		container.CommitContainer(imageName)
		return nil
	},
}

var initCommand = cli.Command{

	Name:  "init",
	Usage: "初始化",
	Action: func(ctx *cli.Context) error {
		logrus.Printf("init  started")
		return container.RunInitContainer(ctx)
	},
}

var stopCommand = cli.Command{

	Name:  "stop",
	Usage: "停止容器",
	Action: func(ctx *cli.Context) error {
		containerName := ctx.Args().Get(0)
		container.StopContainer(containerName)
		return nil
	},
}

var rmCommand = cli.Command{

	Name:  "rm",
	Usage: "删除容器",
	Action: func(ctx *cli.Context) error {
		imagesName := ctx.Args().Get(0)
		container.RmContainer(imagesName)
		return nil
	},
}

var execCommand = cli.Command{

	Name:  "exec",
	Usage: "在容器内执行命令",
	Action: func(ctx *cli.Context) error {
		imageName := ctx.Args().Get(0)
		container.ExecContainer(imageName)
		return nil
	},
}

var pullCommand = cli.Command{
	Name:  "pull",
	Usage: "拉取镜像",
	Action: func(ctx *cli.Context) error {
		imageName := ctx.Args().Get(0)
		images.PullImage(imageName)
		return nil
	},
}
var pushCommand = cli.Command{
	Name:  "push",
	Usage: "推送容器",
	Action: func(ctx cli.Context) error {
		imageName := ctx.Args().Get(0)
		images.PushImage(imageName)
		return nil
	},
}

var psCommand = cli.Command{
	Name:  "ps",
	Usage: "查看容器列表",
	Action: func(ctx *cli.Context) error {
		a := ctx.String("a")
		container.PsContainer(a)
		return nil
	},
}

var logCommand = cli.Command{
	Name:  "log",
	Usage: "容器日志",
	Action: func(ctx *cli.Context) {
		containerName := ctx.Args().Get(0)

		container.GetContainerLog(containerName)

	},
}
