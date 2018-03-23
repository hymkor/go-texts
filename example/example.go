package main

import (
	"fmt"
	"github.com/zetamatta/go-mbcs"
	"os"
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

// vim:set fenc=utf8:
