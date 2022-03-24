package etcd

import "github.com/itozll/go-skep/pkg/command/generator"

var NewEtc = &generator.New{
	Resource: generator.Resource{
		Actions: []*generator.Action{
			{
				Parse: []string{
					"README:README.md",
					"Makefile",
					"main_go:main.go",
					"gomod:go.mod",
				},
				Copy: []string{
					"gitignore:.gitignore",
					"generator:scripts/generator.sh",
				},
			},
			{
				Path: "app/cmd",
				Parse: []string{
					"rootcmd_go:root.go",
					"subcmd_go:server.go",
				},
			},
			{
				Path: "app/internal/runtime/rtinfo",
				Parse: []string{
					"context_go:context.go",
				},
				Copy: []string{
					"rtinfo_go:rtinfo.go",
				},
			},
		},
	},
}
