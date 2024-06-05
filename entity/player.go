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
	builder generic.Map5[component.Cannon, component.Forces, component.Thrusters, component.Polygon, component.Position]
	rng     *rand.Rand

	positionBounds geo.Rectangle
}

func NewPlayerBuilder(w *ecs.World) *PlayerBuilder {
	screenSize := ecs.GetResource[resource.ScreenSize](w)

	return &PlayerBuilder{
		builder: generic.NewMap5[component.Cannon, component.Forces, component.Thrusters, component.Polygon, component.Position](w),
		rng:     rand.New(ecs.GetResource[resource.Rand](w)),

		positionBounds: geo.Rectangle(*screenSize),
	}
}

func (b *PlayerBuilder) Build() {
	b.builder.NewWith(
		&component.Cannon{
			Offset:   geo.V2(0, 6),
			Velocity: 3,
		},
		&component.Forces{
			Friction: .01,
		},
		&component.Thrusters{
			Forward: .2,
			Turn:    math.Pi * 3 / 180,
		},
		b.polygon(),
		b.position(),
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
	return &component.Position{
		Coords:  positionInRectangle(b.rng, b.positionBounds),
		Heading: b.rng.Float32() * math.Pi * 2,
	}
}
