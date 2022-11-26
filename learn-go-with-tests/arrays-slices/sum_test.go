package arraysslices

import (
	"reflect"
	"testing"
)

// [01]
/* func TestSum(t *testing.T) {
  numbers := [5]int{1,2,3,4,5}

  got := Sum(numbers)
  want := 15

  if got != want {
    t.Errorf("expected %d want %d given %v", got, want, numbers)
  }
} */

func TestSum(t *testing.T) {
	/* 	t.Run("collection of 5 numbers", func(t *testing.T) {
		// [01]
		// numbers := [5]int{1, 2, 3, 4, 5}
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("expected %d want %d given %v", got, want, numbers)
		}
	}) */

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("expected %d want %d given %v", got, want, numbers)
		}
	})
}

// go test -cover

func TestSumAll(t *testing.T) {
	input := [][]int{{1, 2}, {0, 9}}
	// got := SumAll(input...)
	got := SumAll(input)
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v want %v given %v", got, want, input)
	}
}

// [03]
/* func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		input := [][]int{{1, 2}, {0, 9}}

		got := SumAllTails(input)
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v want %v given %v", got, want, input)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		input := [][]int{{}, {3,4,5}}

	  got := SumAllTails(input)
	  want := []int{0,9}

	  if !reflect.DeepEqual(got, want) {
	    t.Errorf("expected %v want %v given %v", got, want, input)
	  }
	})
} */

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int, input [][]int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			// t.Errorf("expected %v want %v", got, want)
			t.Errorf("expected %v want %v given %v", got, want, input)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		input := [][]int{{1, 2}, {0, 9}}

		got := SumAllTails(input)
		want := []int{2, 9}

		checkSums(t, got, want, input)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		input := [][]int{{}, {3, 4, 5}}

		got := SumAllTails(input)
		want := []int{0, 9}

		checkSums(t, got, want, input)
	})
}

func BenchmarkSumAllTails(b *testing.B) {
  for i := 0; i < b.N; i++ {
		input := [][]int{{}, {3, 4, 5}}
    SumAllTails(input)
  }
}
