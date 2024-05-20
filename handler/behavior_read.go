package handler

type ReadBehavior[I any] struct {
	DefaultBehavior
	Input I
}

func (b *ReadBehavior[I]) Init() {
	b.Mods = []Modificator{
		&InputModificator[I]{Input: &b.Input},
	}
}
