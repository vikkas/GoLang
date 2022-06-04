package controller

import(
	"github.com/gin-gonic/gin"
	"github.com/pragmaticreviews/golang-gin-poc/entity"
)

type VideoController interface{
	FindAll() []entity.VideoController
	Save(ctx *gin.Context) entity.video
}

type controller struct{
	service service.VideoService
}

func New(service service.VideoService) VideoController{
	return controller{
		service: service,
	}
}

func (c *controller FindAll() []entity.Video{
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context){
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}