package main

import (
	"fmt"
	"testing"
)

func compareSlices(a,b []int)bool{
	if len(a) != len(b){
		return false
	}

	for i := range a{
		if a[i] != b[i]{
			return false
		}
	}

	return true
}

func Test(t *testing.T) {
	var tests = []struct {
		array []int
		target int
		want []int
	}{
		{array: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{array: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{array: []int{3, 3}, target: 6, want: []int{0, 1}},
		{array: []int{2, 7, 11, 15}, target: 9, want: []int{1, 2}},
	}
	for _, test := range tests{
		testname := fmt.Sprintf("%v,%d", test.array, test.target)
		t.Run(testname, func(t *testing.T){
			ans := twoSum(test.array, test.target)
			if !compareSlices(ans, test.want) {
				t.Errorf("ans %v, want %v", ans, test.want)
			}
		})
	}
}
