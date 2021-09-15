package main

import (
	"testing"
)

func TestStack_Push(t *testing.T) {
	s := New(4)

	exp := int64(5)
	s.Push(exp)

	if len(s.Store) != 1 {
		t.Errorf("expected: %v, act: %v", exp, 1)
	}

	act := s.Store[0]

	if act != exp {
		t.Errorf("expected: %v, act: %v", exp, act)
	}
}

func TestStack_Pop(t *testing.T) {
	s := New(4)

	exp := int64(5)
	s.Push(exp)

	act := s.Pop()

	if act != exp {
		t.Errorf("expected: %v, act: %v", exp, act)
	}
}

func TestStack_Pop_Nil(t *testing.T) {
	s := New(4)

	s.Push(5)

	act := s.Pop()
	act = s.Pop()

	if act != -1 {
		t.Errorf("expected: %v, act: %v", nil, act)
	}
}

func TestStack_IsEmpty(t *testing.T) {
	s := New(4)

	s.Push(5)

	act := s.IsEmpty()
	if act {
		t.Errorf("expected false, act: %v", act)
	}

	_ = s.Pop()

	act = s.IsEmpty()
	if !act {
		t.Errorf("expected true, act: %v", act)
	}
}