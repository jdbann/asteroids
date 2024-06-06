package component

import (
	"github.com/jdbann/asteroids/util/geo"
)

type Body struct {
	Polygon geo.Polygon
}

type Cannon struct {
	Offset   geo.Vec2
	Velocity float32
}

type Forces struct {
	Rotation float32
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

type Wrap struct{}
