package bet

import (
	"testing"
)

type Specification[C any] struct {
	t     *testing.T
	setup func(t *testing.T) C
	specs []specRecord[C]
}

func New[C any](t *testing.T) Specification[C] { return Specification[C]{t: t} }

type specRecord[C any] struct {
	name string
	spec func(t *testing.T, configuration C)
}

func (set *Specification[C]) Run() {
	for i := range set.specs {
		set.t.Run(set.specs[i].name, func(t *testing.T) {
			var configuration C
			if set.setup != nil {
				configuration = set.setup(t)
			}
			set.specs[i].spec(set.t, configuration)
		})
	}
}

func (set *Specification[C]) Setup(initConfig func(t *testing.T) C) { set.setup = initConfig }

func (set *Specification[C]) Spec(name string, spec func(t *testing.T, config C)) {
	set.specs = append(set.specs, specRecord[C]{
		name: name,
		spec: spec,
	})
}
