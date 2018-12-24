package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/zetamatta/go-texts/preprocessor"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	count := 0
	lnumFilter := func() ([]byte, error) {
		line, err := br.ReadString('\n')
		if err != nil {
			return []byte{}, err
		}
		count++
		return []byte(fmt.Sprintf("%d: %s", count, line)), nil
	}

	sc := bufio.NewScanner(preprocessor.New(lnumFilter))
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
