package persistence

import (
	"os"
	"time"
)

func getFileTime(fileName string) (time.Time, error) {
	fileName = prefix + fileName
	fileInfo, err := os.Stat(fileName)
	return fileInfo.ModTime(), err
}

func GetAccessTokenTime() (time.Time, error) {
	return getFileTime("access.token")
}

func GetRefreshTokenTime() (time.Time, error) {
	return getFileTime("refresh.token")
}
