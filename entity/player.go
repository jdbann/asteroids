package entity

import (
	"math"

	"github.com/jdbann/asteroids/component"
	"github.com/jdbann/asteroids/resource"
	"github.com/jdbann/asteroids/util/geo"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"golang.org/x/exp/rand"
)

type PlayerBuilder struct {
	builder generic.Map6[component.Heading, component.Friction, component.Thrusters, component.Polygon, component.Position, component.Velocity]
	rng     *rand.Rand

	positionBounds geo.Rectangle
}

func NewPlayerBuilder(w *ecs.World) *PlayerBuilder {
	screenSize := ecs.GetResource[resource.ScreenSize](w)

	return &PlayerBuilder{
		builder: generic.NewMap6[component.Heading, component.Friction, component.Thrusters, component.Polygon, component.Position, component.Velocity](w),
		rng:     rand.New(ecs.GetResource[resource.Rand](w)),

		positionBounds: geo.Rectangle(*screenSize),
	}
}

func (b *PlayerBuilder) Build() {
	heading := component.Heading(b.rng.Float32() * math.Pi * 2)
	friction := component.Friction(.01)
	b.builder.NewWith(
		&heading,
		&friction,
		&component.Thrusters{
			Forward: .2,
			Turn:    math.Pi * 3 / 180,
		},
		b.polygon(),
		b.position(),
		&component.Velocity{},
	)
}

func (b *PlayerBuilder) polygon() *component.Polygon {
	return &component.Polygon{
		Vertices: []geo.Vec2{
			{X: 0, Y: 10},
			{X: 6, Y: -10},
			{X: -6, Y: -10},
		},
	}
}

func (b *PlayerBuilder) position() *component.Position {
	return positionInRectangle(b.rng, b.positionBounds)
}
