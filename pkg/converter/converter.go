package converter

import (
	"fmt"
	"os"
	"os/exec"
)

func ConvertToVideo(imagePath string) (string, error) {
	outputVideo := "output.mp4"
	ffmpegArgs := []string{
		"-i", imagePath,
		"-vf", "zoompan=z='min(zoom+0.001,1.5)':d=200, scale=1080:1920,setsar=1:1",
		"-c:a", "copy",
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
