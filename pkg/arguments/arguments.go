package arguments

import (
	"fmt"
)

func CheckArguments(imagePath string, effect string) error {
	if imagePath == "" {
		return fmt.Errorf("imagePath is empty")
	}

	if effect == "" {
		return fmt.Errorf("effect is empty")
	}

	if !ValidImage(imagePath) {
		return fmt.Errorf("imagePath is not valid")
	}

	return nil
}
