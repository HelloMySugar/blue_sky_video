package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/HelloMySugar/service-video/global/settings"
	"path/filepath"
	"strings"
)

const picturePrefix string = "/pictures/"

func GetPicture(c *gin.Context) {
	path := strings.TrimPrefix(c.Request.URL.Path, picturePrefix)

	c.File(filepath.Join(settings.ServerSetting.DataDir, path))
}
