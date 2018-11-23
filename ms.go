package ms

import (
	"fmt"
	"io"
)

//Memstore operation
type MemStoreOperation interface {
	size() int
	IsEmpty() bool
	Write(b []byte) (n int, err error)
	io.Reader
	fmt.Stringer
}

// MemsStore implementation
type MemStore struct {
	buf []byte
}

func (m MemStore) Size() int {
	return len(m.buf)
}

func (m MemStore) IsEmpty() bool {
	return 0 == m.Size()
}
