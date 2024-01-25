package controller

import (
	"gatewaysvr/log"
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
	_, err = utils.GetVideoSvrClient().UpdateFavoriteCount(c, &pb.UpdateFavoriteCountReq{
		VideoId:    vid,
		ActionType: actionType,
	})
	if err != nil {
		log.Errorf("utils.GetVideoSvrClient().UpdateFavoriteCount err==%v", err)
		response.Fail(c, err.Error(), nil)
		return
	}

	// 查询video表的author_id
	videoInfoResp, err := utils.GetVideoSvrClient().GetVideoInfoList(c, &pb.GetVideoInfoListReq{
		VideoId: []int64{vid},
	})
	if err != nil {
		log.Errorf("GetVideoInfoList failed, err:%v", err)
		response.Fail(c, err.Error(), nil)
		return
	}

	var authorId = videoInfoResp.VideoInfoList[0].AuthorId

	// 更新user表的 total_favorited_count（更新视频作者获赞数）
	_, err = utils.GetUserSvrClient().UpdateUserFavoritedCount(c, &pb.UpdateUserFavoritedCountReq{
		UserId:     authorId,
		ActionType: actionType,
	})
	if err != nil {
		log.Errorf("UpdateUserFavoritedCount failed, err:%v", err)
		response.Fail(c, err.Error(), nil)
		return
	}

	// 更新user表的 favorite_count更新我喜欢的视频数）
	_, err = utils.GetUserSvrClient().UpdateUserFavoriteCount(c, &pb.UpdateUserFavoriteCountReq{
		UserId:     uid.(int64),
		ActionType: actionType,
	})
	if err != nil {
		log.Errorf("UpdateUserFavoriteCount failed, err:%v", err)
		response.Fail(c, err.Error(), nil)
		return
	}

	response.Success(c, "success", nil)
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

	var nums int64
	if actionType == 1 {
		nums = 1
	} else {
		nums = -1
	}
	_, err = utils.CommentSvrClient.CommentLikeAdd(c, &pb.CommentAddLikeNumReq{CommentId: cid, Num: nums})
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}

	response.Success(c, "success", nil)
}
