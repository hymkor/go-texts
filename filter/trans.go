package filter

import (
	"bytes"
	"io"

	"golang.org/x/text/transform"
)

type LineTransformer struct {
	bytes.Buffer
	Filter func([]byte) ([]byte, error)
}

func (this *LineTransformer) Transform(dst, src []byte, atEOF bool) (int, int, error) {
	this.Write(src)

	_src := make([]byte, this.Len())
	copy(_src, this.Bytes())

	backup := _src
	this.Reset()

	rollback := func() {
		this.Reset()
		this.Write(backup)
	}

	if !atEOF {
		pos := bytes.LastIndexByte(_src, '\n')
		if pos >= 0 {
			this.Write(_src[pos+1:])
			_src = _src[:pos+1]
		}
	}

	_dst, err := this.Filter(_src)
	if err != nil {
		rollback()
		return 0, 0, err
	}
	if len(dst) < len(_dst) {
		rollback()
		return 0, 0, transform.ErrShortDst
	}
	copy(dst, []byte(_dst))
	return len(_dst), len(src), nil
}

func NewLineFilter(r io.Reader, filter func([]byte) ([]byte, error)) io.Reader {
	return transform.NewReader(r, &LineTransformer{Filter: filter})
}
