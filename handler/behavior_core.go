package handler

type CoreBehavior[I, O any] struct {
	DefaultBehavior
	Input  I
	Output O
}

func (b *CoreBehavior[I, O]) Init() {
	b.Mods = []Modificator{
		&InputModificator[I]{Input: &b.Input},
		&OutputModificator[O]{Output: &b.Output},
	}
}
