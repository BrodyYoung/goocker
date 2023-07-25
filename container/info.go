package container

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"goocker/common"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type ContainerInfo struct {
	Pid         string
	Id          string
	Name        string
	Command     string
	Volume      string
	Env         []string
	PortMapping []string
	CreateTime  string
	Status      string
}

func RecordContainerInfo(pid int, cmdArr []string, id, containerName string) error {
	info := &ContainerInfo{
		Pid:        strconv.Itoa(pid),
		Id:         id,
		Name:       containerName,
		Command:    strings.Join(cmdArr, " "),
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		Status:     common.Running,
	}

	path := path.Join(common.DefaultContainerPath, containerName)

	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	file, err := os.Create(path)
	if err != nil {
		logrus.Error(err)
	}

	byteSlice, err := json.Marshal(info)
	_, err = file.Write(byteSlice)
	if err != nil {
		logrus.Error(err)
	}
	return nil
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

func DeleteContainerInfo(containerName string) {

	path := path.Join(common.DefaultContainerPath, containerName)

	if err := os.RemoveAll(path); err != nil {
		logrus.Error(err)
	}
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
