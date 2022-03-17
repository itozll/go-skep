package entity

import "errors"

type Model struct {
	Resource string `json:",inline,omitempty" yaml:",inline"`

	Flag string `json:"flag,omitempty" yaml:"flag"`
}

func (eni *Model) Validate() error {
	if eni.Flag == "" {
		return errors.New(`"flag" must not be empty`)
	}

	return nil
}
