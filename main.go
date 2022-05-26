package main

import (
	"log"
	"netsuite-companion-document/lib"
)

func main() {

	v, fn := lib.Validate()
	if v != "" {
		log.Fatalf(v)
	}

	lib.CleanUp(fn)
	lib.Unzip(fn)
	lib.Move(fn)
	lib.Clear()
}
