package arguments_test

import (
	"testing"

	"github.com/evandrolg/insta-motion/pkg/arguments"
)

func TestCheckArguments(t *testing.T) {
	err := arguments.CheckArguments("", "")

	if err == nil {
		t.Errorf("Expected error when image is an empty string, got nil")
	}

	err = arguments.CheckArguments("invalid_path", "")

	if err == nil {
		t.Errorf("Expected error when path is invalid, got nil")
	}

	err = arguments.CheckArguments("test.jpg", "")

	if err == nil {
		t.Errorf("Expected error when image does not exist, got nil")
	}

	err = arguments.CheckArguments("./test.jpeg", "")

	if err != nil {
		t.Errorf("Expected no error when image is valid, got %v", err)
	}

	err = arguments.CheckArguments("./test.jpeg", "invalid_effect")

	if err == nil {
		t.Errorf("Expected error when effect is invalid, got nil")
	}

	err = arguments.CheckArguments("./test.jpeg", "zoom-in")

	if err != nil {
		t.Errorf("Expected no error when effect is valid, got %v", err)
	}

	err = arguments.CheckArguments("./test.jpeg", "zoom-out")

	if err != nil {
		t.Errorf("Expected no error when effect is valid, got %v", err)
	}
}
