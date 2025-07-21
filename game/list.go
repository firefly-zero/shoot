package game

type List[T any] struct {
	item T
	prev *List[T]
	next *List[T]
}

func (l *List[T]) prepend(v T) *List[T] {
	res := &List[T]{item: v, next: l}
	if l != nil {
		l.prev = res
	}
	return res
}
