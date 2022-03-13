package etcd

type New struct {
	SkipGit   bool   `json:"skip_git,omitempty" yaml:"skip_git"`
	Group     string `json:"group,omitempty" yaml:"group"`
	GoVersion string `json:"go_version,omitempty" yaml:"go_version"`
	Workspace string `json:"workspace,omitempty" yaml:"workspace"`

	Resource `json:",inline" yaml:",inline"`
}

var NewEtc = &New{
	Resource: Resource{
		Actions: []*Action{
			{
				Template: []string{
					"README:.md",
					"Makefile",
					"main_go::main.go",
					"gomod::go.mod",
				},
				Copy: []string{
					"gitignore::.gitignore",
					"generator:.sh",
				},
			},
			{
				To: "app/cmd",
				Template: []string{
					"cmdroot::root.go",
					"cmdserver::server.go",
				},
			},
			{
				To: "app/internal/runtime/rtinfo",
				Template: []string{
					"context_go::context.go",
				},
				Copy: []string{
					"rtinfo_go::rtinfo.go",
				},
			},
		},
	},
}
