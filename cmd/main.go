package main

import (
	"flag"
	"fmt"

	"github.com/evandrolg/insta-motion/pkg/arguments"
)

func main() {
	var imagePath, effect string
	flag.StringVar(&imagePath, "image", "", "Path to the image")
	flag.StringVar(&effect, "effect", "", "Effect to apply to the image")
	flag.Parse()

	if err := arguments.CheckArguments(imagePath, effect); err != nil {
		fmt.Errorf("Error: %v", err)
	}

	fmt.Println("Everything is ok")
}
