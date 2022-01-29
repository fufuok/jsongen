package internal

import (
	"bytes"
	"sync"
)

var (
	// 8 MiB
	defaultMaxSize = 8 << 20

	bufferPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(nil)
		},
	}
)

func Get() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

func Put(buf *bytes.Buffer) {
	Release(buf)
}

func Release(buf *bytes.Buffer) bool {
	if buf.Cap() > defaultMaxSize {
		return false
	}
	buf.Reset()
	bufferPool.Put(buf)
	return true
}
