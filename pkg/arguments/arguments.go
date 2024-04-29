package arguments

import (
	"fmt"
)

func CheckArguments(imagePath string) error {
	if imagePath == "" {
		return fmt.Errorf("imagePath is empty")
	}

	if !ValidImage(imagePath) {
		return fmt.Errorf("imagePath is not valid")
	}

	return nil
}
