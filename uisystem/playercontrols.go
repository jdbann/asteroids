package uisystem

import (
	"github.com/jdbann/asteroids/component"
	"github.com/jdbann/asteroids/resource"
	"github.com/jdbann/asteroids/util/geo"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type PlayerControls struct {
	filter *generic.Filter3[component.Heading, component.Thrusters, component.Velocity]

	keyBindingsRes generic.Resource[resource.KeyBindings]
}

func (s *PlayerControls) FinalizeUI(w *ecs.World) {
}

func (s *PlayerControls) InitializeUI(w *ecs.World) {
	s.filter = generic.NewFilter3[component.Heading, component.Thrusters, component.Velocity]()

	s.keyBindingsRes = generic.NewResource[resource.KeyBindings](w)
}

func (s *PlayerControls) PostUpdateUI(w *ecs.World) {
}

func (s *PlayerControls) UpdateUI(w *ecs.World) {
	keyBindings := s.keyBindingsRes.Get()

	query := s.filter.Query(w)
	for query.Next() {
		heading, thrusters, velocity := query.Get()

		if keyBindings.TurnCCW.IsDown() {
			*heading -= component.Heading(thrusters.Turn)
		}

		if keyBindings.TurnCW.IsDown() {
			*heading += component.Heading(thrusters.Turn)
		}

		if keyBindings.FireThrusters.IsDown() {
			thrust := geo.Vec2{Y: thrusters.Forward}
			thrust = thrust.Rotate(float32(*heading))
			newVelocity := geo.Vec2(*velocity).Add(thrust)
			velocity.X = newVelocity.X
			velocity.Y = newVelocity.Y
		}
	}
}

var _ model.UISystem = (*PlayerControls)(nil)
