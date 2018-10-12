package mbcs

import (
	"bytes"
	"io"
	"syscall"
	"unicode/utf8"

	"github.com/zetamatta/go-texts"
	"github.com/zetamatta/go-texts/filter"
)

type UTF16State uint

const (
	NotSet UTF16State = iota
	NotUTF16
	UTF16LE
	UTF16BE
)

// NewAutoDetectReader returns reader object traslating from MBCS/UTF8 to UTF8
func NewAutoDetectReader(fd io.Reader, cp uintptr) io.Reader {
	notutf8 := false
	utf16status := NotSet
	var utf16left []byte
	return filter.New(fd, func(line []byte) ([]byte, error) {
		if utf16status == NotSet {
			if line[0] == 0xFE && line[1] == 0xFF {
				utf16status = UTF16BE
				line = line[2:]
			} else if line[0] == 0xFF && line[1] == 0xFE {
				utf16status = UTF16LE
				line = line[2:]
			} else if pos := bytes.IndexByte(line, 0); pos >= 0 {
				if pos%2 == 0 {
					utf16status = UTF16BE
				} else {
					utf16status = UTF16LE
				}
			} else {
				utf16status = NotUTF16
			}
			//} else if utf16status == UTF16LE && line[0] == 0 {
			//	line = line[1:]
		}
		if utf16status != NotUTF16 {
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
				if utf16status == UTF16BE {
					w = (uint16(line[i]) << 8) | uint16(line[i+1])
				} else {
					w = uint16(line[i]) | (uint16(line[i+1]) << 8)
				}
				utf16s = append(utf16s, w)
			}
			utf8s := syscall.UTF16ToString(utf16s)
			return []byte(utf8s), nil
		}

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
