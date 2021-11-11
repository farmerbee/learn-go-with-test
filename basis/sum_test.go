package main

import (
	"reflect"
	"testing"
)

func equalSlice(t *testing.T, want, got []int) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want:%v, got:%v\n", want, got)
	}
}
func TestSum(t *testing.T) {
	t.Run("test Sum #1", func(t *testing.T) {
		got := Sum([]int{1, 2, 3}, []int{4, 5, 6}, []int{})
		want := []int{6, 15, 0}
		equalSlice(t, want, got)
	})
}

func TestSumTail(t *testing.T) {
	t.Run("test SumTails #1", func(t *testing.T) {
		got := SumTails([]int{1, 2, 3}, []int{4, 5, 6}, []int{})
		want := []int{5, 11, 0}
		equalSlice(t, want, got)
	})
}
