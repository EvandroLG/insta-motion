package validator

import "os"

func ValidImage(imagePath string) bool {
	if imagePath == "" {
		return false
	}

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return false
	}

	if !isImage(imagePath) {
		return false
	}

	return true
}

func isImage(imagePath string) bool {
	file, err := os.Open(imagePath)

	if err != nil {
		return false
	}

	defer file.Close()

	buff := make([]byte, 512)

	if _, err := file.Read(buff); err != nil {
		return false
	}

	return true
}
