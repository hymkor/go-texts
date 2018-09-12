package mbcs

import (
	"syscall"
	"unsafe"
)

var kernel32 = syscall.NewLazyDLL("kernel32")
var multiByteToWideChar = kernel32.NewProc("MultiByteToWideChar")
var wideCharToMultiByte = kernel32.NewProc("WideCharToMultiByte")
var getConsoleCp = kernel32.NewProc("GetConsoleCP")

const ACP = CP_ACP
const THREAD_ACP = CP_THREAD_ACP

// ConsoleCP returns Codepage number of Console.
func ConsoleCP() uintptr {
	cp, _, _ := getConsoleCp.Call()
	return cp
}

// UtoA converts from UTF8 to ANSI(codepage string).
// cp : codepage such as ACP , THREAD_ACP or ConsoleCP()
// chopzero : if it is true trim last \0.
func UtoA(utf8 string, cp uintptr, chopzero bool) ([]byte, error) {
	utf16, err := syscall.UTF16FromString(utf8)
	if err != nil {
		return nil, err
	}
	size, _, _ := wideCharToMultiByte.Call(cp, 0,
		uintptr(unsafe.Pointer(&utf16[0])),
		uintptr(len(utf16)),
		uintptr(0), 0, uintptr(0), uintptr(0))
	if size <= 0 {
		return nil, syscall.GetLastError()
	}
	mbcs := make([]byte, size)
	rc, _, _ := wideCharToMultiByte.Call(cp, 0,
		uintptr(unsafe.Pointer(&utf16[0])),
		uintptr(len(utf16)),
		uintptr(unsafe.Pointer(&mbcs[0])), size, uintptr(0), uintptr(0))
	if rc == 0 {
		return nil, syscall.GetLastError()
	}
	if chopzero && len(mbcs) > 0 && mbcs[len(mbcs)-1] == 0 {
		mbcs = mbcs[:len(mbcs)-1]
	}
	return mbcs, nil
}

// AtoU - Convert ANS(codepage string) to UTF8
// cp : codepage such as ACP , THREAD_ACP or ConsoleCP()
func AtoU(mbcs []byte, cp uintptr) (string, error) {
	if mbcs == nil || len(mbcs) <= 0 {
		return "", nil
	}
	size, _, _ := multiByteToWideChar.Call(cp, 0,
		uintptr(unsafe.Pointer(&mbcs[0])),
		uintptr(len(mbcs)),
		uintptr(0), 0)
	if size <= 0 {
		return "", syscall.GetLastError()
	}
	utf16 := make([]uint16, size)
	rc, _, _ := multiByteToWideChar.Call(cp, 0,
		uintptr(unsafe.Pointer(&mbcs[0])), uintptr(len(mbcs)),
		uintptr(unsafe.Pointer(&utf16[0])), size)
	if rc == 0 {
		return "", syscall.GetLastError()
	}
	return syscall.UTF16ToString(utf16), nil
}
