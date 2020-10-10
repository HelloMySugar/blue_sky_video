package handler

import (
	"log"
	"net/http"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/HelloMySugar/service-video/global/settings"
	"github.com/HelloMySugar/service-video/model"
	"github.com/HelloMySugar/service-video/types"
)

type UploadVideoRequest struct {
	CID         string `form:"cid"`
	Keywords    string `form:"keywords"`
	VName       string `form:"vname"`
	ReleaseTime string `form:"releaseTime"`
}

func UploadVideo(c *gin.Context) {
	req := UploadVideoRequest{}
	if err := c.ShouldBind(&req); err != nil {
		log.Printf("%v", errors.WithMessage(err, "bind query"))
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "bind query error",
		})
		return
	}

	rt, err := time.Parse("2006-01-02", req.ReleaseTime)
	if err != nil {
		log.Printf("%v", errors.WithMessage(err, "parse release time"))
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "release time format error",
		})
		return
	}

	videoFile, err := c.FormFile("video")
	if err != nil {
		log.Printf("%v", errors.WithMessage(err, "video file not found"))
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "video file not found",
		})
		return
	}
	vfn := videoFile.Filename
	vfp := filepath.Join(settings.ServerSetting.DataDir, vfn)
	if err = c.SaveUploadedFile(videoFile, vfp); err != nil {
		log.Printf("%v", errors.WithMessage(err, "video file save"))
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "video file save error",
		})
		return
	}

	pfn := videoFile.Filename + ".jpg"
	pfp := filepath.Join(settings.ServerSetting.DataDir, pfn)
	cmd := exec.Command("ffmpeg", "-i", vfp, "-y", "-f", "image2", "-ss", "1", "-vframes", "1", pfp)
	if err := cmd.Run(); err != nil {
		log.Printf("%v", errors.WithMessage(err, "ffmpeg convert"))
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "extract frame error",
		})
		return
	}

	cid, err := strconv.Atoi(req.CID)
	if err != nil {
		log.Printf("%v", errors.WithMessage(err, "cid atoi"))
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "cid should be number",
		})
		return
	}
	video := model.Video{
		Name:        req.VName,
		Keyword:     req.Keywords,
		CategoryID:  uint(cid),
		Url:         path.Join(videoPrefix, vfn),
		CoverUrl:    path.Join(picturePrefix, pfn),
		ReleaseTime: rt,
	}
	if err := video.Create(model.DBObj); err != nil {
		log.Printf("%v", errors.WithMessage(err, "create video"))
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "db create video",
		})
		return
	}
}
