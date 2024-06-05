package uisystem

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jdbann/asteroids/resource"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type Window struct {
	screenSizeRes  generic.Resource[resource.ScreenSize]
	terminationRes generic.Resource[resource.Termination]
}

func (s *Window) FinalizeUI(w *ecs.World) {
	rl.CloseWindow()
}

func (s *Window) InitializeUI(w *ecs.World) {
	s.screenSizeRes = generic.NewResource[resource.ScreenSize](w)
	s.terminationRes = generic.NewResource[resource.Termination](w)

	screenSize := s.screenSizeRes.Get()

	rl.InitWindow(int32(screenSize.To.X), int32(screenSize.To.Y), "asteroids")
}

func (s *Window) PostUpdateUI(w *ecs.World) {
	s.terminationRes.Get().Terminate = rl.WindowShouldClose()
}

func (s *Window) UpdateUI(w *ecs.World) {
}

var _ model.UISystem = (*Window)(nil)
