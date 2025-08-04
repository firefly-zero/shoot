package game

func clamp[T int](a, aMin, aMax T) T {
	if a <= aMin {
		return aMin
	}
	if a >= aMax {
		return aMax
	}
	return a
}
