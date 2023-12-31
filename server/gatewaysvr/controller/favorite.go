package controller

import (
	"gatewaysvr/response"
	"gatewaysvr/utils"
	"github.com/Keqing-win/camp_tiktok/pkg/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FavoriteVideoAction(c *gin.Context) {
	vidStr := c.Query("video_id")
	vid, err := strconv.ParseInt(vidStr, 10, 64)
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}
	actionTypeStr := c.Query("action_type")
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}
	uid, _ := c.Get("UserId")
	client := utils.GetFavoriteSvrClient()
	_, err = client.FavoriteVideoAction(c, &pb.FavoriteVideoActionReq{
		UserId:     uid.(int64),
		VideoId:    vid,
		ActionType: actionType,
	})
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}
	response.Success(c, "success", nil)
	//todo:给视频的喜爱总数增加，给被点赞视频的创作者被喜爱数加一，给点赞者的喜爱数加一
}

func GetFavoriteVideoList(c *gin.Context) {
	uid, _ := c.Get("UserId")
	client := utils.GetFavoriteSvrClient()
	_, err := client.GetFavoriteVideoIdList(c, &pb.GetFavoriteVideoIdListReq{UserId: uid.(int64)})
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}
	response.Success(c, "success", nil)
}

func FavoriteCommentAction(c *gin.Context) {
	cidStr := c.Query("comment_id")
	cid, err := strconv.ParseInt(cidStr, 10, 64)
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}
	actionTypeStr := c.Query("action_type")
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}
	uid, _ := c.Get("UserId")
	client := utils.GetFavoriteSvrClient()
	_, err = client.FavoriteCommentAction(c, &pb.FavoriteCommentActionReq{
		UserId:     uid.(int64),
		CommentId:  cid,
		ActionType: actionType,
	})
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}
	response.Success(c, "success", nil)
	//todo:给评论的喜爱总数增加
}
