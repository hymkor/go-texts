package mbcs

import (
	"bufio"
	"io"
)

// Filter is the converter for io.Reader
type Filter struct {
	r          io.Reader
	br         *bufio.Reader
	buffer     []byte
	translator func([]byte) ([]byte, error)
}

// Read is the method satisfy the requirements of io.Reader
func (this *Filter) Read(p []byte) (int, error) {
	for len(this.buffer) < len(p) {
		line, err1 := this.br.ReadBytes('\n')
		line, err2 := this.translator(line)
		this.buffer = append(this.buffer, line...)
		if err1 != nil || err2 != nil {
			result := copy(p, this.buffer)
			var err error
			if err1 != nil {
				err = err1
			} else {
				err = err2
			}
			return result, err
		}
	}
	copy(p, this.buffer)
	this.buffer = this.buffer[len(p):]
	return len(p), nil
}

// NewFilter is the constructor for Filter
func NewFilter(fd io.Reader, translator func([]byte) ([]byte, error)) io.Reader {
	return &Filter{
		r:          fd,
		br:         bufio.NewReader(fd),
		translator: translator,
	}
}
