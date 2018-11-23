package ms

func (m MemStore) Read(b []byte) (n int, err error) {
	n := copy(b, m.buf)

	return n, nil
}
