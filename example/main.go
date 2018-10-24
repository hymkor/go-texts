package main

import (
	"fmt"

	"github.com/zetamatta/go-texts"
)

func main() {
	map1 := map[string]string{
		"A": "alpha",
		"B": "beta",
		"C": "gamma",
	}

	for _, key1 := range texts.SortedKeys(map1) {
		fmt.Printf("%s: %s\n", key1, map1[key1])
	}
}
