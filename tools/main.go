package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	// 定义输入文件名和输出文件名
	inputFilename := "input.mp4"
	outputFilename := "output.mp4"

	//os.TempDir("recently")
	// 获取视频时长
	duration, err := getVideoDuration(inputFilename)
	if err != nil {
		fmt.Println("获取视频时长失败：", err)
		return
	}

	// 计算开始时间和结束时间
	startTime := duration - 30
	endTime := duration

	// 使用FFmpeg进行剪辑
	cmd := exec.Command("ffmpeg", "-i", inputFilename, "-ss", fmt.Sprintf("%.2f", startTime), "-to", fmt.Sprintf("%.2f", endTime), "-c", "copy", outputFilename)
	err = cmd.Run()
	if err != nil {
		fmt.Println("剪辑视频失败：", err)
		return
	}

	fmt.Println("视频回放成功，已保存为", outputFilename)
}

// 获取视频时长
func getVideoDuration(filename string) (float64, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=duration", "-of", "default=noprint_wrappers=1:nokey=1", filename)
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	duration, err := time.ParseDuration(fmt.Sprintf("%ss", string(output)))
	if err != nil {
		return 0, err
	}

	return duration.Seconds(), nil
}
