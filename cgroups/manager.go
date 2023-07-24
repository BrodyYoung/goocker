package cgroups

import (
	"github.com/sirupsen/logrus"
	"goocker/cgroups/subsystem"
)

type CgroupManager struct {
	Path string
}

func NewCgroupManager(path string) *CgroupManager {
	return &CgroupManager{Path: path}
}

func (c CgroupManager) Set(res *subsystem.ResourceConfig) {
	for _, ss := range subsystem.Subsystems {
		err := ss.Set(c.Path, res)
		if err != nil {
			logrus.Error(err)
		}
	}
}

func (c CgroupManager) Apply(pid int) {
	for _, ss := range subsystem.Subsystems {
		err := ss.Apply(c.Path, pid)
		if err != nil {
			logrus.Error(err)
		}
	}
}

func (c CgroupManager) Destroy() {
	for _, ss := range subsystem.Subsystems {
		err := ss.Remove(c.Path)
		if err != nil {
			logrus.Error(err)
		}
	}
}
