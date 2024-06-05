package system

import (
	"github.com/jdbann/asteroids/component"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type Friction struct {
	filter *generic.Filter2[component.Friction, component.Velocity]
}

func (s *Friction) Finalize(w *ecs.World) {
}

func (s *Friction) Initialize(w *ecs.World) {
	s.filter = generic.NewFilter2[component.Friction, component.Velocity]()
}

func (s *Friction) Update(w *ecs.World) {
	query := s.filter.Query(w)
	for query.Next() {
		friction, velocity := query.Get()
		velocity.X *= (1 - float32(*friction))
		velocity.Y *= (1 - float32(*friction))
	}
}

var _ model.System = (*Friction)(nil)
