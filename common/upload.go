package common

import (
	"bytes"
	"fmt"
	"github.com/u2takey/ffmpeg-go"
	"io"
	"os"
	"os/exec"
	"strconv"
)

// GetUserURL 得到一名用户对应的云端存储路径
func GetUserURL(userId int64) string {
	return "https://" + UploadServerBucket + ".oss-cn-hangzhou.aliyuncs.com/" + UploadBucketDirectory + "/" + strconv.FormatInt(userId, 10) + "/"
}

func exampleReadFrameAsJpeg(inFileName string, frameNum int) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg_go.Input(inFileName).
		Filter("select", ffmpeg_go.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg_go.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// ExtractCoverFromVideo 从视频中截取图像的第一帧
func ExtractCoverFromVideo(pathVideo, pathImg string) error {
	command := "ffmpeg"
	frameExtractionTime := "0"
	image_mode := "image2"
	vtime := "0.001"

	// create the command
	cmd := exec.Command(command,
		"-i", pathVideo,
		"-y",
		"-f", image_mode,
		"-ss", frameExtractionTime,
		"-t", vtime,
		"-y", pathImg)

	// run the command and don't wait for it to finish. waiting exec is run
	// fmt.Println(cmd.String())
	err := cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}

func GetUploadURL(userId int64, fileName string) string {
	return GetUserURL(userId) + fileName
}
