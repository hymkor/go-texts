package main

import (
	"github.com/hymkor/go-texts/dos"
)

func main() {
	dos.System(`dir /w "C:\Program Files" & echo Done!`)

	dos.SystemWith("/V:ON", "echo with /V:ON,COMSPEC=!COMSPEC!")
	dos.SystemWith("/V:OFF", "echo with /V:OFF,COMSPEC=!COMSPEC!")
}
