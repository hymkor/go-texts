go-texts/mbcs
=============

`go-mbcs` is the library for the programming language Go for Windows,
to convert characters between the current codepage string(ANSI) and UTF8

This library uses `MultiByteToWideChar` and `WideCharToMultiByte`
in `kernel32.dll`

* `AtoU` - Convert ANSI(Codepage String) to UTF8
* `UtoA` - Convert UTF8 to ANSI(Codepage String)
* `NewAutoDetectReader` - reader traslating from ANSI/UTF8 to UTF8
* `ACP` - ANSI Codepage
* `THREAD_ACP` - ANSI Codepage of the current thread.
* `ConsoleCP` - Returns ANSI Codepage of the current console.

```
package main

import (
	"fmt"
	"github.com/zetamatta/go-texts/mbcs"
	"os"
)

func main() {
	ansi, err := mbcs.UtoA("UTF8文字列",mbcs.ACP,true)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	utf8, err := mbcs.AtoU(ansi,mbcs.ACP)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	fmt.Printf("Ok: %s\n", utf8)
}
```

<!-- vim:set fenc=utf8: -->
