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

type GetCollectionNumRequest struct {
	VID string `form:"vid"`
}

func GetCollectionNum(c *gin.Context) {
	req := GetCollectionNumRequest{}
	if err := c.ShouldBind(&req); err != nil {
		log.Printf("%v", errors.WithMessage(err, "bind query"))
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "bind query error",
		})
		return
	}

	vid, err := strconv.Atoi(req.VID)
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
	if err := video.Fetch(model.DBObj); err != nil {
		log.Printf("%v", errors.WithMessage(err, "fetch video"))
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "fetch video from db",
		})
		return
	}
	video.Visited = video.Visited + 1
	if err := video.Update(model.DBObj); err != nil {
		log.Printf("%v", errors.WithMessage(err, "update video"))
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "update video",
		})
		return
	}

	c.Status(http.StatusOK)
}
