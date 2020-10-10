package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/HelloMySugar/service-video/model"
	"github.com/HelloMySugar/service-video/types"
)

type UpdateVideoRequest struct {
	CID         string `form:"cid"`
	Name        string `form:"vname"`
	ReleaseTime string `form:"releaseTime"`
	Keywords    string `form:"keywords"`
}

func UpdateVideo(c *gin.Context) {
	req := UpdateVideoRequest{}
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

	cid, err := strconv.Atoi(req.CID)
	if err != nil {
		log.Printf("%v", errors.WithMessage(err, "cid atoi"))
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "cid should be number",
		})
		return
	}

	vidParam := c.Param("vid")
	vid, err := strconv.Atoi(vidParam)
	if err != nil {
		log.Printf("%v", errors.WithMessage(err, "vid atoi"))
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "vid should be number",
		})
		return
	}
	video := model.Video{
		Model: model.Model{
			ID: uint(vid),
		},
		Name:        req.Name,
		ReleaseTime: rt,
		CategoryID:  uint(cid),
		Keyword:     req.Keywords,
	}
	if err := video.Update(model.DBObj); err != nil {
		log.Printf("%v", errors.WithMessage(err, "update video"))
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "db update video",
		})
		return
	}
}
