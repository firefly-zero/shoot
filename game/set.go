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

func (s *Set[T]) len() int {
	res := 0
	for _, x := range s.items {
		if x != nil {
			res++
		}
	}
	return res
}

func (s *Set[T]) iter() *SetIter[T] {
	return &SetIter[T]{items: s.items}
}

type SetIter[T any] struct {
	items []*T
	i     int
}

func (s *SetIter[T]) remove() {
	s.items[s.i-1] = nil
}

func (s *SetIter[T]) next() *T {
	for s.i < len(s.items) {
		v := s.items[s.i]
		s.i++
		if v != nil {
			return v
		}
	}
	return nil
}
