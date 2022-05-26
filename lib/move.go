package lib

import (
	"os"
	"strings"
)

func Move(fn string) {
	oldPath := "./output/FileCabinet/SuiteScripts"
	newPath := "./website/docs/" + strings.Replace(fn, ".zip", "", 1)
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return
	}
}
