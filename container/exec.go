package container

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"goocker/common"
	"os"
	"os/exec"
	"strings"
)

func Exec(cmdSlice []string, containerName string) {
	info, err := getContainerInfo(containerName)
	if err != nil {
		logrus.Errorf("err")
	}
	cmd := exec.Command("/etc/exe", "exec")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	os.Setenv(common.ExecEnvPid, info.Pid)
	os.Setenv(common.ExecEnvCmd, strings.Join(cmdSlice, " "))

	envs := getEnvsByPid(info.Pid)
	cmd.Env = append(os.Environ(), envs...)

	if err := cmd.Run(); err != nil {
		logrus.Error(err)
	}
}

func getEnvsByPid(pid string) []string {

	path := fmt.Sprintf("/proc/%s/environ", pid)

	byteSlice, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	return strings.Split(string(byteSlice), "\u0000")
}
