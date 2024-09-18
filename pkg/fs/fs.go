package fs

import (
	"os"
)

func FileExists(filePath string) bool {
    _, err := os.Stat(filePath)
    return !os.IsNotExist(err)
}

func DirectoryExists(dirPath string) bool {
	info, err := os.Stat(dirPath)
	return !os.IsNotExist(err) && info.IsDir()
}
