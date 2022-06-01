package intersection

import (
	"reflect"
	"testing"
)

func TestIntersectV1(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		exp   []int
	}{
		{
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			exp:   []int{2, 2},
		},
		{
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			exp:   []int{9, 4},
		},
	}

	for _, tt := range tests {
		act := intersectV1(tt.nums1, tt.nums2)
		if !reflect.DeepEqual(tt.exp, act) {
			t.Errorf("exp: %v, act: %v", tt.exp, act)
		}
	}
}

func TestIntersectV2(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		exp   []int
	}{
		{
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			exp:   []int{2, 2},
		},
		{
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			exp:   []int{4, 9},
		},
	}

	for _, tt := range tests {
		act := intersectV2(tt.nums1, tt.nums2)
		if !reflect.DeepEqual(tt.exp, act) {
			t.Errorf("exp: %v, act: %v", tt.exp, act)
		}
	}
}
