package subsystem

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type CpusetSubsystem struct {
	apply bool
}

func (ss *CpusetSubsystem) Name() string {

	return "cpuset"
}

func (ss *CpusetSubsystem) Set(cgroupPath string, res *ResourceConfig) error {
	info, err := GetCgroupPath(ss.Name(), cgroupPath, true)
	if err != nil {
		logrus.Error("error")
	}
	if res.CpuSet != "" {
		ss.apply = true

		err = ioutil.WriteFile(path.Join(info, "cpuset.txt"), []byte(res.CpuSet), os.ModePerm)
		if err != nil {
			logrus.Error("error")
			return err
		}
	}
	return nil
}

func (ss *CpusetSubsystem) Apply(cgroupPath string, pid int) error {
	if ss.apply {

		//第三个参数是做什么的
		info, err := GetCgroupPath(ss.Name(), cgroupPath, true)

		if err != nil {
			logrus.Error("error")
			return err
		}
		err = ioutil.WriteFile(path.Join(info, "tasks"), []byte(strconv.Itoa(pid)), os.ModePerm)
		if err != nil {
			logrus.Error("error")
			return err
		}
	}
	return nil
}

func (ss *CpusetSubsystem) Remove(cgroupPath string) error {

	info, err := GetCgroupPath(ss.Name(), cgroupPath, true)
	if err != nil {
		logrus.Error("error")
	}
	return os.RemoveAll(info)
}
