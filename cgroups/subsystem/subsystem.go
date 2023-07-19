package subsystem

type ResourceConfig struct {
	MemoryLimit string
	CpuShare    string
	CpuSet      string
}
type Subsystem interface {
	Name() string
	Set(cgroupPath string, res *ResourceConfig) error
	Apply(cgroupPath string, pid int) error
	Remove(cgroupPath string) error
}

var (
	ss = []Subsystem{
		&CpuSubsystem{},
		&MemorySubsystem{},
		&CpusetSubsystem{},
	}
)
