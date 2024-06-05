package keys

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Binding struct {
	key Keys
}

type BindingOpt func(*Binding)

func NewBinding(opts ...BindingOpt) Binding {
	b := Binding{}
	for _, opt := range opts {
		opt(&b)
	}
	return b
}

func WithKey(k Key) BindingOpt {
	return func(b *Binding) { b.key = Keys{key: k} }
}

func (b Binding) IsDown() bool {
	return b.key.Down()
}

func (b Binding) IsPressed() bool {
	return b.key.Pressed()
}

type Keys struct {
	key Key
}

func (k Keys) Down() bool {
	return rl.IsKeyDown(int32(k.key))
}

func (k Keys) Pressed() bool {
	return rl.IsKeyPressed(int32(k.key))
}

type Key int32

const (
	KeySpace Key = rl.KeySpace
	KeyRight Key = rl.KeyRight
	KeyLeft  Key = rl.KeyLeft
	KeyUp    Key = rl.KeyUp
)
