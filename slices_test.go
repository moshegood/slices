package slices

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	isEven := func(x int) bool { return x%2 == 0 }
	expected := []int{2, 4}
	output := Filter(input, isEven)
	if !reflect.DeepEqual(expected, output) {
		t.Errorf("expected output %v, got %v", expected, output)
	}
}

func TestMap(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	half := func(x int) float64 { return float64(x) / 2 }
	expected := []float64{0.5, 1, 1.5, 2, 2.5}
	output := Map(input, half)
	if !reflect.DeepEqual(expected, output) {
		t.Errorf("expected output %v, got %v", expected, output)
	}
}

func TestReduce(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	less := func(x, y int) int { return x - y }
	expected := 7
	output := Reduce(input, less, 22)
	if expected != output {
		t.Errorf("expected output %v, got %v", expected, output)
	}
}
