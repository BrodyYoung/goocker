package container

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"goocker/common"
	"math/rand"
	"os"
	"path"
	"time"
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
	byteSlice, err := os.ReadFile(path)
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

func GenContainerId(n int) string {
	letterChar := "0123456789"

	var res = make([]byte, n)
	rand.Seed(time.Now().UnixNano()) //???????做什么的

	for i, _ := range res {
		res[i] = letterChar[rand.Intn(n)]
	}

	return string(res)
}
