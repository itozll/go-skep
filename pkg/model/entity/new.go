package entity

import "github.com/itozll/go-skep/pkg/runtime/rtinfo"

type New struct {
	SkipGit bool   `json:"skip_git,omitempty" yaml:"skip_git"`
	Group   string `json:"group,omitempty" yaml:"group"`
	Go      string `json:"go_version,omitempty" yaml:"go_version"`

	Resource `json:",inline" yaml:",inline"`
}

func (eni *New) Parse() {
	if !rtinfo.FlagSkipGit.Changed && eni.SkipGit {
		rtinfo.SkipGit = true
	}

	if !rtinfo.FlagGroup.Changed && eni.Group != "" {
		rtinfo.Group = eni.Group
	}

	if !rtinfo.FlagGoVersion.Changed && eni.Go != "" {
		rtinfo.Go = eni.Go
	}
}
