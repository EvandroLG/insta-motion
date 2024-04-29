package main

import (
	"flag"
	"fmt"

	"github.com/evandrolg/insta-motion/pkg/arguments"
	"github.com/evandrolg/insta-motion/pkg/converter"
)

func main() {
	var imagePath string
	flag.StringVar(&imagePath, "image", "", "Path to the image")
	flag.Parse()

	if err := arguments.CheckArguments(imagePath); err != nil {
		fmt.Println(err)
	}

	outputVideo, err := converter.ConvertToVideo(imagePath)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Video created: %s\n", outputVideo)
}
