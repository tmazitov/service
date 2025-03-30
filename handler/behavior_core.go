package handler

import mode "github.com/tmazitov/service/mode"

type CoreBehavior[I, O any] struct {
	DefaultBehavior
	Input  I
	Output O
}

func (b *CoreBehavior[I, O]) Init() {
	b.Mods = []mode.IMode{
		&mode.InputMode[I]{Input: &b.Input},
		&mode.OutputMode[O]{Output: &b.Output},
	}
}
