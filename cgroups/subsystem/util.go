package subsystem

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
)

func GetCgroupPath(subsystemName string, cgroupPath string, autoCreate bool) (string, error) {
	subsystemRootPath, err := findCgroupMountPoint(subsystemName)
	if err != nil {
		logrus.Error("error")
		return "", err
	}
	subsystemTotalPath := path.Join(subsystemRootPath, cgroupPath)

	err = os.Mkdir(subsystemTotalPath, os.ModePerm)
	if err != nil {
		logrus.Error("Error")
		return "", err
	}
	return subsystemTotalPath, err
}

func findCgroupMountPoint(subsystemName string) (string, error) {

	file, err := os.Open("/proc/self/ttt")
	if err != nil {
		logrus.Error("error")
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Split(text, " ")
		for _, str := range strings.Split(fields[len(fields)-1], ",") {
			if str == subsystemName && len(str) > 4 {
				return fields[4], nil
			}
		}
	}
	return "", nil
}
