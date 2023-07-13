package container

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"goocker/common"
	"io/ioutil"
	"os"
	"path"
)

func GetContainerLog(containerName string) {

	path := path.Join(common.DefaultContainerPath, containerName, common.ContainerLogFileName)
	file, err := os.Open(path)
	if err != nil {
		logrus.Print(err)
	}

	byteSlice, err := ioutil.ReadAll(file)
	if err != nil {
		logrus.Print(err)
	}
	_, _ = fmt.Fprint(os.Stdout, byteSlice)

}
