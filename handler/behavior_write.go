package handler

type WriteBehavior[O any] struct {
	DefaultBehavior
	Output O
}

func (b *WriteBehavior[O]) Init() {
	b.Mods = []Modificator{
		&OutputModificator[O]{Output: &b.Output},
	}
}
