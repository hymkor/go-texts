package mbcs

import (
	"io"
	"unicode/utf8"

	"github.com/zetamatta/go-texts/filter"
)

// NewAutoDetectReader returns reader object traslating from MBCS/UTF8 to UTF8
func NewAutoDetectReader(fd io.Reader, cp uintptr) io.Reader {
	notutf8 := false
	return filter.New(fd, func(line []byte) ([]byte, error) {
		if !notutf8 && utf8.Valid(line) {
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
