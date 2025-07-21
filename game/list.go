package game

type List[T any] struct {
	head T
	tail *List[T]
}

func (l List[T]) prepend(v T) List[T] {
	return List[T]{head: v, tail: &l}
}
