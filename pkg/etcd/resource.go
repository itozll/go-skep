package etcd

type Resource struct {
	Before []string `json:"before,omitempty" yaml:"before"`
	After  []string `json:"after,omitempty" yaml:"after"`

	Provider string                 `json:"provider,omitempty" yaml:"provider"`
	Binder   map[string]interface{} `json:"binder,omitempty" yaml:"binder"`

	Actions []*Action `json:"actions,omitempty" yaml:"actions"`
}

type Action struct {
	Before []string `json:"before,omitempty" yaml:"before"`
	After  []string `json:"after,omitempty" yaml:"after"`

	Provider string                 `json:"provider,omitempty" yaml:"provider"`
	Binder   map[string]interface{} `json:"binder,omitempty" yaml:"binder"`

	To       string   `json:"to,omitempty" yaml:"to"`
	Template []string `json:"template,omitempty" yaml:"template"`
	Copy     []string `json:"copy,omitempty" yaml:"copy"`
}
