package entity

import (
	"github.com/jdbann/asteroids/component"
	"github.com/jdbann/asteroids/util/geo"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type RoundBuilder struct {
	builder generic.Map3[component.Body, component.Forces, component.Position]

	toAdd []roundComponents
}

func NewRoundBuilder(w *ecs.World) *RoundBuilder {
	return &RoundBuilder{
		builder: generic.NewMap3[component.Body, component.Forces, component.Position](w),
	}
}

func (b *RoundBuilder) Add(position, velocity geo.Vec2) {
	b.toAdd = append(b.toAdd, roundComponents{
		&component.Body{
			Polygon: geo.Polygon{
				Vertices: []geo.Vec2{
					{},
					velocity.Invert(),
				},
			},
		},
		&component.Forces{
			Velocity: velocity,
		},
		&component.Position{
			Coords: position,
		},
	})
}

func (b *RoundBuilder) Build() {
	for _, c := range b.toAdd {
		b.builder.NewWith(
			c.body,
			c.forces,
			c.position,
		)
	}
	b.toAdd = b.toAdd[:0]
}

type roundComponents struct {
	body     *component.Body
	forces   *component.Forces
	position *component.Position
}
