package main

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	// 0-2 (1) + 2-1 (2) beats 0-1 (4); 0..3 best is 0-2-1-3 = 4.
	edges := []Edge{{0, 1, 4}, {0, 2, 1}, {2, 1, 2}, {1, 3, 1}, {2, 3, 5}}
	want := []int{0, 3, 1, 4}
	if got := Dijkstra(4, edges, 0); !reflect.DeepEqual(got, want) {
		t.Errorf("Dijkstra = %v, want %v", got, want)
	}

	// Unreachable vertex reported as -1.
	got := Dijkstra(3, []Edge{{0, 1, 7}}, 0)
	if !reflect.DeepEqual(got, []int{0, 7, -1}) {
		t.Errorf("unreachable = %v", got)
	}

	// Zero-weight edges are fine (non-negative is the only requirement).
	got = Dijkstra(2, []Edge{{0, 1, 0}}, 0)
	if !reflect.DeepEqual(got, []int{0, 0}) {
		t.Errorf("zero weight = %v", got)
	}

	// Source only.
	if got := Dijkstra(1, nil, 0); !reflect.DeepEqual(got, []int{0}) {
		t.Errorf("single vertex = %v", got)
	}
}
