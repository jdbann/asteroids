package uisystem

import (
	"github.com/jdbann/asteroids/component"
	"github.com/jdbann/asteroids/resource"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type PlayerControls struct {
	filter *generic.Filter1[component.Velocity]

	keyBindingsRes generic.Resource[resource.KeyBindings]
}

func (s *PlayerControls) FinalizeUI(w *ecs.World) {
}

func (s *PlayerControls) InitializeUI(w *ecs.World) {
	s.filter = generic.NewFilter1[component.Velocity]().
		With(generic.T[component.PlayerControlled]())

	s.keyBindingsRes = generic.NewResource[resource.KeyBindings](w)
}

func (s *PlayerControls) PostUpdateUI(w *ecs.World) {
}

func (s *PlayerControls) UpdateUI(w *ecs.World) {
	keyBindings := s.keyBindingsRes.Get()

	query := s.filter.Query(w)
	for query.Next() {
		velocity := query.Get()

		if keyBindings.FireThrusters.IsDown() {
			velocity.Y += .2
		}
	}
}

var _ model.UISystem = (*PlayerControls)(nil)
