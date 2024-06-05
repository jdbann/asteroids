package entity

import (
	"golang.org/x/exp/rand"

	"github.com/jdbann/asteroids/component"
	"github.com/jdbann/asteroids/util/geo"
)

func positionInRectangle(rng *rand.Rand, rec geo.Rectangle) *component.Position {
	return &component.Position{
		X: rec.Min.X + rng.Float32()*(rec.Max.X-rec.Min.X),
		Y: rec.Min.Y + rng.Float32()*(rec.Max.Y-rec.Min.Y),
	}
}
