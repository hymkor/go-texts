//go:build !windows
// +build !windows

package mbcs

import "io"

func NewAutoDetectReader(fd io.Reader, _ uintptr) io.Reader {
	return fd
}

func NewAtoUReader(fd io.Reader, _ uintptr) io.Reader {
	return fd
}
