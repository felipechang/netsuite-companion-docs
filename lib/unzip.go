package lib

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(path string) {
	archive, err := zip.OpenReader(path)
	if err != nil {
		panic(err)
	}
	defer func(archive *zip.ReadCloser) {
		err = archive.Close()
		if err != nil {
			panic(err)
		}
	}(archive)
	for _, f := range archive.File {
		filePath := filepath.Join("output", f.Name)
		if !strings.HasPrefix(filePath, filepath.Clean("output")+string(os.PathSeparator)) {
			return
		}
		if f.FileInfo().IsDir() {
			err = os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				panic(err)
			}
			continue
		}
		if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}
		if strings.Contains(f.Name, ".md") == false {
			continue
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}
		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}
		if _, err = io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}
		err = dstFile.Close()
		if err != nil {
			return
		}
		err = fileInArchive.Close()
		if err != nil {
			return
		}
		if strings.Contains(filePath, "README.md") {
			sfn := strings.Split(filePath, "\\")
			nfn := sfn[len(sfn)-2] + ".md"
			err := os.Rename(filePath, strings.Replace(filePath, "README.md", nfn, 1))
			if err != nil {
				return
			}
		}
	}
}
