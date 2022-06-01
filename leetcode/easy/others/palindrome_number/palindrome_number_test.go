package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		x   int
		exp bool
	}{
		{
			x:   121,
			exp: true,
		},
		{
			x:   -121,
			exp: false,
		},
		{
			x:   10,
			exp: false,
		},
		{
			x:   -1011,
			exp: false,
		},
	}

	for _, tt := range tests {
		act := IsPalindrome(tt.x)
		if act != tt.exp {
			t.Errorf("expected: %v, act: %v", tt.exp, act)
		}
	}
}
