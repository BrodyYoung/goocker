package main

import "goocker/container"

func Run(cmdArray, envs, portMapping []string, res *subsystem.ResourceConfig imageName, containerName, volume, net string, it bool) {
	containerId := container.GenContainerId(10)
	if containerName == "" {
		containerName = containerId
	}

	info := &container.ContainerInfo{
		Id:          containerId,
		Name:        containerName,
		Volume:      volume,
		Env:         envs,
		PortMapping: portMapping,
	}

	if it {
		parent := GetParentProcess()
		parent.Process.Pid

		parent.Start()

	}

}
