package command

import (
	"github.com/itozll/go-skep/pkg/runtime/initd"
	"github.com/itozll/go-skep/pkg/runtime/rtstatus"
	"github.com/itozll/go-skep/pkg/tmpl"
)

type ExecHandler interface {
	Exec(parent WorkerHandler) error

	WorkerHandler
}

type WorkerHandler interface {
	Dir() string
	Provider() tmpl.Provider
	MapBinder() map[string]interface{}
}

type Worker struct {
	Before func() error
	After  func() error

	Path string
	P    tmpl.Provider

	Binder map[string]interface{}

	Handlers []ExecHandler
}

func (w *Worker) Dir() string                       { return w.Path }
func (w *Worker) Provider() tmpl.Provider           { return w.P }
func (w *Worker) MapBinder() map[string]interface{} { return w.Binder }

func (w *Worker) Exec() (err error) {
	if w.Binder == nil {
		w.Binder = initd.MapBinder()
	} else {
		for key, val := range initd.MapBinder() {
			if _, ok := w.Binder[key]; !ok {
				w.Binder[key] = val
			}
		}
	}

	if w.Before != nil {
		if err = w.Before(); err != nil {
			return
		}
	}

	for _, action := range w.Handlers {
		rtstatus.ExitIfError(action.Exec(w))
	}

	if w.After != nil {
		return w.After()
	}

	return nil
}
