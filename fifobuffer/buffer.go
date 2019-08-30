package fifobuffer

import (
	"fmt"
	"sync"
)

type Buffer struct {
	Data []interface{}
	capacity int
	lock *sync.Mutex
}


func NewBuffer(capacity int) *Buffer {
	return &Buffer{Data: make([]interface{}, 0), capacity: capacity, lock: &sync.Mutex{}}
}

func (b *Buffer) Add(entry interface{}) {
	b.lock.Lock()
	defer b.lock.Unlock()

	// if at or over capacity drop the oldest entry
	if (len(b.Data) >= b.capacity) {
		b.Data = b.Data[1:]
	}

	b.Data = append(b.Data, entry)
}

func (b *Buffer) Pop() (interface{}, error) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if (len(b.Data) != 0) {
		value := b.Data[:1]
		b.Data = b.Data[1:]
		return value, nil
	}

	return nil, fmt.Errorf("Buffer is empty")
}