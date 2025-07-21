package game

func clamp[T int](a, aMin, aMax T) T {
	return min(max(a, aMin), aMax)
}

func min[T int](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T int](a, b T) T {
	if a > b {
		return a
	}
	return b
}
