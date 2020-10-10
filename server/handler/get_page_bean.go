package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/HelloMySugar/service-video/model"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/HelloMySugar/service-video/types"
)

type GetPageBeanRequest struct {
	CID         string `form:"cid"`
	CurrentPage int    `form:"currentPage"`
	PageSize    int    `form:"pageSize"`
}

type Video struct {
	VID         string `json:"vid"`
	VName       string `json:"vname"`
	ReleaseTime string `json:"releaseTime"`
	Visited     int    `json:"visited"`
	Keyword     string `json:"keyword"`
	VUrl        string `json:"vurl"`
	CoverUrl    string `json:"coverurl"`
	CID         string `json:"cid"`
}

type GetPageBeanResponse struct {
	TotalCount  int     `json:"totalCount"`
	TotalPage   int     `json:"totalPage"`
	CurrentPage int     `json:"currentPage"`
	PageSize    int     `json:"pageSize"`
	List        []Video `json:"list"`
}

func GetPageBean(c *gin.Context) {
	req := GetPageBeanRequest{}
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
	cid, err := strconv.Atoi(req.CID)
	if err != nil {
		log.Printf("%v", errors.WithMessage(err, "cid atoi"))
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "cid should be number",
		})
		return
	}
	video := model.Video{
		CategoryID: uint(cid),
	}
	videoList, cnt, err := video.FetchList(model.DBObj, page)
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
