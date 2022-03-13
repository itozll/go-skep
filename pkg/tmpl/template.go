package tmpl

import "github.com/itozll/go-skep/pkg/runtime/rtstatus"

var _s = map[string]Provider{}

func GetTemplateProvider(name string) Provider {
	if name == "" {
		name = "base"
	}

	v, ok := _s[name]
	if !ok {
		rtstatus.Fatal("no such template named '" + name + "'")
	}

	return v
}

func AddTemplateProvider(name string, p Provider) {
	_s[name] = p
}
