package core

import (
	"errors"
	"fmt"
	"io"
)

// type Queue[T any] interface {
// 	Enqueue(T) error
// 	Dequeue() (T, error)
// }

type message[T any] struct {
	Val  T
	prev *message[T]
	next *message[T]
}

type queue[T any] struct {
	new      bool
	capacity int
	msgCount int
	head     *message[T]
	tail     *message[T]
}

func (q *queue[T]) validQueue() {
	if q == nil || !q.new {
		panic("queue not initialized properly")
	}
}

func NewQueue[T any](Capacity int) *queue[T] {
	return &queue[T]{new: true, capacity: Capacity}
}

func (q *queue[T]) Enqueue(Val T) error {
	q.validQueue()
	if q.msgCount == q.capacity {
		return errors.New("max capacity reached")
	}
	m := &message[T]{Val: Val}
	if q.head == nil && q.tail == nil {
		q.head, q.tail = m, m
	} else {
		m.prev = q.tail
		q.tail.next = m
		q.tail = m
	}
	q.msgCount += 1
	return nil
}

func (q *queue[T]) Dequeue() (T, error) {
	q.validQueue()
	var msgVal T
	if q.msgCount == 0 {
		return msgVal, errors.New("queue is empty")
	}
	if q.msgCount == 1 {
		msgVal = q.head.Val
		q.head = nil
		q.tail = nil
	} else {
		msgVal = q.head.Val
		q.head = q.head.next
	}
	q.msgCount -= 1
	return msgVal, nil
}

func (q *queue[T]) GetOldest() T {
	return q.head.Val
}

func (q *queue[T]) GetNewest() T {
	return q.tail.Val
}

func (q *queue[T]) Traverse(w io.Writer) {
	q.validQueue()
	tmp := q.head
	for tmp != nil {
		fmt.Fprint(w, tmp.Val)
		tmp = tmp.next
	}
}
