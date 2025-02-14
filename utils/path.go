package utils

import (
	"errors"
	"os"
)

// PathExists checks if the directory exists
func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("there is a file with the same name")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// PathFileExists checks if the file exists
func PathFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
