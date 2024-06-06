package entity

import (
	"math"

	"golang.org/x/exp/rand"

	"github.com/jdbann/asteroids/component"
	"github.com/jdbann/asteroids/resource"
	"github.com/jdbann/asteroids/util/geo"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type AsteroidBuilder struct {
	builder generic.Map5[component.Body, component.CollisionGroup, component.Forces, component.Position, component.Wrap]
	rng     *rand.Rand

	collisionGroup           ecs.Entity
	radiusMin, radiusMax     float32
	roughness                float32
	sidesMin, sidesMax       int32
	positionBounds           geo.Rectangle
	velocityMin, velocityMax geo.Vec2
}

func NewAsteroidBuilder(w *ecs.World, collisionGroup ecs.Entity) *AsteroidBuilder {
	screenSize := ecs.GetResource[resource.ScreenSize](w)

	return &AsteroidBuilder{
		builder: generic.NewMap5[component.Body, component.CollisionGroup, component.Forces, component.Position, component.Wrap](w, generic.T[component.CollisionGroup]()),
		rng:     rand.New(ecs.GetResource[resource.Rand](w)),

		collisionGroup: collisionGroup,
		radiusMin:      50,
		radiusMax:      100,
		sidesMin:       9,
		sidesMax:       18,
		roughness:      0.2,
		positionBounds: geo.Rectangle(*screenSize),
		velocityMin:    geo.Vec2{X: -2, Y: -2},
		velocityMax:    geo.Vec2{X: 2, Y: 2},
	}
}

func (b *AsteroidBuilder) BuildBatch(count int) {
	query := b.builder.NewBatchQ(count, b.collisionGroup)
	for query.Next() {
		body, _, forces, position, _ := query.Get()
		*body = *b.body()
		*forces = *b.forces()
		*position = *b.position()
	}
}

func (b *AsteroidBuilder) body() *component.Body {
	radius := b.radiusMin + b.rng.Float32()*(b.radiusMax-b.radiusMin)
	sides := b.sidesMin + b.rng.Int31n(b.sidesMax-b.sidesMin)
	angle := math.Pi * 2 / float64(sides)

	p := geo.Polygon{}
	p.Vertices = make([]geo.Vec2, sides)
	for i := range p.Vertices {
		vRadius := radius * (1 + (b.roughness * 2 * b.rng.Float32()) - b.roughness)
		p.Vertices[i] = geo.Vec2{
			X: vRadius * float32(math.Sin(angle*float64(i))),
			Y: vRadius * float32(math.Cos(angle*float64(i))),
		}
	}

	return &component.Body{
		Polygon: p,
	}
}

func (b *AsteroidBuilder) forces() *component.Forces {
	return &component.Forces{
		Rotation: (b.rng.Float32() - 0.5) * math.Pi / 180,
		Velocity: geo.Vec2{
			X: b.velocityMin.X + b.rng.Float32()*(b.velocityMax.X-b.velocityMin.X),
			Y: b.velocityMin.Y + b.rng.Float32()*(b.velocityMax.Y-b.velocityMin.Y),
		},
	}
}

func (b *AsteroidBuilder) position() *component.Position {
	return &component.Position{
		Coords: positionInRectangle(b.rng, b.positionBounds),
	}
}
