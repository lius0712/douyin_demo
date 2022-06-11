package video

import (
	"fmt"
	"path/filepath"

	"github.com/RaymondCode/simple-demo/config"
)

func GetVideoLocalPath(name string) string {
	return filepath.ToSlash(filepath.Join(config.Config.LocalVideoPath, fmt.Sprintf("%v.mp4", name)))
}

func GetCoverLocalPath(name string) string {
	return filepath.ToSlash(filepath.Join(config.Config.LocalVideoPath, fmt.Sprintf("%v.jpg", name)))
}

func GetVideoRemotePath(name string) string {
	return config.Config.DbUrl +
		filepath.ToSlash(filepath.Join(config.Config.RemoteVideoPath, fmt.Sprintf("%v.mp4", name)))
}

func GetCoverRemotePath(name string) string {
	return config.Config.DbUrl +
		filepath.ToSlash(filepath.Join(config.Config.RemoteVideoPath, fmt.Sprintf("%v.jpg", name)))
}
