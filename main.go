package main

import (
	"github.com/jdbann/asteroids/entity"
	"github.com/jdbann/asteroids/resource"
	"github.com/jdbann/asteroids/system"
	"github.com/jdbann/asteroids/uisystem"
	"github.com/jdbann/asteroids/util/geo"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
)

func main() {
	m := model.New()

	m.FPS = 60
	m.TPS = 60

	m.AddUISystem(&uisystem.Window{})
	m.AddUISystem(&uisystem.Render{})
	m.AddUISystem(&uisystem.PlayerControls{})

	m.AddSystem(&system.Movement{})
	m.AddSystem(&system.Wrap{})
	m.AddSystem(&system.Collision{})

	ecs.AddResource(&m.World, &resource.ScreenSize{Max: geo.Vec2{X: 1280, Y: 720}})
	ecs.AddResource(&m.World, resource.DefaultKeyBindings())

	cgb := entity.NewCollisionGroupBuilder(&m.World)
	cgb.Build()

	entity.NewAsteroidBuilder(&m.World, cgb.AsteroidGroup).BuildBatch(10)
	entity.NewPlayerBuilder(&m.World, cgb.PlayerGroup).Build()

	ecs.AddResource(&m.World, entity.NewRoundBuilder(&m.World, cgb.RoundGroup))

	m.Run()
}
