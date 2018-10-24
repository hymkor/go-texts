package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/zetamatta/go-texts/mbcs"
)

func main() {
	sc := bufio.NewScanner(mbcs.NewAutoDetectReader(os.Stdin, mbcs.ConsoleCP()))
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
