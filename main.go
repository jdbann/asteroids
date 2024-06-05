package main

import (
	"github.com/jdbann/asteroids/uisystem"
	"github.com/mlange-42/arche-model/model"
)

func main() {
	m := model.New()

	m.AddUISystem(&uisystem.Window{})
	m.AddUISystem(&uisystem.Render{})

	m.Run()
}
