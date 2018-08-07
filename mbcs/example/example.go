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

// vim:set fenc=utf8:
