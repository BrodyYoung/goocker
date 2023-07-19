package subsystem

type ResourceConfig struct {
	MemoryLimit string
	CpuShare    string
	CpuSet      string
}
type Subsystem interface {
	Name() string
	Set(c *cgroup.Cgroup) error
	Apply(c *cgroup.Cgroup) error
	Remove(c *cgroup.Cgroup) error
}

var (
	ss = []Subsystem{
		&CpuSubsystem{},
		&MemorySubsystem{},
		&CpusetSubsystem{},
	}
)
