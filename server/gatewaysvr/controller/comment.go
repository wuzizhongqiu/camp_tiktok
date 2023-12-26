package controller

import (
	"gatewaysvr/log"
	"gatewaysvr/response"
	"gatewaysvr/utils"
	"github.com/Keqing-win/camp_tiktok/pkg/pb"
	"github.com/gin-gonic/gin"
)

type CommentActionTopReq struct {
	VideoId     int64  `json:"video_id" form:"video_id" binding:"required"`
	CommentText string `json:"comment_text" form:"comment_text"  binding:"required"`
	ActionType  int64  `json:"action_type" form:"action_type" binding:"required"`
	CommentId   int64  `json:"comment_id" form:"comment_id"  binding:"required"`
}

type CommentActionTopResp struct {
	StatusCode int32       `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
	StatusMsg  string      `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	Comment    *pb.Comment `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment"`
}

type CommentTopListReq struct {
	VideoId int64 `json:"video_id" form:"video_id" binding:"required"`
	Page    int64 `json:"page" form:"page" binding:"required"`
	Size    int64 `json:"size" form:"size"  binding:"required"`
}

type CommentTopListResp struct {
	StatusCode int32            `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
	StatusMsg  string           `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	Comments   []*pb.CommentGet `protobuf:"bytes,3,opt,name=comments_top,proto3" json:"comments_top"`
}

type CommentOtherReq struct {
	VideoId      int64  `json:"video_id" form:"video_id" binding:"required"`
	ParentId     int64  `json:"parent_id" form:"parent_id" binding:"required"`
	ParentUserId int64  `json:"parent_user_id" form:"parent_user_id" binding:"required"`
	CommentText  string `json:"comment_text" form:"comment_text"  binding:"required"`
	ReplyId      int64  `json:"reply_id" form:"reply_id" binding:""`
	ActionType   int64  `json:"action_type" form:"action_type" binding:"required"`
	CommentId    int64  `json:"comment_id" form:"comment_id"  binding:""`
}

type CommentActionOtherResp struct {
	StatusCode int32       `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
	StatusMsg  string      `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	Comment    *pb.Comment `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment"`
}

type CommentOtherListReq struct {
	VideoId  int64 `json:"video_id" form:"video_id" binding:"required"`
	ParentId int64 `json:"parent_id" form:"parent_id" binding:"required"`
	Page     int64 `json:"page" form:"page" binding:"required"`
	Size     int64 `json:"size" form:"size"  binding:"required"`
}

type CommentOtherListResp struct {
	StatusCode    int32              `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
	StatusMsg     string             `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	Comments      []*pb.Comment      `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment"`
	ReplyComments []*pb.ReplyComment `protobuf:"bytes,3,opt,name=reply_comment,proto3" json:"reply_comment"`
}

func CommentActionTop(ctx *gin.Context) {
	var req CommentActionTopReq
	err := ctx.ShouldBind(&req)
	log.Info("get req %+v", req)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	client := utils.GetCommentSvrClient()
	uid, _ := ctx.Get("UserId")
	pbRsp, err := client.CommentTopAction(ctx, &pb.CommentActionTopReq{
		UserId:      uid.(int64),
		VideoId:     req.VideoId,
		CommentText: req.CommentText,
		ActionType:  req.ActionType,
		CommentId:   req.CommentId,
	})
	log.Info("pbRsp =%+v", pbRsp)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	log.Info("create comment %+v", pbRsp)
	response.Success(ctx, "success", &CommentActionTopResp{
		Comment: pbRsp.Comment,
	})
	//增加视频评论数不在这里写，打算单独给video写个服务

}

func GetCommentTopList(ctx *gin.Context) {
	var req CommentTopListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	client := utils.GetCommentSvrClient()
	resp, err := client.GetTopCommentList(ctx, &pb.GetTopCommentListReq{
		VideoId: req.VideoId,
		Page:    req.Page,
		Size:    req.Size,
	})
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", &CommentTopListResp{Comments: resp.Comments})
}

func GetCommentOtherList(ctx *gin.Context) {
	var req CommentOtherListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	log.Info("GetCommentOtherList req==", req)
	client := utils.GetCommentSvrClient()
	resp, err := client.GetOtherCommentList(ctx, &pb.GetOtherCommentListReq{
		VideoId:  req.VideoId,
		ParentId: req.ParentId,
		Page:     req.Page,
		Size:     req.Size,
	})
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", &CommentOtherListResp{Comments: resp.Comment, ReplyComments: resp.ReplyComment})
}

func CommentActionOther(ctx *gin.Context) {
	var req CommentOtherReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	log.Info("CommentActionOther req==", req)
	uid, _ := ctx.Get("UserId")
	client := utils.GetCommentSvrClient()
	resp, err := client.CommentOtherAction(ctx, &pb.CommentActionOtherReq{
		UserId:       uid.(int64),
		VideoId:      req.VideoId,
		ParentId:     req.ParentId,
		ParentUserId: req.ParentUserId,
		CommentText:  req.CommentText,
		ReplyId:      req.ReplyId,
		ActionType:   req.ActionType,
		CommentId:    req.CommentId,
	})
	if err != nil {
		return
	}
	response.Success(ctx, "success", &CommentActionOtherResp{
		Comment: resp.Comment,
	})
}
