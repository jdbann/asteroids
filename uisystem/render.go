package uisystem

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
)

type Render struct {
	message   string
	fontSize  int32
	textWidth int32
}

func (s *Render) FinalizeUI(w *ecs.World) {
}

func (s *Render) InitializeUI(w *ecs.World) {
	s.message = "Let's play asteroids!"
	s.fontSize = 10
	s.textWidth = rl.MeasureText(s.message, s.fontSize)
}

func (s *Render) PostUpdateUI(w *ecs.World) {
}

func (s *Render) UpdateUI(w *ecs.World) {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.RayWhite)

	rl.DrawText(s.message, (int32(rl.GetScreenWidth())-s.textWidth)/2, (int32(rl.GetScreenHeight())-s.fontSize)/2, s.fontSize, rl.Black)
}

var _ model.UISystem = (*Render)(nil)
