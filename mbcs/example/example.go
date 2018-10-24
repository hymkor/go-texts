package main

import (
	"fmt"
	"os"

	"github.com/zetamatta/go-texts/mbcs"
)

func main1() error {
	ansi, err := mbcs.UtoA("UTF8文字列", mbcs.ConsoleCP(), true)
	if err != nil {
		return err
	}

	utf8, err := mbcs.AtoU(ansi, mbcs.ConsoleCP())
	if err != nil {
		return err
	}
	fmt.Printf("Ok: %s\n", utf8)
	return nil
}

func main() {
	if err := main1(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// vim:set fenc=utf8:
