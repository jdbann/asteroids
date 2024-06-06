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
	filter *generic.Filter3[component.Body, component.Disappear, component.Position]
}

func (s *Render) FinalizeUI(w *ecs.World) {
}

func (s *Render) InitializeUI(w *ecs.World) {
	s.filter = generic.NewFilter3[component.Body, component.Disappear, component.Position]().
		Optional(generic.T[component.Disappear]())
}

func (s *Render) PostUpdateUI(w *ecs.World) {
}

func (s *Render) UpdateUI(w *ecs.World) {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.RayWhite)

	query := s.filter.Query(w)
	for query.Next() {
		body, disappear, position := query.Get()

		color := rl.Black

		if disappear != nil {
			color = rl.ColorAlpha(color, 1-disappear.Progress)
			disappear.Progress += disappear.Rate
		}

		drawPolygon(
			body.Polygon.
				Rotate(position.Heading).
				Translate(position.Coords),
			color,
		)
	}
}

func drawPolygon(polygon geo.Polygon, color rl.Color) {
	for i := 0; i < polygon.Edges(); i++ {
		a, b := polygon.Edge(i)
		rl.DrawLineV(rl.Vector2(a), rl.Vector2(b), color)
	}
}

var _ model.UISystem = (*Render)(nil)
