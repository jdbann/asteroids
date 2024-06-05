package entity

import (
	"golang.org/x/exp/rand"

	"github.com/jdbann/asteroids/util/geo"
)

func positionInRectangle(rng *rand.Rand, rec geo.Rectangle) geo.Vec2 {
	return geo.Vec2{
		X: rec.Min.X + rng.Float32()*(rec.Max.X-rec.Min.X),
		Y: rec.Min.Y + rng.Float32()*(rec.Max.Y-rec.Min.Y),
	}
}
