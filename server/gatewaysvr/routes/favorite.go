package routes

import (
	"gatewaysvr/controller"
	"gatewaysvr/utils/common"
	"github.com/gin-gonic/gin"
)

func FavoriteRoutes(r *gin.RouterGroup) {
	like := r.Group("like")
	{
		like.POST("/action/video/", common.AuthMiddleware(), controller.FavoriteVideoAction)
		like.GET("/list/video/", common.AuthMiddleware(), controller.GetFavoriteVideoList)
		like.POST("/action/comment/", common.AuthMiddleware(), controller.FavoriteCommentAction)
	}
}
