//go:build ignore
// +build ignore

// Test case: source stream's last line has no LF.

package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/zetamatta/go-texts/mbcs"
)

func main() {
	reader := mbcs.NewAtoUReader(strings.NewReader("1stLine\nNoLfLine"),
		mbcs.ConsoleCP())
	sc := bufio.NewScanner(reader)
	fmt.Println("--- mbcs.NewAtoUReader ---")
	for sc.Scan() {
		fmt.Println(sc.Text())
	}

	fmt.Println("--- mbcs.NewAutoDetectReader ---")
	reader = mbcs.NewAutoDetectReader(strings.NewReader("1stLine\nNoLfLine"),
		mbcs.ConsoleCP())
	sc = bufio.NewScanner(reader)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
}
