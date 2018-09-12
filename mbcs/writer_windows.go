package mbcs

import (
	"bytes"
	"io"

	"github.com/zetamatta/go-texts/filter"
)

var lf = []byte{'\n'}
var crlf = []byte{'\r', '\n'}

func NewWriter(fd io.Writer, cp uintptr) io.Writer {
	return filter.NewWriter(fd, func(source []byte) ([]byte, error) {
		bin, err := UtoA(string(source), cp, true)
		if err != nil {
			return bin, err
		}
		bin = bytes.Replace(bin, crlf, lf, -1)
		bin = bytes.Replace(bin, lf, crlf, -1)
		return bin, nil
	})
}
