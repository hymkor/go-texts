package mbcs

import (
	"bytes"
	"io"
	"syscall"
	"unicode/utf8"

	"github.com/zetamatta/go-texts/filter"
)

const (
	NotSet = 0
)

type UTF16State uint

const (
	NotUTF16 UTF16State = iota + 1
	UTF16LE
	UTF16BE
)

type UTF8State int

const (
	NotUTF8 UTF8State = iota + 1
	UTF8
)

var BOM8 = []byte{0xEF, 0xBB, 0xBF}

// NewAutoDetectReader returns reader object traslating from MBCS/UTF8 to UTF8
func NewAutoDetectReader(fd io.Reader, cp uintptr) io.Reader {
	var utf16state UTF16State = NotSet
	var utf8state UTF8State = NotSet
	var utf16left []byte
	return filter.New(fd, func(line []byte) ([]byte, error) {
		if utf16state == NotSet {
			if len(line) >= 2 && line[0] == 0xFE && line[1] == 0xFF {
				utf16state = UTF16BE
				line = line[2:]
			} else if len(line) >= 2 && line[0] == 0xFF && line[1] == 0xFE {
				utf16state = UTF16LE
				line = line[2:]
			} else if pos := bytes.IndexByte(line, 0); pos >= 0 {
				if pos%2 == 0 {
					utf16state = UTF16BE
				} else {
					utf16state = UTF16LE
				}
			} else {
				utf16state = NotUTF16
			}
			//} else if utf16state == UTF16LE && line[0] == 0 {
			//	line = line[1:]
		}
		if utf16state != NotUTF16 {
			if utf16left != nil && len(utf16left) > 0 {
				tmp := append(utf16left, line...)
				line = tmp
				utf16left = nil
			}
			if len(line)%2 == 1 {
				utf16left = []byte{line[len(line)-1]}
				line = line[:len(line)-1]
			}
			utf16s := make([]uint16, 0, len(line)/2+1)
			for i := 0; i+1 < len(line); i += 2 {
				var w uint16
				if utf16state == UTF16BE {
					w = (uint16(line[i]) << 8) | uint16(line[i+1])
				} else {
					w = uint16(line[i]) | (uint16(line[i+1]) << 8)
				}
				utf16s = append(utf16s, w)
			}
			utf8s := syscall.UTF16ToString(utf16s)
			return []byte(utf8s), nil
		}
		if utf8state == NotSet {
			if len(line) >= 3 &&
				line[0] == BOM8[0] &&
				line[1] == BOM8[1] &&
				line[2] == BOM8[2] {

				line = line[3:]
				utf8state = UTF8
			} else if !utf8.Valid(line) {
				utf8state = NotUTF8
			}
		}
		if utf8state == NotUTF8 {
			text, err := AtoU(line, cp)
			if err != nil {
				return nil, err
			}
			return []byte(text), nil
		}
		return []byte(line), nil
	})
}

// NewAtoUReader returns new reader translate from mbcs to utf8.
func NewAtoUReader(r io.Reader, cp uintptr) io.Reader {
	return filter.New(r, func(line []byte) ([]byte, error) {
		text, err := AtoU(line, cp)
		return []byte(text), err
	})
}
