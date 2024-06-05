package uisystem

import (
	"github.com/fzipp/geom"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jdbann/asteroids/component"
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

		drawPolygon(*polygon, geom.Vec2(*position))
	}
}

func drawPolygon(polygon []geom.Vec2, position geom.Vec2) {
	var vB geom.Vec2
	for i, vA := range polygon {
		if i == 0 {
			vB = (polygon)[len(polygon)-1]
		}

		rl.DrawLineV(rl.Vector2(vA.Add(position)), rl.Vector2(vB.Add(position)), rl.Black)

		vB = vA
	}
}

var _ model.UISystem = (*Render)(nil)
