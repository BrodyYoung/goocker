package container

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"goocker/common"
	"os"
	"path"
)

func GetContainerLog(containerName string) {

	path := path.Join(common.DefaultContainerPath, containerName, common.ContainerLogFileName)

	byteSlice, err := os.ReadFile(path)
	if err != nil {
		logrus.Print(err)
	}
	_, _ = fmt.Fprint(os.Stdout, byteSlice)

}
