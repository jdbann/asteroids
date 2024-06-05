package resource

import "github.com/jdbann/asteroids/util/keys"

type KeyBindings struct {
	FireThrusters keys.Binding
	TurnCCW       keys.Binding
	TurnCW        keys.Binding
}

func DefaultKeyBindings() *KeyBindings {
	return &KeyBindings{
		FireThrusters: keys.NewBinding(keys.WithKey(keys.KeyUp)),
		TurnCCW:       keys.NewBinding(keys.WithKey(keys.KeyLeft)),
		TurnCW:        keys.NewBinding(keys.WithKey(keys.KeyRight)),
	}
}
