package main

import (
	"os"
	"strings"
	"path/filepath"
)

type folderData []string
type fileExtensions map[string]string


func CheckPathType(path string) (int, error) {
	info, err := os.Stat(path)
	if err != nil {
		return -1, err
	}

	switch mode := info.Mode(); {
	case mode.IsDir():
		return 1, nil
	case mode.IsRegular():
		return 2, nil
	}
	return 0, nil

}

/*
func CheckSuffix(path string) bool {
	if !strings.HasSuffix(path, "/") || filepath.Ext == "" {
		return false
	}
	return true
}
*/

func (fd *folderData) ParseFolder(location string, extensions fileExtensions) error{
	if !strings.HasSuffix(location, "/") {
		location = location + "/"
	}

	f, err := os.Open(location)
	if err != nil {
		return err
	}
	fileInfo, err := f.Readdir(-1)
	defer f.Close()
	if err != nil {
		return err
	}

	for _, file := range fileInfo {
		for _, ext := range extensions {
			if filepath.Ext(file.Name()) == ext {
				*fd = append(*fd, location + file.Name())
			}

		}
	}
	return nil
}

