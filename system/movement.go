package system

import (
	"github.com/jdbann/asteroids/component"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type Movement struct {
	filter *generic.Filter2[component.Position, component.Velocity]
}

func (s *Movement) Finalize(w *ecs.World) {
}

func (s *Movement) Initialize(w *ecs.World) {
	s.filter = generic.NewFilter2[component.Position, component.Velocity]()
}

func (s *Movement) Update(w *ecs.World) {
	query := s.filter.Query(w)
	for query.Next() {
		position, velocity := query.Get()
		position.X = position.X + velocity.X
		position.Y = position.Y + velocity.Y
	}
}

var _ model.System = (*Movement)(nil)
