package lib

import (
	"os"
	"strings"
)

func CleanUp(fn string) {
	err := os.RemoveAll("./website/docs/" + strings.Replace(fn, ".zip", "", 1))
	if err != nil {
		return
	}
}

func Clear() {
	err := os.RemoveAll("output")
	if err != nil {
		return
	}
}
