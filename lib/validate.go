package lib

import (
	"os"
	"strings"
)

func Validate() (string, string) {
	if len(os.Args) == 1 {
		return "zip lib name is required", ""
	}
	var fn = os.Args[1]
	if strings.Contains(fn, ".zip") == false {
		return "lib should be .zip", ""
	}
	return "", fn
}
