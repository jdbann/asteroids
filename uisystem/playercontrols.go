package uisystem

import (
	"github.com/jdbann/asteroids/component"
	"github.com/jdbann/asteroids/entity"
	"github.com/jdbann/asteroids/resource"
	"github.com/jdbann/asteroids/util/geo"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type PlayerControls struct {
	filter *generic.Filter4[component.Cannon, component.Forces, component.Position, component.Thrusters]

	keyBindingsRes generic.Resource[resource.KeyBindings]

	roundBuilder *entity.RoundBuilder
}

func (s *PlayerControls) FinalizeUI(w *ecs.World) {
}

func (s *PlayerControls) InitializeUI(w *ecs.World) {
	s.filter = generic.NewFilter4[component.Cannon, component.Forces, component.Position, component.Thrusters]()

	s.keyBindingsRes = generic.NewResource[resource.KeyBindings](w)

	s.roundBuilder = entity.NewRoundBuilder(w)
}

func (s *PlayerControls) PostUpdateUI(w *ecs.World) {
	s.roundBuilder.Build()
}

func (s *PlayerControls) UpdateUI(w *ecs.World) {
	keyBindings := s.keyBindingsRes.Get()

	query := s.filter.Query(w)
	for query.Next() {
		cannon, forces, position, thrusters := query.Get()

		if keyBindings.FireCannon.IsPressed() {
			s.roundBuilder.Add(
				position.Coords.Add(cannon.Offset.Rotate(position.Heading)),
				geo.V2(0, cannon.Velocity).Rotate(position.Heading),
			)
		}

		if keyBindings.TurnCCW.IsDown() {
			position.Heading -= thrusters.Turn
		}

		if keyBindings.TurnCW.IsDown() {
			position.Heading += thrusters.Turn
		}

		if keyBindings.FireThrusters.IsDown() {
			thrust := geo.V2(0, thrusters.Forward).Rotate(position.Heading)
			forces.Velocity = forces.Velocity.Add(thrust)
		}
	}
}

var _ model.UISystem = (*PlayerControls)(nil)
