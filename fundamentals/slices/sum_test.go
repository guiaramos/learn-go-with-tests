package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectMessage(t, got, want, numbers)
	})

}

func ExampleSum() {
	numbers := []int{1, 2, 3}
	fmt.Println(Sum(numbers))

	// Output: 6
}

func TestSumAll(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("sum multiple slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		assertCorrectMessage(t, got, want)
	})

	t.Run("sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assertCorrectMessage(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		assertCorrectMessage(t, got, want)
	})

}

func ExampleSumAll() {
	got := SumAll([]int{1, 2}, []int{0, 9})
	fmt.Println(got)

	// Output:  [3 9]
}
