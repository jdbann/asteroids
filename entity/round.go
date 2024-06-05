package entity

import (
	"github.com/jdbann/asteroids/component"
	"github.com/jdbann/asteroids/util/geo"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type RoundBuilder struct {
	builder generic.Map4[component.Body, component.CollisionGroup, component.Forces, component.Position]

	collisionGroup ecs.Entity

	toAdd []roundComponents
}

func NewRoundBuilder(w *ecs.World, collisionGroup ecs.Entity) *RoundBuilder {
	return &RoundBuilder{
		builder: generic.NewMap4[component.Body, component.CollisionGroup, component.Forces, component.Position](w, generic.T[component.CollisionGroup]()),

		collisionGroup: collisionGroup,
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
	if len(b.toAdd) == 0 {
		return
	}
	query := b.builder.NewBatchQ(len(b.toAdd), b.collisionGroup)
	var i int
	for query.Next() {
		body, _, forces, position := query.Get()
		*body = *b.toAdd[i].body
		*forces = *b.toAdd[i].forces
		*position = *b.toAdd[i].position
		i++
	}
	b.toAdd = b.toAdd[:0]
}

type roundComponents struct {
	body     *component.Body
	forces   *component.Forces
	position *component.Position
}
