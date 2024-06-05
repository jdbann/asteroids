package component

import (
	"github.com/jdbann/asteroids/util/geo"
)

type Cannon struct {
	Offset   geo.Vec2
	Velocity float32
}

type Forces struct {
	Friction float32
	Velocity geo.Vec2
}

type Polygon geo.Polygon

type Position struct {
	Coords  geo.Vec2
	Heading float32
}

type Thrusters struct {
	Forward float32
	Turn    float32
}
