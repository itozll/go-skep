package initd

var (
	DefaultGroup     = ""
	DefaultGoVersion = "1.17"
	DefaultCommand   = "root"
	SkipGit          = false

	Data     string
	File     string
	FileType string

	DefaultTemplateName = "base"
	ConfigFile          = ".skeprc"
	Version             = "v0.1.0"
)
