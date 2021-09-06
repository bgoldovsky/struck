package deque

import (
	"errors"
)

/*
	Deque (Double ended queue)
	Очередь с двумя концами, двусторонняя очередь
	==============================
	I. Time:
		Constant, O(1)
	==============================
	II. Space
		Linear for init, O(n)
		Constant for all operations, O(1)
*/

type Deque interface {
	PushBack(item int64) error
	PopBack() (int64, error)
	PushFront(item int64) error
	PopFront() (int64, error)
	Size() int
}

type deque struct {
	buf   []int64
	max   int
	size  int
	front int
	back  int
}

func New(maxSize int) *deque {
	return &deque{
		buf: make([]int64, maxSize),
		max: maxSize,
	}
}

// Size Возвращает размер очереди
// Time O(1)
// Space O(1)
func (d *deque) Size() int {
	return d.size
}

// PushFront Добавляет элемент в начало очереди
// Time O(1)
// Space O(1)
func (d *deque) PushFront(item int64) error {
	if d.size == d.max {
		return errors.New("push front: max deque size")
	}

	d.front = d.prev(d.front)
	d.buf[d.front] = item
	d.size++

	return nil
}

// PushBack Добавляет элемент в конец очереди
// Time O(1)
// Space O(1)
func (d *deque) PushBack(item int64) error {
	if d.size == d.max {
		return errors.New("push back: max deque size")
	}

	d.buf[d.back] = item
	d.back = d.next(d.back)
	d.size++

	return nil
}

// PopFront Извлекает элемент из конца очереди
// Time O(1)
// Space O(1)
func (d *deque) PopFront() (int64, error) {
	if d.size == 0 {
		return 0, errors.New("exp front: empty deque error")
	}

	ret := d.buf[d.front]
	d.front = d.next(d.front)
	d.size--

	return ret, nil
}

// PopBack Извлекает элемент из начала очереди
// Time O(1)
// Space O(1)
func (d *deque) PopBack() (int64, error) {
	if d.size == 0 {
		return 0, errors.New("exp back: empty deque error")
	}

	d.back = d.prev(d.back)
	ret := d.buf[d.back]

	d.size--

	return ret, nil
}

func (d *deque) prev(pos int) int {
	pos--

	if pos < 0 {
		pos = d.max - 1
	}

	return pos
}

func (d *deque) next(pos int) int {
	pos++

	if pos >= d.max {
		pos = 0
	}

	return pos
}