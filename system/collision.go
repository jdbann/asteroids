package system

import (
	"github.com/jdbann/asteroids/component"
	"github.com/jdbann/asteroids/entity"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type Collision struct {
	groupFilter  *generic.Filter1[component.CollisionParams]
	memberFilter *generic.Filter2[component.Body, component.Position]
	targetFilter *generic.Filter2[component.Body, component.Position]

	explosionBuilderRes generic.Resource[entity.ExplosionBuilder]

	toExplode []ecs.Entity
	toRemove  []ecs.Entity
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

	s.explosionBuilderRes = generic.NewResource[entity.ExplosionBuilder](w)
}

func (s *Collision) Update(w *ecs.World) {
	// Iterate through the collision groups
	groupQuery := s.groupFilter.Query(w)
	for groupQuery.Next() {
		collisionParams := groupQuery.Get()
		collisionGroup := groupQuery.Entity()

		// Iterate over each group that should be destroyed on collision
		for targetGroup, actions := range collisionParams.Interactions {
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
						s.prepareCollisionAction(memberQuery.Entity(), actions.Self)
						s.prepareCollisionAction(targetQuery.Entity(), actions.Other)
					}
				}
			}
		}
	}

	s.handleCollisionActions(w)
}

func (s *Collision) prepareCollisionAction(entity ecs.Entity, action component.CollisionAction) {
	switch action {
	case component.CollisionActionExplode:
		s.toExplode = append(s.toExplode, entity)
	case component.CollisionActionRemove:
		s.toRemove = append(s.toRemove, entity)
	}
}

func (s *Collision) handleCollisionActions(w *ecs.World) {
	for _, e := range s.toRemove {
		if !w.Alive(e) {
			continue
		}
		w.RemoveEntity(e)
	}
	s.toRemove = s.toRemove[:0]

	explosionBuilder := s.explosionBuilderRes.Get()

	for _, e := range s.toExplode {
		if !w.Alive(e) {
			continue
		}
		explosionBuilder.BuildFrom(e)
		w.RemoveEntity(e)
	}
	s.toExplode = s.toExplode[:0]
}

var _ model.System = (*Collision)(nil)
