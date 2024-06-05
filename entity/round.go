package entity

import (
	"github.com/jdbann/asteroids/component"
	"github.com/jdbann/asteroids/util/geo"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type RoundBuilder struct {
	builder generic.Map3[component.Polygon, component.Position, component.Velocity]

	toAdd []roundComponents
}

func NewRoundBuilder(w *ecs.World) *RoundBuilder {
	return &RoundBuilder{
		builder: generic.NewMap3[component.Polygon, component.Position, component.Velocity](w),
	}
}

func (b *RoundBuilder) Add(position, velocity geo.Vec2) {
	b.toAdd = append(b.toAdd, roundComponents{
		&component.Polygon{
			Vertices: []geo.Vec2{
				{},
				velocity.Invert(),
			},
		},
		&component.Position{
			X: position.X,
			Y: position.Y,
		},
		&component.Velocity{
			X: velocity.X,
			Y: velocity.Y,
		},
	})
}

func (b *RoundBuilder) Build() {
	for _, c := range b.toAdd {
		b.builder.NewWith(
			c.polygon,
			c.position,
			c.velocity,
		)
	}
	b.toAdd = b.toAdd[:0]
}

type roundComponents struct {
	polygon  *component.Polygon
	position *component.Position
	velocity *component.Velocity
}
