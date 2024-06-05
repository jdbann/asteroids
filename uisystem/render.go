package uisystem

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jdbann/asteroids/component"
	"github.com/jdbann/asteroids/util/geo"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type Render struct {
	filter *generic.Filter2[component.Polygon, component.Position]
}

func (s *Render) FinalizeUI(w *ecs.World) {
}

func (s *Render) InitializeUI(w *ecs.World) {
	s.filter = generic.NewFilter2[component.Polygon, component.Position]()
}

func (s *Render) PostUpdateUI(w *ecs.World) {
}

func (s *Render) UpdateUI(w *ecs.World) {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.RayWhite)

	query := s.filter.Query(w)
	for query.Next() {
		polygon, position := query.Get()

		drawPolygon(geo.Polygon(*polygon).Translate(geo.Vec2(*position)))
	}
}

func drawPolygon(polygon geo.Polygon) {
	for i := 0; i < polygon.Edges(); i++ {
		a, b := polygon.Edge(i)
		rl.DrawLineV(rl.Vector2(a), rl.Vector2(b), rl.Black)
	}
}

var _ model.UISystem = (*Render)(nil)
