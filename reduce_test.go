package reduce

import "testing"

func TestReduce(t *testing.T) {
	IntSumReducer := func(a, b int) int {
		return a + b
	}
	IntProductReducer := func(a, b int) int {
		return a * b
	}
	t.Run("with empty slice", func(t *testing.T) {
		reduction := Reduce([]int{}, IntSumReducer, 0)
		if reduction != 0 {
			t.Fatalf("expected %d but got %d", 0, reduction)
		}
	})

	t.Run("with int sum", func(t *testing.T) {
		reduction := Reduce([]int{1, 2, 3, 4, 5}, IntSumReducer, 0)
		if reduction != 15 {
			t.Fatalf("expected %d but got %d", 15, reduction)
		}
	})

	t.Run("with int product", func(t *testing.T) {
		reduction := Reduce([]int{-2, -2, -2}, IntProductReducer, 1)
		if reduction != -8 {
			t.Fatalf("expected %d but got %d", -8, reduction)
		}
	})
}
