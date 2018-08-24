package filter

import (
	"io"
)

type writeFilter struct {
	w          io.Writer
	translator func([]byte) ([]byte, error)
}

func (this *writeFilter) Write(p []byte) (int, error) {
	data, err := this.translator(p)
	if err != nil {
		return 0, err
	}
	return this.w.Write(data)
}

func NewWriteFilter(fd io.Writer, translator func([]byte) ([]byte, error)) io.Writer {
	return &writeFilter{
		w:          fd,
		translator: translator,
	}
}
