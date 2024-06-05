package resource

import "github.com/jdbann/asteroids/util/keys"

type KeyBindings struct {
	FireThrusters keys.Binding
}

func DefaultKeyBindings() *KeyBindings {
	return &KeyBindings{
		FireThrusters: keys.NewBinding(
			keys.WithKey(keys.KeyUp),
		),
	}
}
