package converter

import (
	"fmt"
	"os"
	"os/exec"
)

// ConvertToVideo converts an image to a video with the specified effect.
// The video is optimized for Instagram stories with a 9:16 aspect ratio.
// The function uses ffmpeg to create a video from the image.
func ConvertToVideo(imagePath string, effect string) (string, error) {
	outputVideo := "output.mp4"

	// Instagram story optimal specifications
	// 1080x1920 (9:16 aspect ratio)
	// Frame rate: 30fps
	// Bit rate: 3-5 Mbps
	ffmpegArgs := []string{
		"-i", imagePath,
		"-vf", fmt.Sprintf("zoompan=z='%s':d=200:s=1080x1920, fps=30, format=yuv420p", getEffect(effect)),
		"-c:v", "libx264", // Use H.264 codec
		"-b:v", "4M", // Set bitrate to 4 Mbps
		"-pix_fmt", "yuv420p", // Pixel format for better compatibility
		"-movflags", "+faststart", // Optimize for web streaming
		"-loglevel", "error",
		"-t", "15", // 15-second duration
		"-y", // Overwrite output file if it exists
		outputVideo,
	}

	cmd := exec.Command("ffmpeg", ffmpegArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error converting image to video: %v", err)
	}

	return outputVideo, nil
}

// getEffect returns the ffmpeg zoompan filter expression based on the effect.
// The effect can be "zoom-in" or "zoom-out".
func getEffect(effect string) string {
	switch effect {
	case "zoom-out":
		// Start zoomed in and slowly zoom out
		return "if(lte(zoom,1.0),1.5,max(1.0,zoom-0.0008))"
	default: // zoom-in
		// Start zoomed out and slowly zoom in
		return "min(zoom+0.0008,1.5)"
	}
}
