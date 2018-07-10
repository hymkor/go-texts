package mbcs

import (
	"bufio"
	"fmt"
	"io"
	"unicode/utf8"
)

// MbcsReader is the io.Reader to convert from mbcs to utf8
type MbcsReader struct {
	r       io.Reader
	sc      *bufio.Scanner
	buffer  []byte
	notutf8 bool
}

// Read is the method satisfy the requirements of io.Reader
func (this *MbcsReader) Read(p []byte) (int, error) {
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
		line := this.sc.Bytes()
		if !this.notutf8 && utf8.Valid(line) {
			this.buffer = append(this.buffer, line...)
		} else {
			text, err := AtoU(line)
			if err != nil {
				return 0, err
			}
			this.notutf8 = true
			this.buffer = append(this.buffer, []byte(text)...)
		}
		this.buffer = append(this.buffer, '\n')
	}
	for i := 0; i < len(p); i++ {
		p[i] = this.buffer[i]
	}
	this.buffer = this.buffer[len(p):]
	return len(p), nil
}

// NewReader returns reader object traslating from MBCS to UTF8
func NewReader(fd io.Reader) io.Reader {
	return &MbcsReader{
		r:  fd,
		sc: bufio.NewScanner(fd),
	}
}

// Reader is obsolete. Use NewReader()
func Reader(fd io.Reader, onError func(error, io.Writer) bool) io.ReadCloser {
	reader, writer := io.Pipe()
	go func() {
		sc := bufio.NewScanner(fd)
		defer writer.Close()
		notUtf8 := false
		for sc.Scan() {
			line := sc.Bytes()
			if !notUtf8 && utf8.Valid(line) {
				fmt.Fprintln(writer, string(line))
			} else {
				text, err := AtoU(line)
				if err != nil {
					if !onError(err, writer) {
						return
					}
				} else {
					notUtf8 = true
					fmt.Fprintln(writer, text)
				}
			}
		}
	}()
	return reader
}
