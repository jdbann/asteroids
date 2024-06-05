package system

import (
	"github.com/jdbann/asteroids/component"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type Collision struct {
	groupFilter  *generic.Filter1[component.CollisionParams]
	memberFilter *generic.Filter2[component.Body, component.Position]
	targetFilter *generic.Filter2[component.Body, component.Position]
}

func (s *Collision) Finalize(w *ecs.World) {
}

func (s *Collision) Initialize(w *ecs.World) {
	s.groupFilter = generic.NewFilter1[component.CollisionParams]()
	s.memberFilter = generic.NewFilter2[component.Body, component.Position]().
		With(generic.T[component.CollisionGroup]()).
		WithRelation(generic.T[component.CollisionGroup]())
	s.targetFilter = generic.NewFilter2[component.Body, component.Position]().
		With(generic.T[component.CollisionGroup]()).
		WithRelation(generic.T[component.CollisionGroup]())
}

func (s *Collision) Update(w *ecs.World) {
	var toRemove []ecs.Entity

	// Iterate through the collision groups
	groupQuery := s.groupFilter.Query(w)
	for groupQuery.Next() {
		collisionParams := groupQuery.Get()
		collisionGroup := groupQuery.Entity()

		// Iterate over each group that should be destroyed on collision
		for _, targetGroup := range collisionParams.DestroyGroups {
			memberQuery := s.memberFilter.Query(w, collisionGroup)

			// Loop over members of the primary group
			for memberQuery.Next() {
				memberBody, memberPosition := memberQuery.Get()
				memberPolygon := memberBody.Polygon.Rotate(memberPosition.Heading).Translate(memberPosition.Coords)

				// Loop over members of the target group
				targetQuery := s.targetFilter.Query(w, targetGroup)
				for targetQuery.Next() {
					targetBody, targetPosition := targetQuery.Get()
					targetPolygon := targetBody.Polygon.Rotate(targetPosition.Heading).Translate(targetPosition.Coords)

					// If these polygons collide, remove the appropriate entities
					if memberPolygon.InPolygon(targetPolygon) {
						toRemove = append(toRemove, targetQuery.Entity())

						if collisionParams.DestroySelf {
							toRemove = append(toRemove, memberQuery.Entity())
						}
					}
				}
			}
		}
	}

	for _, e := range toRemove {
		if !w.Alive(e) {
			continue
		}
		w.RemoveEntity(e)
	}
}

var _ model.System = (*Collision)(nil)
