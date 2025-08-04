package game

type Set[T any] struct {
	items []*T
}

func newSet[T any]() *Set[T] {
	return &Set[T]{items: make([]*T, 0)}
}

func (s *Set[T]) add(v *T) {
	for i, x := range s.items {
		if x == nil {
			s.items[i] = v
			return
		}
	}
	s.items = append(s.items, v)
}

func (s *Set[T]) remove(i int) {
	s.items[i] = nil
}

func (s *Set[T]) len() int {
	res := 0
	for _, x := range s.items {
		if x != nil {
			res++
		}
	}
	return res
}

func (s *Set[T]) iter() []*T {
	return s.items
}

func (s *Set[T]) empty() bool {
	for _, x := range s.items {
		if x != nil {
			return false
		}
	}
	return true
}
