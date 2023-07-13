package container

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"goocker/common"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

func CommitContainer(imagePath, imageName string) error {
	if imagePath == "" {
		imagePath = common.RootPath
	}
	image := path.Join(imagePath, imageName, ".tar")
	_, err := exec.Command("tar", "-czf", image, "-C", common.MntPath, ".").CombinedOutput()
	if err != nil {
		logrus.Printf("error")
		return err
	}
	return nil

}

func RunInitContainer(ctx *cli.Context) error {

}
func StopContainer(containerName string) {

	info, err := getContainerInfo(containerName)
	if err != nil {
		logrus.Errorf("not container")
		return
	}

	if info.Pid != "" {
		pid, _ := strconv.Atoi(info.Pid)
		if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
			logrus.Errorf("error")
			return
		}

		info.Status = common.Stop
		info.Pid = ""
		byteSlice, _ := json.Marshal(info)

		path := path.Join(common.DefaultContainerPath, containerName, common.ContainerInfoFileName)
		err = ioutil.WriteFile(path, byteSlice, 0622)
		if err != nil {
			logrus.Error(err)
			return
		}
	}

}

func RmContainer(containerName string) {

	info, err := getContainerInfo(containerName)
	if err != nil {
		logrus.Error("error")
		return
	}
	if info.Status != common.Stop {
		logrus.Error("error")
		return
	}

	path := path.Join(common.DefaultContainerPath, containerName)
	err = os.RemoveAll(path)
	if err != nil {
		logrus.Error("error")
	}
}

func ExecContainer(imageName string) error {

}

func PsContainer(arg string) {

}
