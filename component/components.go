package component

import (
	"github.com/jdbann/asteroids/util/geo"
	"github.com/mlange-42/arche/ecs"
)

type Body struct {
	Polygon geo.Polygon
}

type Cannon struct {
	Offset   geo.Vec2
	Velocity float32
}

type CollisionGroup struct {
	ecs.Relation
}

type CollisionParams struct {
	DestroyGroups []ecs.Entity
}

type Forces struct {
	Friction float32
	Velocity geo.Vec2
}

type Position struct {
	Coords  geo.Vec2
	Heading float32
}

type Thrusters struct {
	Forward float32
	Turn    float32
}
