package mbcs

import (
	"bufio"
	"io"
)

// Filter is the converter for io.Reader
type Filter struct {
	r          io.Reader
	sc         *bufio.Scanner
	buffer     []byte
	translator func([]byte) ([]byte, error)
}

// Read is the method satisfy the requirements of io.Reader
func (this *Filter) Read(p []byte) (int, error) {
	if this.r == nil {
		if err := this.sc.Err(); err != nil {
			return 0, err
		}
		return 0, io.EOF
	}
	for len(this.buffer) < len(p) {
		if !this.sc.Scan() {
			if len(this.buffer) <= 0 {
				this.r = nil
				if err := this.sc.Err(); err != nil {
					return 0, err
				}
				return 0, io.EOF
			}
			for i := 0; i < len(this.buffer); i++ {
				p[i] = this.buffer[i]
			}
			this.r = nil
			return len(this.buffer), nil
		}
		line, err := this.translator(this.sc.Bytes())
		if err != nil {
			return 0, err
		}
		this.buffer = append(this.buffer, line...)
		this.buffer = append(this.buffer, '\n')
	}
	for i := 0; i < len(p); i++ {
		p[i] = this.buffer[i]
	}
	this.buffer = this.buffer[len(p):]
	return len(p), nil
}

// NewFilter is the constructor for Filter
func NewFilter(fd io.Reader, translator func([]byte) ([]byte, error)) io.Reader {
	return &Filter{
		r:          fd,
		sc:         bufio.NewScanner(fd),
		translator: translator,
	}
}
