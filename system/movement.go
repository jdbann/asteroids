package system

import (
	"github.com/jdbann/asteroids/component"
	"github.com/jdbann/asteroids/util/geo"
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
		position.Coords = position.Coords.Add(geo.Vec2(*velocity))
	}
}

var _ model.System = (*Movement)(nil)
