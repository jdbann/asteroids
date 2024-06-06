package component

import "github.com/mlange-42/arche/ecs"

type CollisionAction int

const (
	CollisionActionNone CollisionAction = iota
	CollisionActionExplode
	CollisionActionRemove
)

type CollisionActions struct {
	Self  CollisionAction
	Other CollisionAction
}

type CollisionGroup struct {
	ecs.Relation
}

type CollisionParams struct {
	Interactions map[ecs.Entity]CollisionActions
}
