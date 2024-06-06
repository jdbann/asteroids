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
	filter *generic.Filter3[component.Body, component.Position, component.Wrap]

	screenSizeRes generic.Resource[resource.ScreenSize]

	toRemove []ecs.Entity
}

func (s *Wrap) Finalize(w *ecs.World) {
}

func (s *Wrap) Initialize(w *ecs.World) {
	s.filter = generic.NewFilter3[component.Body, component.Position, component.Wrap]().
		Optional(generic.T[component.Wrap]())

	s.screenSizeRes = generic.NewResource[resource.ScreenSize](w)
}

func (s *Wrap) Update(w *ecs.World) {
	screenSizeRes := s.screenSizeRes.Get()
	screenSize := geo.Rectangle(*screenSizeRes)

	query := s.filter.Query(w)
	for query.Next() {
		body, position, wrap := query.Get()
		bb := body.Polygon.Translate(position.Coords).BoundingBox()

		if bb.Overlaps(screenSize) {
			continue
		}

		if wrap == nil {
			s.toRemove = append(s.toRemove, query.Entity())
			continue
		}

		wrapCoords := screenSize.Inset(-bb.Dx()/2, -bb.Dy()/2).WrapVec2(position.Coords)

		position.Coords = wrapCoords
	}

	for _, e := range s.toRemove {
		if w.Alive(e) {
			w.RemoveEntity(e)
		}
	}
	s.toRemove = s.toRemove[:0]
}

var _ model.System = (*Wrap)(nil)
