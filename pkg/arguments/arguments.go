package arguments

import (
	"fmt"
)

var Effects = map[string]bool{
	"zoom-in":  true,
	"zoom-out": true,
}

func CheckArguments(imagePath string, effect string) error {
	if imagePath == "" {
		return fmt.Errorf("imagePath is empty")
	}

	fmt.Println(effect)

	if effect != "" && !Effects[effect] {
		return fmt.Errorf("effect is not valid")
	}

	if !ValidImage(imagePath) {
		return fmt.Errorf("imagePath is not valid")
	}

	return nil
}
