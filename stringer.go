package ms

import (
	"fmt"
)

func (m MemStore) String() string {

	return fmt.Sprintf("%T (size:%v,  buf:%v)", m, m.Size(), m.buf)

}
