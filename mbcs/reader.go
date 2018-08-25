package mbcs

import (
	"bytes"
	"io"
	"unicode/utf8"

	"golang.org/x/text/transform"

	"github.com/zetamatta/go-texts"
	"github.com/zetamatta/go-texts/filter"
)

// NewAutoDetectReader returns reader object traslating from MBCS/UTF8 to UTF8
func NewAutoDetectReader(fd io.Reader, cp uintptr) io.Reader {
	notutf8 := false
	return filter.New(fd, func(line []byte) ([]byte, error) {
		if !notutf8 && utf8.Valid(line) {
			line = bytes.Replace(line, texts.ByteOrderMark, []byte{}, -1)
			return line, nil
		} else {
			text, err := AtoU(line, cp)
			if err != nil {
				return nil, err
			}
			notutf8 = true
			return []byte(text), nil
		}
	})
}

// NewAtoUReader returns new reader translate from mbcs to utf8.
func NewAtoUReader(r io.Reader, cp uintptr) io.Reader {
	return filter.New(r, func(line []byte) ([]byte, error) {
		text, err := AtoU(line, cp)
		return []byte(text), err
	})
}

// NewA2UReader returns new reader translate from mbcs to utf8 with transfomer interface. This version is slower than NewAtoUReader but few allocation.
func NewA2UReader(r io.Reader, cp uintptr) io.Reader {
	return transform.NewReader(r, &filter.LineTransformer{
		Filter: func(line []byte) ([]byte, error) {
			text, err := AtoU(line, cp)
			return []byte(text), err
		},
	})
}
