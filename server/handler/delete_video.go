package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/HelloMySugar/service-video/model"
	"github.com/HelloMySugar/service-video/types"
	"log"
	"net/http"
	"strconv"
)

func DeleteVideo(c *gin.Context) {
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
	}
	if err := video.Delete(model.DBObj); err != nil {
		log.Printf("%v", errors.WithMessage(err, "delete video"))
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "db delete video",
		})
		return
	}
}
