package daakia

import (
	"os"
	"strings"
)

func Clean(path string) error {
	return os.RemoveAll(path)
}

func Mkdir(path string) error {
	return os.MkdirAll(path, os.FileMode(0755))
}

func MkFile(dir string, filename string) (*os.File, error) {
	err := Mkdir(dir)
	if err != nil {
		return nil, err
	}
	return os.Create(dir + "/" + strings.ToLower(filename))
}
