package main

import (
	"github.com/jdbann/asteroids/entity"
	"github.com/jdbann/asteroids/system"
	"github.com/jdbann/asteroids/uisystem"
	"github.com/mlange-42/arche-model/model"
)

func main() {
	m := model.New()

	m.FPS = 60
	m.TPS = 60

	m.AddUISystem(&uisystem.Window{})
	m.AddUISystem(&uisystem.Render{})

	m.AddSystem(&system.Movement{})

	entity.NewAsteroidBuilder(&m.World).Build(10)

	m.Run()
}
