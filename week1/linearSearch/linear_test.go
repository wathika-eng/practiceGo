package main

import "testing"

func TestLinear(t *testing.T) {
	got, val := search([]int{1, 2, 6, 10, 100, 4, -9, 10}, 4)
	want := 5
	val2 := true
	if got != want && val != val2 {
		t.Errorf("%v is not equal to %v\n", got, want)
	}
}
