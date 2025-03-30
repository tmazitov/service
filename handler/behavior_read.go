package handler

import mode "github.com/tmazitov/service/mode"

type ReadBehavior[I any] struct {
	DefaultBehavior
	Input I
}

func (b *ReadBehavior[I]) Init() {
	b.Mods = []mode.IMode{
		&mode.InputMode[I]{Input: &b.Input},
	}
}
