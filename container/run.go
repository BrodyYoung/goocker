package container

import (
	"github.com/sirupsen/logrus"
	"goocker/cgroups"
	"goocker/cgroups/subsystem"
	"goocker/network"
	"os"
	"strings"
)

func Run(cmdArray, envs, portMapping []string, res *subsystem.ResourceConfig, imageName, containerName, volume, net string, it bool) {
	containerId := GenContainerId(10)
	if containerName == "" {
		containerName = containerId
	}

	parent, writePipe := NewParentProcess(volume, envs, imageName, containerName, it)
	if parent == nil {
		logrus.Error("error")
		return
	}
	if err := parent.Start(); err != nil {
		logrus.Error(err)
		return
	}

	err := RecordContainerInfo(parent.Process.Pid, cmdArray, containerId, containerName)
	if err != nil {
		logrus.Error(err)
		return
	}

	cgroupManager := cgroups.NewCgroupManager("goocker")
	//cgroupManager:=cgroups.CgroupManager{Path:"goocker"}
	cgroupManager.Set(res)
	cgroupManager.Apply(parent.Process.Pid)
	defer cgroupManager.Destroy()

	if net != "" {
		if err := network.Init(); err != nil {
			logrus.Error(err)
			return
		}
		info := &ContainerInfo{
			Id:          containerId,
			Name:        containerName,
			Pid:         parent.Process.Pid,
			Env:         envs,
			PortMapping: portMapping,
		}
		if err := network.Connect(info, net); err != nil {
			logrus.Error(err)
			return
		}
	}
	sendInitCommand(cmdArray, writePipe)

	if it {
		if err := parent.Wait(); err != nil {
			logrus.Error(err)
		}
		if err := DeleteContainerWorkSpace(info, volume); err != nil {
			logrus.Error(err)
		}
		DeleteContainerInfo(containerName)
	}

}

func sendInitCommand(array []string, writePipe *os.File) {
	cmd := strings.Join(array, " ")

	writePipe.WriteString(cmd)
	writePipe.Close()
}
