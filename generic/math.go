package generic

func Add[T int | float64](a, b T) T {
	return a + b
}
