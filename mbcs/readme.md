go-mbcs
=======

`go-mbcs` is the library for the programming language Go for Windows,
to convert characters between the current codepage string(ANSI) and UTF8

This library uses `MultiByteToWideChar` and `WideCharToMultiByte`
in `kernel32.dll`

* `AtoU` - Convert ANSI to UTF8
* `UtoAc` - Convert UTF8 to ANSI string without '\000'
* `UtoAz` - Convert UTF8 to ANSI string with '\000'
* `UtoA` - Compatible `UtoAz`

```
    package main

    import (
            "fmt"
            "os"
            "github.com/zetamatta/go-mbcs"
    )

    func main() {
            ansi, err := mbcs.UtoAc("UTF8文字列")
            if err != nil {
                    fmt.Fprintln(os.Stderr, err.Error())
                    return
            }

            utf8, err := mbcs.AtoU(ansi)
            if err != nil {
                    fmt.Fprintln(os.Stderr, err.Error())
                    return
            }
            fmt.Printf("Ok: %s\n", utf8)
    }
```

<!-- vim:set fenc=utf8: -->
