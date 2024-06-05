package entity

import (
	"github.com/jdbann/asteroids/component"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type CollisionGroupBuilder struct {
	builder generic.Map1[component.CollisionParams]

	AsteroidGroup ecs.Entity
	PlayerGroup   ecs.Entity
	RoundGroup    ecs.Entity
}

func NewCollisionGroupBuilder(w *ecs.World) *CollisionGroupBuilder {
	return &CollisionGroupBuilder{
		builder:       generic.NewMap1[component.CollisionParams](w),
		AsteroidGroup: w.NewEntity(),
		PlayerGroup:   w.NewEntity(),
		RoundGroup:    w.NewEntity(),
	}
}

func (b *CollisionGroupBuilder) Build() {
	b.PlayerGroup = b.builder.NewWith(&component.CollisionParams{
		DestroyGroups: nil,
	})
	b.AsteroidGroup = b.builder.NewWith(&component.CollisionParams{
		DestroyGroups: []ecs.Entity{b.PlayerGroup},
	})
	b.RoundGroup = b.builder.NewWith(&component.CollisionParams{
		DestroyGroups: []ecs.Entity{b.AsteroidGroup},
		DestroySelf:   true,
	})
}
