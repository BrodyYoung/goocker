package container

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"goocker/common"
	"io/ioutil"
	"path"
)

type ContainerInfo struct {
	Pid         string
	Id          string
	Name        string
	Command     string
	Volume      string
	Env         string
	PortMapping []string
	CreateTime  string
	Status      string
}

func getContainerInfo(containerName string) (*ContainerInfo, error) {
	path := path.Join(common.DefaultContainerPath, containerName, common.ContainerInfoFileName)
	byteSlice, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Errorf("error")
		return nil, err
	}
	info := &ContainerInfo{}
	err = json.Unmarshal(byteSlice, info)
	if err != nil {
		logrus.Errorf("error")
		return nil, err
	}
	return info, nil
}
