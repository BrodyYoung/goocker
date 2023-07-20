package subsystem

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type MemorySubsystem struct {
	apply bool
}

func (ss *MemorySubsystem) Name() string {
	return "memory"
}

func (ss *MemorySubsystem) Set(cgroupPath string, res *ResourceConfig) error {

	info, err := GetCgroupPath(ss.Name(), cgroupPath, true)
	if err != nil {
		logrus.Error("error")
		return err
	}
	if res.MemoryLimit != "" {
		ss.apply = true
		err := ioutil.WriteFile(path.Join(info, "memory.txt"), []byte(res.MemoryLimit), os.ModePerm)

		if err != nil {
			logrus.Error("error")
			return err
		}
	}
	return nil
}

func (ss *MemorySubsystem) Apply(cgroupPath string, pid int) error {
	if ss.apply {
		info, err := GetCgroupPath(ss.Name(), cgroupPath, true)
		if err != nil {
			logrus.Error("error")
		}
		err = ioutil.WriteFile(path.Join(info, ""), []byte(strconv.Itoa(pid)), os.ModePerm)
		if err != nil {
			logrus.Error("error")
			return err
		}
	}
	return nil
}

func (ss *MemorySubsystem) Remove(cgroupPath string) error {

	info, err := GetCgroupPath(ss.Name(), cgroupPath, true)
	if err != nil {
		logrus.Error("error")
		return err
	}
	return os.RemoveAll(info)
}
