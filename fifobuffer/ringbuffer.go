package fifobuffer

import (
	"container/ring"
	"sync"
	//"fmt"
)

// RingBuffer is a FIFO queue/buffer built on the ring container
type RingBuffer struct {
	read *ring.Ring
	write *ring.Ring
	lock *sync.Mutex
	length int
	capacity int
}

// NewRingBuffer create a new RingBuffer
func NewRingBuffer(capacity int) *RingBuffer {
	r := ring.New(capacity)
	rBuff := RingBuffer{read: r, write: r, lock: &sync.Mutex{}, capacity: capacity}
	return &rBuff
}

func (buffer *RingBuffer) Write(item interface{}) (interface{}, error) {
	buffer.lock.Lock()
	defer buffer.lock.Unlock()

	// if buffer.length == buffer.capacity {
	// 	return nil, fmt.Errorf("Buffer full")
	// }

	old := buffer.write.Value
	buffer.write.Value = item
	buffer.write = buffer.write.Next()

	buffer.length++

	return old, nil
}

func (buffer *RingBuffer) Read() (interface{}, error) {
	buffer.lock.Lock()
	defer buffer.lock.Unlock()

	value := buffer.read.Value
	buffer.read = buffer.read.Next()
	
	buffer.length--

	return value, nil
}