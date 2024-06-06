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

type ExplosionBuilder struct {
	builder generic.Map3[component.Body, component.Forces, component.Position]
	rng     *rand.Rand
}

func NewExplosionBuilder(w *ecs.World) *ExplosionBuilder {
	return &ExplosionBuilder{
		builder: generic.NewMap3[component.Body, component.Forces, component.Position](w),
		rng:     rand.New(ecs.GetResource[resource.Rand](w)),
	}
}

func (b *ExplosionBuilder) BuildFrom(entity ecs.Entity) {
	sourceBody, sourceForces, sourcePosition := b.builder.Get(entity)

	if sourceBody == nil || sourceForces == nil || sourcePosition == nil {
		panic("attempted to exploded an unexplodable entity")
	}

	sourcePolygon := sourceBody.Polygon.Rotate(sourcePosition.Heading).Translate(sourcePosition.Coords)

	var i int
	query := b.builder.NewBatchQ(sourcePolygon.Edges())
	for query.Next() {
		body, forces, position := query.Get()

		start, end := sourcePolygon.Edge(i)
		mid := start.Add(end).Scale(0.5)

		body.Polygon = geo.Polygon{
			Vertices: []geo.Vec2{start.Sub(mid), end.Sub(mid)},
		}

		forces.Rotation = (b.rng.Float32() - 0.5) * math.Pi * 2 / 180
		forces.Velocity = sourceForces.Velocity.Add(
			mid.Sub(sourcePosition.Coords).Normalize().Scale(0.75),
		)

		position.Coords = mid

		i++
	}
}
