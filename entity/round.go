package entity

import (
	"github.com/jdbann/asteroids/component"
	"github.com/jdbann/asteroids/util/geo"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type RoundBuilder struct {
	builder generic.Map3[component.Forces, component.Polygon, component.Position]

	toAdd []roundComponents
}

func NewRoundBuilder(w *ecs.World) *RoundBuilder {
	return &RoundBuilder{
		builder: generic.NewMap3[component.Forces, component.Polygon, component.Position](w),
	}
}

func (b *RoundBuilder) Add(position, velocity geo.Vec2) {
	b.toAdd = append(b.toAdd, roundComponents{
		&component.Forces{
			Velocity: velocity,
		},
		&component.Polygon{
			Vertices: []geo.Vec2{
				{},
				velocity.Invert(),
			},
		},
		&component.Position{
			Coords: position,
		},
	})
}

func (b *RoundBuilder) Build() {
	for _, c := range b.toAdd {
		b.builder.NewWith(
			c.forces,
			c.polygon,
			c.position,
		)
	}
	b.toAdd = b.toAdd[:0]
}

type roundComponents struct {
	forces   *component.Forces
	polygon  *component.Polygon
	position *component.Position
}
