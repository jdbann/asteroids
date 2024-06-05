package system

import (
	"github.com/jdbann/asteroids/component"
	"github.com/jdbann/asteroids/resource"
	"github.com/jdbann/asteroids/util/geo"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type Wrap struct {
	filter *generic.Filter2[component.Polygon, component.Position]

	screenSizeRes generic.Resource[resource.ScreenSize]
}

func (s *Wrap) Finalize(w *ecs.World) {
}

func (s *Wrap) Initialize(w *ecs.World) {
	s.filter = generic.NewFilter2[component.Polygon, component.Position]()

	s.screenSizeRes = generic.NewResource[resource.ScreenSize](w)
}

func (s *Wrap) Update(w *ecs.World) {
	screenSizeRes := s.screenSizeRes.Get()
	screenSize := geo.Rectangle(*screenSizeRes)

	query := s.filter.Query(w)
	for query.Next() {
		polygon, position := query.Get()
		bb := geo.Polygon(*polygon).Translate(geo.Vec2(*position)).BoundingBox()

		if bb.Overlaps(screenSize) {
			continue
		}

		wrapPosition := screenSize.Inset(-bb.Dx()/2, -bb.Dy()/2).WrapVec2(geo.Vec2(*position))

		position.X = wrapPosition.X
		position.Y = wrapPosition.Y
	}
}

var _ model.System = (*Wrap)(nil)
