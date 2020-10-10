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

type VideoCategoryResponse struct {
	CID   string `json:"cid"`
	CName string `json:"cname"`
}

func HandleVideoCategory(c *gin.Context) {
	category := model.Category{}
	categoryList, _, err := category.FetchAllList(model.DBObj)
	if err != nil {
		log.Printf("%v", errors.WithMessage(err, "fetch all category list"))
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "fetch all category list error",
		})
	}
	resp := make([]VideoCategoryResponse, 0, len(categoryList))
	for _, ca := range categoryList {
		resp = append(resp, VideoCategoryResponse{
			CID:   strconv.Itoa(int(ca.ID)),
			CName: ca.Name,
		})
	}
	c.JSON(http.StatusOK, resp)
}
