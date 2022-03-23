package etcd

import "github.com/itozll/go-skep/pkg/command/generator"

var NewEtc = &generator.New{
	Resource: generator.Resource{
		Actions: []*generator.Action{
			{
				Parse: []string{
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
				Base: generator.Base{
					Path: "app/cmd",
				},
				Parse: []string{
					"cmdroot::root.go",
					"cmdserver::server.go",
				},
			},
			{
				Base: generator.Base{
					Path: "app/internal/runtime/rtinfo",
				},
				Parse: []string{
					"context_go::context.go",
				},
				Copy: []string{
					"rtinfo_go::rtinfo.go",
				},
			},
		},
	},
}
