package etcd

var NewEtc = []byte(`
actions:
- parse:
  - README:README.md
  - Makefile
  - main_go:main.go
  - gomod:go.mod
  - rootcmd_go:app/cmd/root.go
  - subcmd_go:app/cmd/server.go
  - context_go:app/internal/runtime/rtinfo/context.go
  copy:
  - gitignore:.gitignore
  - generator:scripts/generator.sh
  - rtinfo_go:app/internal/runtime/rtinfo/rtinfo.go
`)

// var NewEtc = &generator.New{
// 	Resource: generator.Resource{
// 		Actions: []*generator.Action{
// 			{
// 				Parse: []string{
// 					"Makefile",
// 					"README:README.md",
// 					"main_go:main.go",
// 					"gomod:go.mod",
// 					"rootcmd_go:app/cmd/root.go",
// 					"subcmd_go:app/cmd/server.go",
// 					"context_go:app/internal/runtime/rtinfo/context.go",
// 				},
// 				Copy: []string{
// 					"gitignore:.gitignore",
// 					"generator:scripts/generator.sh",
// 					"rtinfo_go:app/internal/runtime/rtinfo/rtinfo.go",
// 				},
// 			},
// 		},
// 	},
// }
