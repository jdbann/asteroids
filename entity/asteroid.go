package entity

import (
	"math"

	"golang.org/x/exp/rand"

	"github.com/fzipp/geom"
	"github.com/jdbann/asteroids/component"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type AsteroidBuilder struct {
	builder generic.Map3[component.Polygon, component.Position, component.Velocity]
	rng     *rand.Rand

	radiusMin, radiusMax     float32
	roughness                float32
	sidesMin, sidesMax       int32
	positionMin, positionMax geom.Vec2
	velocityMin, velocityMax geom.Vec2
}

func NewAsteroidBuilder(w *ecs.World) *AsteroidBuilder {
	return &AsteroidBuilder{
		builder: generic.NewMap3[component.Polygon, component.Position, component.Velocity](w),
		rng:     rand.New(ecs.GetResource[resource.Rand](w)),

		radiusMin:   50,
		radiusMax:   100,
		sidesMin:    9,
		sidesMax:    18,
		roughness:   0.2,
		positionMin: geom.Vec2{X: 0, Y: 0},
		positionMax: geom.Vec2{X: 1280, Y: 720},
		velocityMin: geom.Vec2{X: -2, Y: -2},
		velocityMax: geom.Vec2{X: 2, Y: 2},
	}
}

func (b *AsteroidBuilder) Build(count int) {
	query := b.builder.NewBatchQ(count)
	for query.Next() {
		polygon, position, velocity := query.Get()
		*polygon = *b.polygon()
		*position = *b.position()
		*velocity = *b.velocity()
	}
}

func (b *AsteroidBuilder) polygon() *component.Polygon {
	radius := b.radiusMin + b.rng.Float32()*(b.radiusMax-b.radiusMin)
	sides := b.sidesMin + b.rng.Int31n(b.sidesMax-b.sidesMin)
	angle := math.Pi * 2 / float64(sides)

	p := make(component.Polygon, sides)
	for i := range p {
		vRadius := radius * (1 + (b.roughness * 2 * b.rng.Float32()) - b.roughness)
		p[i] = geom.Vec2{
			X: vRadius * float32(math.Sin(angle*float64(i))),
			Y: vRadius * float32(math.Cos(angle*float64(i))),
		}
	}
	return &p
}

func (b *AsteroidBuilder) position() *component.Position {
	return &component.Position{
		X: b.positionMin.X + b.rng.Float32()*(b.positionMax.X-b.positionMin.X),
		Y: b.positionMin.Y + b.rng.Float32()*(b.positionMax.Y-b.positionMin.Y),
	}
}

func (b *AsteroidBuilder) velocity() *component.Velocity {
	return &component.Velocity{
		X: b.velocityMin.X + b.rng.Float32()*(b.velocityMax.X-b.velocityMin.X),
		Y: b.velocityMin.Y + b.rng.Float32()*(b.velocityMax.Y-b.velocityMin.Y),
	}
}
