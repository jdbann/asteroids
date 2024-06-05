package component

import (
	"github.com/jdbann/asteroids/util/geo"
)

type Heading float32

type Friction float32

type Polygon geo.Polygon

type Position geo.Vec2

type Thrusters struct {
	Forward float32
	Turn    float32
}

type Velocity geo.Vec2
