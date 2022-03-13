package rtinfo

const (
	Version = "v0.1.0"
)

var (
	DefaultGoVersion = "1.17"
	DefaultGroup     = ""
)

var (
	Go string

	// Workspace = [Group/]<Project>
	Workspace string

	// github.com/<Group>/<Project>
	Github  string
	Group   string
	Project string
	SkipGit bool

	Data     string
	File     string
	FileType string
)
