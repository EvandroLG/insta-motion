package converter

import (
	"fmt"
	"os"
	"os/exec"
)

func ConvertToVideo(imagePath string, effect string) (string, error) {
	outputVideo := "output.mp4"

	ffmpegArgs := []string{
		"-i", imagePath,
		"-vf", fmt.Sprintf("zoompan=z='%s':d=200, scale=1080:1920,setsar=1:1", getEffect(effect)),
		"-c:a", "copy",
		"-loglevel", "error",
		"-t", "15",
		outputVideo,
	}

	cmd := exec.Command("ffmpeg", ffmpegArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("Error converting image to video: %v", err)
	}

	return outputVideo, nil
}

func getEffect(effect string) string {
	switch effect {
	case "zoom-out":
		return "if(lte(zoom,1.0),1.5,max(1.0,zoom-0.001))"
	default:
		return "min(zoom+0.001,1.5)"
	}
}
