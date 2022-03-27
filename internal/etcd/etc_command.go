package etcd

var CommandEtc = []byte(`
actions:
- parse:
  - subcmd_go:app/cmd/{{- .file_name -}}.go
`)
