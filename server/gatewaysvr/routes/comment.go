package routes

import (
	"gatewaysvr/controller"
	"gatewaysvr/utils/common"
	"github.com/gin-gonic/gin"
)

func CommentRoutes(r *gin.RouterGroup) {
	comment := r.Group("comment")
	{
		comment.POST("/actionTop/", common.AuthMiddleware(), controller.CommentActionTop)
		comment.GET("/listTop/", common.AuthWithOutMiddleware(), controller.GetCommentTopList)
		comment.GET("/listOther/", common.AuthWithOutMiddleware(), controller.GetCommentOtherList)
		comment.POST("/actionOther/", common.AuthMiddleware(), controller.CommentActionOther)
	}

}
