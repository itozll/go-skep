package model

import (
	"github.com/itozll/go-skep/pkg/generator"
	"github.com/itozll/go-skep/pkg/model/entity"
	"github.com/itozll/go-skep/pkg/process"
	"github.com/itozll/go-skep/pkg/tmpl"
)

func genCommand(res *entity.Resource) *generator.Command {
	c := &generator.Command{
		Binder: res.Binder,
		P:      tmpl.GetTemplateProvider(res.Provider),
		Before: process.Command(res.Before),
		After:  process.Command(res.After),
	}

	for _, action := range res.Actions {
		c.Actions = append(c.Actions, &generator.Action{
			Before:   process.Command(action.Before),
			After:    process.Command(action.After),
			Binder:   action.Binder,
			P:        tmpl.GetTemplateProvider(action.Provider),
			To:       action.To,
			Template: action.Template,
			Copy:     action.Copy,
		})
	}

	return c
}
