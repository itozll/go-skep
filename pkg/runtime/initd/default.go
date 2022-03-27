package initd

var (
	DefaultGroup     = ""
	DefaultGoVersion = "1.17"
	DefaultCommand   = "root"
	SkipGit          = false

	Verbose bool

	DefaultTemplateName = "base"
	ConfigFile          = ".skeprc"
	ConfigPath          = ".skep.d"
	Version             = "v0.1.0"
)
