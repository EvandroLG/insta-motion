package main

import (
	"os"

	"github.com/evandrolg/insta-motion/pkg/validator"
)

func main() {
	args := os.Args

	if len(args) == 0 {
		println("Image path is required")
		os.Exit(1)
	}

	imagePath := args[1]

	if !validator.ValidImage(imagePath) {
		println("Image is not valid")
		os.Exit(1)
	}

	println("Image is valid")
}
