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

var BOM8 = []byte{0xEF, 0xBB, 0xBF}

func anyToUtf8(line []byte, cp uintptr) (string, error) {
	if len(line) >= 3 &&
		line[0] == BOM8[0] &&
		line[1] == BOM8[1] &&
		line[2] == BOM8[2] {

		return string(line[3:]), nil
	}
	if utf8.Valid(line) {
		return string(line), nil
	}
	text, err := AtoU(line, cp)
	if err != nil {
		return "", err
	}
	return text, nil
}

// NewAutoDetectReader returns reader object traslating from MBCS/UTF8 to UTF8
func NewAutoDetectReader(fd io.Reader, cp uintptr) io.Reader {
	var utf16state UTF16State = NotSet
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
		var buffer bytes.Buffer
		for {
			endl := bytes.IndexByte(line, '\n')
			if endl < 0 {
				text, err := anyToUtf8(line, cp)
				if err != nil {
					return nil, err
				}
				buffer.WriteString(text)
				return buffer.Bytes(), nil
			}
			text, err := anyToUtf8(line[:endl+1], cp)
			if err != nil {
				return nil, err
			}
			buffer.WriteString(text)
			line = line[endl+1:]
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
