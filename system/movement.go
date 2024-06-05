package system

import (
	"github.com/jdbann/asteroids/component"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type Movement struct {
	filter *generic.Filter2[component.Forces, component.Position]
}

func (s *Movement) Finalize(w *ecs.World) {
}

func (s *Movement) Initialize(w *ecs.World) {
	s.filter = generic.NewFilter2[component.Forces, component.Position]()
}

func (s *Movement) Update(w *ecs.World) {
	query := s.filter.Query(w)
	for query.Next() {
		forces, position := query.Get()
		position.Coords = position.Coords.Add(forces.Velocity)
		forces.Velocity = forces.Velocity.Scale(1 - forces.Friction)
	}
}

var _ model.System = (*Movement)(nil)
