package utils

import (
	"os"
)

func DirectoryExists(dirPath string) bool {
    info, err := os.Stat(dirPath)
    return !os.IsNotExist(err) && info.IsDir()
}
