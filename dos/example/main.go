package main

import (
	"github.com/zetamatta/go-texts/dos"
)

func main() {
	dos.System(`dir /w "C:\Program Files" & echo Done!`)
}
