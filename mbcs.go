package mbcs

import (
	"syscall"
	"unsafe"
)

var kernel32 = syscall.NewLazyDLL("kernel32")
var multiByteToWideChar = kernel32.NewProc("MultiByteToWideChar")
var wideCharToMultiByte = kernel32.NewProc("WideCharToMultiByte")

func utoa(utf8 string) ([]byte, error) {
	utf16, err := syscall.UTF16FromString(utf8)
	if err != nil {
		return nil, err
	}
	size, _, _ := wideCharToMultiByte.Call(CP_THREAD_ACP, 0,
		uintptr(unsafe.Pointer(&utf16[0])),
		uintptr(len(utf16)),
		uintptr(0), 0, uintptr(0), uintptr(0))
	if size <= 0 {
		return nil, syscall.GetLastError()
	}
	mbcs := make([]byte, size)
	rc, _, _ := wideCharToMultiByte.Call(CP_THREAD_ACP, 0,
		uintptr(unsafe.Pointer(&utf16[0])),
		uintptr(len(utf16)),
		uintptr(unsafe.Pointer(&mbcs[0])), size, uintptr(0), uintptr(0))
	if rc == 0 {
		return nil, syscall.GetLastError()
	}
	return mbcs, nil
}

// UtoAz - Convert UTF8 to Ansi string with \0
func UtoAz(utf8 string) ([]byte, error) { return utoa(utf8) }

// UtoA - Convert UTF8 to Ansi string with \0 (for compatible)
func UtoA(utf8 string) ([]byte, error) { return utoa(utf8) }

// UtoAc - Convert UTF8 to Ansi string without \0 from UTF8 (chop \0)
func UtoAc(utf8 string) ([]byte, error) {
	ansi, err := utoa(utf8)
	if err != nil {
		return nil, err
	}
	if ansi[len(ansi)-1] == 0 {
		ansi = ansi[:len(ansi)-1]
	}
	return ansi, nil
}

// AtoU - Convert Ansi string to UTF8
func AtoU(mbcs []byte) (string, error) {
	if mbcs == nil || len(mbcs) <= 0 {
		return "", nil
	}
	size, _, _ := multiByteToWideChar.Call(CP_THREAD_ACP, 0,
		uintptr(unsafe.Pointer(&mbcs[0])),
		uintptr(len(mbcs)),
		uintptr(0), 0)
	if size <= 0 {
		return "", syscall.GetLastError()
	}
	utf16 := make([]uint16, size)
	rc, _, _ := multiByteToWideChar.Call(CP_THREAD_ACP, 0,
		uintptr(unsafe.Pointer(&mbcs[0])), uintptr(len(mbcs)),
		uintptr(unsafe.Pointer(&utf16[0])), size)
	if rc == 0 {
		return "", syscall.GetLastError()
	}
	return syscall.UTF16ToString(utf16), nil
}
