package main

import (
	"github.com/gin-gonic/gin"
	"github.com/HelloMySugar/service-video/global/settings"
	"github.com/HelloMySugar/service-video/handler"
	"github.com/HelloMySugar/service-video/model"
)

func main() {
	settings.Setup()
	model.Setup()

	r := gin.Default()
	r.GET("/videoCategory", handler.HandleVideoCategory)
	r.POST("/video/getPageBean", handler.GetPageBean)
	r.POST("/video/getPageBeanByValue", handler.GetPageBeanByValue)
	r.POST("/collection/getCollectionNum", handler.GetCollectionNum)
	r.GET("/videos/*path", handler.GetVideo)
	r.GET("/pictures/*path", handler.GetPicture)

	r.POST("/v0/videos", handler.UploadVideo)
	r.DELETE("/v0/videos/:vid", handler.DeleteVideo)
	r.PUT("/v0/videos/:vid", handler.UpdateVideo)
	r.Run(settings.ServerSetting.Address)
}
