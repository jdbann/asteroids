package uisystem

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type Window struct {
	terminationRes generic.Resource[resource.Termination]
}

func (s *Window) FinalizeUI(w *ecs.World) {
	rl.CloseWindow()
}

func (s *Window) InitializeUI(w *ecs.World) {
	s.terminationRes = generic.NewResource[resource.Termination](w)

	rl.InitWindow(1280, 720, "asteroids")
}

func (s *Window) PostUpdateUI(w *ecs.World) {
	s.terminationRes.Get().Terminate = rl.WindowShouldClose()
}

func (s *Window) UpdateUI(w *ecs.World) {
}

var _ model.UISystem = (*Window)(nil)
