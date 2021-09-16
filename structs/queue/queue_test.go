package main

import (
	"errors"
	"testing"
)

func TestQueue_Put(t *testing.T) {
	exp := int64(123)
	q := NewQueue()

	q.Put(exp)

	if q.tail.data != exp {
		t.Errorf("expected: %d, act: %d", exp, q.tail.data)
	}

	if q.tail.next != nil {
		t.Errorf("expected: nil, act: %d", q.tail.data)
	}
}

func TestQueue_Get(t *testing.T) {
	exp := int64(123)
	q := NewQueue()

	q.Put(exp)

	act, err := q.Get()
	if err != nil {
		t.Errorf("expected: nil, act: %d", err)
	}

	if act != exp {
		t.Errorf("expected: %d, act: %d", exp, act)
	}
}

func TestQueue_Get_Error(t *testing.T) {
	q := NewQueue()

	q.Put(123)

	_, err := q.Get()
	_, err = q.Get()
	if err == nil {
		t.Errorf("expected: nil, act: %d", err)
	}

	if !errors.Is(err, errEmpty) {
		t.Errorf("expected: %d, act: %d", errEmpty, err)
	}
}

func TestQueue_Size(t *testing.T) {
	q := NewQueue()

	act := q.Size()
	if act != 0 {
		t.Errorf("expected: %d, act: %d", 0, act)
	}

	q.Put(123)
	q.Put(234)

	act = q.Size()
	if act != 2 {
		t.Errorf("expected: %d, act: %d", 2, act)
	}

	_, _ = q.Get()
	act = q.Size()
	if act != 1 {
		t.Errorf("expected: %d, act: %d", 1, act)
	}
}
