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
	builder generic.Map6[component.Body, component.Cannon, component.CollisionGroup, component.Forces, component.Thrusters, component.Position]
	rng     *rand.Rand

	collisionGroup ecs.Entity
	positionBounds geo.Rectangle
}

func NewPlayerBuilder(w *ecs.World, collisionGroup ecs.Entity) *PlayerBuilder {
	screenSize := ecs.GetResource[resource.ScreenSize](w)

	return &PlayerBuilder{
		builder: generic.NewMap6[component.Body, component.Cannon, component.CollisionGroup, component.Forces, component.Thrusters, component.Position](w, generic.T[component.CollisionGroup]()),
		rng:     rand.New(ecs.GetResource[resource.Rand](w)),

		collisionGroup: collisionGroup,
		positionBounds: geo.Rectangle(*screenSize),
	}
}

func (b *PlayerBuilder) Build() {
	query := b.builder.NewBatchQ(1, b.collisionGroup)
	for query.Next() {
		body, cannon, _, forces, thrusters, position := query.Get()
		*body = *b.polygon()
		*cannon = component.Cannon{
			Offset:   geo.V2(0, 6),
			Velocity: 3,
		}
		*forces = component.Forces{
			Friction: .01,
		}
		*thrusters = component.Thrusters{
			Forward: .2,
			Turn:    math.Pi * 3 / 180,
		}
		*position = *b.position()
	}
}

func (b *PlayerBuilder) polygon() *component.Body {
	return &component.Body{
		Polygon: geo.Polygon{
			Vertices: []geo.Vec2{
				{X: 0, Y: 10},
				{X: 6, Y: -10},
				{X: -6, Y: -10},
			},
		},
	}
}

func (b *PlayerBuilder) position() *component.Position {
	return &component.Position{
		Coords:  positionInRectangle(b.rng, b.positionBounds),
		Heading: b.rng.Float32() * math.Pi * 2,
	}
}
