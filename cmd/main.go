package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/evandrolg/insta-motion/pkg/arguments"
	"github.com/evandrolg/insta-motion/pkg/converter"
)

func main() {
	var imagePath string
	var effect string
	var help bool
	flag.StringVar(&imagePath, "image", "", "Path to the image")
	flag.StringVar(&effect, "effect", "", "Effect to apply to the video")
	flag.BoolVar(&help, "help", false, "Show help")
	flag.Parse()

	if help || len(os.Args) == 1 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if err := arguments.CheckArguments(imagePath, effect); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outputVideo, err := converter.ConvertToVideo(imagePath, effect)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Video created: %s\n", outputVideo)
}
