package video

import ffmpeg_go "github.com/u2takey/ffmpeg-go"

// ExtractFrame 提取videoPath视频的第n帧保存至framePath中
func ExtractFrame(videoPath string, framePath string) error {
	err := ffmpeg_go.Input(videoPath).
		Output(framePath, ffmpeg_go.KwArgs{
			"vframes": "1",
		}).
		OverWriteOutput().
		Run()
	return err
}
