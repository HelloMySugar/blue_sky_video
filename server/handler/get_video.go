package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/HelloMySugar/service-video/global/settings"
	"path/filepath"
	"strings"
)

const videoPrefix string = "/videos/"

func GetVideo(c *gin.Context) {
	path := strings.TrimPrefix(c.Request.URL.Path, videoPrefix)

	c.File(filepath.Join(settings.ServerSetting.DataDir, path))
}
