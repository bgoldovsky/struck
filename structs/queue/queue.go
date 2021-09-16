package main

import (
	"errors"
)

/*
	Queue

	Очередь работает по принципу FIFO (First in first out)
	При том, что интерфейс не накладывает ограничений на внутреннее устройство,
	очередь принято реализовывать таким образом, чтобы операции вставки и удаления элемента выполнялись за O(1)
	==============================
	I. Time: O(1)
	II. Space: O(1)
*/

var errEmpty = errors.New("queue empty error")

type node struct {
	data interface{}
	next *node
}

func newNode(data interface{}, next *node) *node {
	return &node{
		data: data,
		next: next,
	}
}

func NewQueue() *queue {
	return &queue{}
}

type queue struct {
	head *node
	tail *node
	size int
}

// Get Извлекает элемент из очереди
// Time O(1)
// Space O(1)
func (q *queue) Get() (interface{}, error) {
	if q.IsEmpty() {
		return 0, errEmpty
	}

	q.size--
	data := q.head.data
	q.head = q.head.next

	return data, nil
}

// Put Добавляет элемент в очередь
// Time O(1)
// Space O(1)
func (q *queue) Put(item int64) {
	if q.IsEmpty() {
		q.size++
		n := newNode(item, nil)
		q.tail = n
		q.head = n
		return
	}

	q.size++
	current := q.tail

	newQueueNode := newNode(item, nil)
	current.next = newQueueNode
	q.tail = newQueueNode
}

// Size Возвращает количество элементов в очереди
// Time O(1)
// Space O(1)
func (q *queue) Size() int {
	return q.size
}

// IsEmpty Возвращает признак пуста ли очередь
// Time O(1)
// Space O(1)
func (q *queue) IsEmpty() bool {
	return q.size == 0
}