package ms

func (m *MemStore) Write(b []byte) (n int, err error) {

	temp := make([]byte, len(m.buf)+len(b))
	copy(temp, m.buf)
	copy(temp[len(m.buf):], b)
	m.buf = temp
	return len(b), nil
}
