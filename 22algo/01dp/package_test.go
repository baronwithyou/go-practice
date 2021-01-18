package dp

import "testing"

func TestSolve1(t *testing.T) {
	minePrices := []int{200, 300, 350, 400, 500}
	mineWorkload := []int{3, 4, 3, 5, 5}
	workerNum := 10

	want := 900
	if got := solve1(workerNum, minePrices, mineWorkload); got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestSolve2(t *testing.T) {
	minePrices := []int{200, 300, 350, 400, 500}
	mineWorkload := []int{3, 4, 3, 5, 5}
	workerNum := 10

	want := 900
	if got := solve2(workerNum, minePrices, mineWorkload); got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
