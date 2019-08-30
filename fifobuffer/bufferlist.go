package fifobuffer

import (
	"sync"
	"time"
)

type Metric struct {
	startTime time.Time
	endTime time.Time
	metric metricData
}

type metricData struct {
	metricName string
	values []uint64
}

type node struct {
	data interface{}
	next *node
}

type ListBuffer struct {
	head *node
	tail *node
	size int
	capacity int
	lock *sync.Mutex
}

func NewListBuffer(capacity int) *ListBuffer {
	b := &ListBuffer{lock: &sync.Mutex{}, capacity: capacity}
	return b
}

func (b *ListBuffer) Len() int {
	b.lock.Lock()
	defer b.lock.Unlock()

	return b.size
}

func (b *ListBuffer) Push(item interface{}) interface{} {
	b.lock.Lock()
	defer b.lock.Unlock()

	n := &node{data: item}
	var popped interface{}

	if b.size + 1 >= b.capacity {
		popped = b.popInternal()
	}

	if b.tail == nil {
		b.tail = n
		b.head = n
	} else {
		b.tail.next = n
		b.tail = n
	}

	b.size++

	return popped
}


func (b *ListBuffer) popInternal() interface{} {
	if b.head == nil {
		return nil
	}

	n := b.head
	b.head = n.next

	if b.head == nil {
		b.tail = nil
	}

	b.size--

	return n.data
}

func (b *ListBuffer) Pop() interface{} {
	b.lock.Lock()
	defer b.lock.Unlock()

	return b.popInternal()
}
