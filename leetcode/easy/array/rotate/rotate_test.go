package rotate

import (
	"reflect"
	"testing"
)

func TestRotateSpace(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		exp  []int
	}{
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    2,
			exp:  []int{6, 7, 1, 2, 3, 4, 5},
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			exp:  []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums: []int{-1, -100, 3, 99},
			k:    2,
			exp:  []int{3, 99, -1, -100},
		},
	}

	for _, tt := range tests {
		rotateV1(tt.nums, tt.k)
		if !reflect.DeepEqual(tt.exp, tt.nums) {
			t.Errorf("exp: %v, act: %v", tt.exp, tt.nums)
		}
	}
}

func TestRotateTime(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		exp  []int
	}{
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    2,
			exp:  []int{6, 7, 1, 2, 3, 4, 5},
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			exp:  []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums: []int{-1, -100, 3, 99},
			k:    2,
			exp:  []int{3, 99, -1, -100},
		},
	}

	for _, tt := range tests {
		rotateV2(tt.nums, tt.k)
		if !reflect.DeepEqual(tt.exp, tt.nums) {
			t.Errorf("exp: %v, act: %v", tt.exp, tt.nums)
		}
	}
}
