package rtinfo

const (
	Version = "v0.1.0"
)

var (
	DefaultGoVersion = "1.17"
	DefaultGroup     = ""
)

var (
	GoVersion string

	// Workspace = [Group/]<Project>
	Workspace string

	// github.com/<Group>/<Project>
	Github  string
	Group   string
	Project string
	SkipGit bool

	Data string
)
