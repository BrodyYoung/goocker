package subsystem

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type CpuSubsystem struct {
	apply bool
}

func (*CpuSubsystem) Name() string {
	return "cpu"
}

func (ss *CpuSubsystem) Set(cgroupPath string, res *ResourceConfig) error {
	info, err := GetCgroupPath(ss.Name(), cgroupPath, true)
	if err != nil {
		logrus.Error("error")
		return err
	}
	if res.CpuShare != "" {
		ss.apply = true
		err = ioutil.WriteFile(path.Join(info, "cpushare.txt"), []byte(res.CpuShare), os.ModePerm)
		if err != nil {
			logrus.Error("error")
			return err
		}
	}
	return nil
}

func (ss *CpuSubsystem) Apply(cgroupPath string, pid int) error {
	if ss.apply {
		info, err := GetCgroupPath(ss.Name(), cgroupPath, true)
		if err != nil {
			logrus.Error("error")
			return err
		}
		err = ioutil.WriteFile(path.Join(info, ""), []byte(strconv.Itoa(pid)), os.ModePerm)
		if err != nil {
			logrus.Error("error")
			return err
		}
	}
	return nil

}

func (ss *CpuSubsystem) Remove(cgroupPath string) error {
	info, err := GetCgroupPath(ss.Name(), cgroupPath, true)
	if err != nil {
		logrus.Error("error")
	}
	return os.RemoveAll(info)
}
