package mbcs

import (
	"bufio"
	"fmt"
	"io"
	"unicode/utf8"
)

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
