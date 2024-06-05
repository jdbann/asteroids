package resource

import "github.com/jdbann/asteroids/util/keys"

type KeyBindings struct {
	FireCannon    keys.Binding
	FireThrusters keys.Binding
	TurnCCW       keys.Binding
	TurnCW        keys.Binding
}

func DefaultKeyBindings() *KeyBindings {
	return &KeyBindings{
		FireCannon:    keys.NewBinding(keys.WithKey(keys.KeySpace)),
		FireThrusters: keys.NewBinding(keys.WithKey(keys.KeyUp)),
		TurnCCW:       keys.NewBinding(keys.WithKey(keys.KeyLeft)),
		TurnCW:        keys.NewBinding(keys.WithKey(keys.KeyRight)),
	}
}
