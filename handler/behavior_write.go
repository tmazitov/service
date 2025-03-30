package handler

import (
	mode "github.com/tmazitov/service/mode"
)

type WriteBehavior[O any] struct {
	DefaultBehavior
	Output O
}

func (b *WriteBehavior[O]) Init() {
	b.Mods = []mode.IMode{
		&mode.OutputMode[O]{Output: &b.Output},
	}
}
