package controller

import (
	"gatewaysvr/log"
	"gatewaysvr/response"
	"gatewaysvr/utils"
	"github.com/Keqing-win/camp_tiktok/pkg/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DouyinRelationListResponse struct {
	StatusCode int32          `json:"status_code"`
	StatusMsg  string         `json:"status_msg,omitempty"`
	UserList   []*pb.UserInfo `json:"user_list,omitempty"`
}

func RelationAction(ctx *gin.Context) {
	tokens, _ := ctx.Get("UserId")
	tokenUserId := tokens.(int64)

	toUserId := ctx.Query("to_user_id")
	toUid, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	actionStr := ctx.Query("action_type")

	actionType, err := strconv.ParseInt(actionStr, 10, 64)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	log.Infof("RelationAction tokenUserId:%d, toUid:%d, actionType:%d", tokenUserId, toUid, actionType)
	// 1.关注 2.取消关注
	_, err = utils.GetRelationSvrClient().RelationAction(ctx, &pb.RelationActionReq{
		ToUserId:   toUid,
		SelfUserId: tokenUserId,
		ActionType: actionType,
	})
	if err != nil {
		log.Errorf("RelationAction error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}

	// 我关注的人++
	_, err = utils.GetUserSvrClient().UpdateUserFollowCount(ctx, &pb.UpdateUserFollowCountReq{
		UserId:     tokenUserId,
		ActionType: actionType,
	})
	if err != nil {
		log.Errorf("RelationAction error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}

	// 被我关注的人的粉丝++
	_, err = utils.GetUserSvrClient().UpdateUserFollowerCount(ctx, &pb.UpdateUserFollowerCountReq{
		UserId:     toUid,
		ActionType: actionType,
	})
	if err != nil {
		log.Errorf("RelationAction error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}

	response.Success(ctx, "success", nil)
}

func GetFollowList(ctx *gin.Context) {
	UserId := ctx.Query("user_id")
	tokens, _ := ctx.Get("UserId")
	tokenUserId := tokens.(int64)
	uid, err := strconv.ParseInt(UserId, 10, 64)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}

	// 获取关注列表
	getRelationFollowListRsp, err := utils.GetRelationSvrClient().GetRelationFollowList(ctx, &pb.GetRelationFollowListReq{
		UserId: uid,
	})
	if err != nil {
		log.Errorf("GetFollowList error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}

	// 去用户服务获取用户信息
	resp, err := utils.GetUserSvrClient().GetUserInfoList(ctx, &pb.GetUserInfoListRequest{
		IdList: getRelationFollowListRsp.FollowList,
	})
	if err != nil {
		log.Errorf("GetFollowList error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	var followUintList = make([]*pb.FollowUint, 0)
	for _, user := range resp.UserInfoList {
		followUintList = append(followUintList, &pb.FollowUint{
			SelfUserId: tokenUserId,
			UserIdList: user.Id,
		})
	}
	isFollowedRsp, err := utils.GetRelationSvrClient().IsFollowDict(ctx, &pb.IsFollowDictReq{
		FollowUintList: followUintList,
	})
	if err != nil {
		log.Errorf("utils.GetRelationSvrClient().IsFollowDict err==%v", err)
	}
	for _, user := range resp.UserInfoList {
		var followUint = strconv.FormatInt(tokenUserId, 10) + "_" + strconv.FormatInt(user.Id, 10)
		user.IsFollow = isFollowedRsp.IsFollowDict[followUint]
	}
	response.Success(ctx, "success", &DouyinRelationListResponse{
		UserList: resp.UserInfoList,
	})
}

func GetFollowerList(ctx *gin.Context) {
	UserId := ctx.Query("user_id")
	uid, err := strconv.ParseInt(UserId, 10, 64)
	if err != nil {
		log.Errorf("GetFollowerList ParseInt error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	log.Info("utils.GetRelationSvrClient().GetRelationFollowerList init")
	// 获取关注者列表
	getRelationFollowerListRsp, err := utils.GetRelationSvrClient().GetRelationFollowerList(ctx, &pb.GetRelationFollowerListReq{
		UserId: uid,
	})
	log.Info(getRelationFollowerListRsp.FollowerList)
	if err != nil {
		log.Errorf("GetFollowerList error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}

	// 去用户服务获取用户信息
	resp, err := utils.GetUserSvrClient().GetUserInfoList(ctx, &pb.GetUserInfoListRequest{
		IdList: getRelationFollowerListRsp.FollowerList,
	})

	response.Success(ctx, "success", &DouyinRelationListResponse{
		UserList: resp.UserInfoList,
	})
}
