package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/HelloMySugar/service-video/model"
	"github.com/HelloMySugar/service-video/types"
)

type GetPageBeanByValueRequest struct {
	Value       string `form:"value"`
	CurrentPage int    `form:"currentPage"`
	PageSize    int    `form:"pageSize"`
}

func GetPageBeanByValue(c *gin.Context) {
	req := GetPageBeanByValueRequest{}
	if err := c.ShouldBind(&req); err != nil {
		log.Printf("%v", errors.WithMessage(err, "bind query"))
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "bind query error",
		})
		return
	}

	page := types.Pager{
		Limit:  req.PageSize,
		Offset: (req.CurrentPage - 1) * req.PageSize,
	}
	video := model.Video{}
	videoList, cnt, err := video.FetchListNameLike(model.DBObj, page, req.Value)
	if err != nil {
		log.Printf("%v", errors.WithMessage(err, "fetch video list"))
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "fetch video list error",
		})
		return
	}

	totalPage := 0
	if req.PageSize != 0 {
		totalPage = (cnt-1)/req.PageSize + 1
	}
	resp := GetPageBeanResponse{
		TotalCount:  cnt,
		TotalPage:   totalPage,
		CurrentPage: req.CurrentPage,
		PageSize:    req.PageSize,
		List:        make([]Video, 0, len(videoList)),
	}

	for _, v := range videoList {
		resp.List = append(resp.List, Video{
			VID:         strconv.Itoa(int(v.ID)),
			VName:       v.Name,
			ReleaseTime: v.ReleaseTime.Format("2006-01-02"),
			Visited:     int(v.Visited),
			Keyword:     v.Keyword,
			VUrl:        v.Url,
			CoverUrl:    v.CoverUrl,
			CID:         strconv.Itoa(int(v.CategoryID)),
		})
	}

	c.JSON(http.StatusOK, resp)
}
