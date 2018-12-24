package preprocessor

import (
	"bytes"
	"io"
)

type preprocessor struct {
	f       func() ([]byte, error)
	buffer  [2]bytes.Buffer
	current uint
}

func (pp *preprocessor) Read(b []byte) (int, error) {
	buf := &pp.buffer[pp.current]
	for buf.Len() < len(b) {
		tmp, err := pp.f()
		if err != nil {
			end := buf.Bytes()
			copy(b, end)
			buf.Reset()
			return len(end), err
		}
		buf.Write(tmp)
	}
	end := buf.Bytes()
	copy(b, end)

	pp.current ^= 1
	next := &pp.buffer[pp.current]
	next.Reset()
	next.Write(end[len(b):])
	return len(b), nil
}

// New returns new preprocessor instance.
func New(f func() ([]byte, error)) io.Reader {
	return &preprocessor{f: f}
}
