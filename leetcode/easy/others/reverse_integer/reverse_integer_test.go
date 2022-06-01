package main

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		x   int
		exp int
	}{
		{
			x:   123,
			exp: 321,
		},
		{
			x:   -123,
			exp: -321,
		},
		{
			x:   120,
			exp: 21,
		},
		{
			x:   0,
			exp: 0,
		},
	}

	for _, tt := range tests {
		act := Reverse(tt.x)
		if act != tt.exp {
			t.Errorf("expected: %v, act: %v", tt.exp, act)
		}
	}
}
