package persistence

import (
	"os"
)

const prefix = "storage/"

func readFile(fileName string) (string, error) {
	fileName = prefix + fileName
	bytes, err := os.ReadFile(fileName)
	return string(bytes), err
}

func writeFile(fileName, content string) error {
	fileName = prefix + fileName
	data := []byte(content)
	return os.WriteFile(fileName, data, 0644) // 0644 sets read/write for owner, read-only for othersd
}
