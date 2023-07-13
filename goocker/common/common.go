package common

//系统路径相关常量
const (
	RootPath   = "/root"
	MntPath    = "/root/mnt"
	WriteLayer = "WriteLayer"
)

//容器运行状态常量
const (
	Running = "Running"
	Stop    = "Stopped"
	Exit    = "Exited"
)

//容器文件路径相关常量
const (
	DefaultContainerPath  = "/var/run/goocker/"
	ContainerInfoFileName = "config.json"
	ContainerLogFileName  = "container.log"
)

//执行环境相关常量
const (
	ExecEnvPid = "goocker_pid"
	ExecEnvCmd = "goocker_cmd"
)
