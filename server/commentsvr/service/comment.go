package service

import (
	"commentsvr/constant"
	"commentsvr/log"
	"commentsvr/respository"
	"commentsvr/utils"
	"github.com/Keqing-win/camp_tiktok/pkg/pb"
	"golang.org/x/net/context"
	"sync"
)

type CommentService struct {
	pb.UnimplementedCommentServiceServer
}

func (c *CommentService) GetCommentSum(ctx context.Context, in *pb.GetCommentNumReq) (*pb.GetCommentNumRsp, error) {
	sum, err := respository.GetCommentSum(in.VideoId)
	if err != nil {
		return nil, err
	}
	return &pb.GetCommentNumRsp{Sum: sum}, nil
}

func (c *CommentService) GetOtherCommentList(ctx context.Context, in *pb.GetOtherCommentListReq) (*pb.GetOtherCommentListRsp, error) {
	log.Debugf("GetOtherCommentList 调用 req=%+v", in)
	normalComments, replyComments, err := respository.GetOtherCommentList(in.VideoId, in.ParentId, int(in.Page), int(in.Size))

	if err != nil {
		return nil, err
	}
	var CommentPb []*pb.Comment
	for _, normalComment := range normalComments {
		log.Debugf("normalComment=%+v", normalComment)
		info, err := utils.GetUserSvrClient().GetUserInfo(ctx, &pb.GetUserInfoRequest{Id: normalComment.UserId})
		if err != nil {
			return nil, err
		}
		CommentPb = append(CommentPb, ChangeToPbComment(info.UserInfo, normalComment))
	}
	var replyCommentPb []*pb.ReplyComment
	for _, replyComment := range replyComments {
		log.Debugf("replyComment=%+v", replyComment)

		data, err := ChangToPbReplyComment(ctx, replyComment)
		if err != nil {
			return nil, err
		}
		replyCommentPb = append(replyCommentPb, data)
	}
	return &pb.GetOtherCommentListRsp{
		Comment:      CommentPb,
		ReplyComment: replyCommentPb,
	}, nil
}

func (c *CommentService) GetTopCommentList(ctx context.Context, in *pb.GetTopCommentListReq) (*pb.GetTopCommentListRsp, error) {
	comments, err := respository.GetTopCommentList(in.VideoId, int(in.Page), int(in.Size))
	if err != nil {
		return nil, err
	}
	var getCommentList []*pb.CommentGet
	for _, comment := range comments {
		commentResp, err := GetCommentGet(ctx, comment)
		if err != nil {
			return nil, err
		}
		getCommentList = append(getCommentList, commentResp)
	}

	return &pb.GetTopCommentListRsp{Comments: getCommentList}, nil

}

func (c *CommentService) CommentOtherAction(ctx context.Context, in *pb.CommentActionOtherReq) (*pb.CommentActionRsp, error) {
	if in.ActionType == constant.Create {
		log.Info("调用 CommentOtherAction")
		comment, err := respository.CreateOtherComment(in.VideoId, in.UserId, in.ParentId, in.ReplyId, in.ParentUserId, in.CommentText)
		if err != nil {
			return nil, err
		}
		info, err := utils.GetUserSvrClient().GetUserInfo(ctx, &pb.GetUserInfoRequest{Id: in.UserId})
		if err != nil {
			return nil, err
		}
		return &pb.CommentActionRsp{Comment: ChangeToPbComment(info.UserInfo, comment)}, nil
	} else {
		err := respository.DeleteOtherComment(in.VideoId, in.CommentId)
		if err != nil {
			return nil, err
		}
		return &pb.CommentActionRsp{}, nil
	}
}

func (c *CommentService) CommentTopAction(ctx context.Context, in *pb.CommentActionTopReq) (*pb.CommentActionRsp, error) {
	if in.ActionType == constant.Create {
		comment, err := respository.CreateTopComment(in.UserId, in.VideoId, in.CommentText)
		if err != nil {
			return nil, err
		}
		info, err := utils.GetUserSvrClient().GetUserInfo(ctx, &pb.GetUserInfoRequest{Id: in.UserId})
		if err != nil {
			return nil, err
		}
		return &pb.CommentActionRsp{Comment: ChangeToPbComment(info.UserInfo, comment)}, nil
	} else {
		err := respository.DeleteTopComment(in.VideoId, in.CommentId)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

func GetCommentGet(ctx context.Context, comment *respository.Comment) (*pb.CommentGet, error) {
	var wg sync.WaitGroup
	wg.Add(2)
	var err error
	commentGet := &pb.CommentGet{} // 初始化 commentGet
	go func() {
		defer wg.Done()
		var info *pb.GetUserInfoResponse
		info, err = utils.GetUserSvrClient().GetUserInfo(ctx, &pb.GetUserInfoRequest{Id: comment.UserId})
		commentPb := ChangeToPbComment(info.UserInfo, comment)
		commentGet.Comment = commentPb
	}()
	go func() {
		defer wg.Done()
		commentGet.Count, err = respository.GetTopCommentReplyNum(comment.Id)
	}()
	wg.Wait()
	if err != nil {
		return nil, err
	}
	return commentGet, nil
}

func ChangToPbReplyComment(ctx context.Context, comment *respository.ReplyComment) (*pb.ReplyComment, error) {
	var wg sync.WaitGroup
	wg.Add(2)
	var Userinfo *pb.GetUserInfoResponse
	var err error
	var ReplyUserInfo *pb.GetUserInfoResponse
	go func() {
		defer wg.Done()
		Userinfo, err = utils.GetUserSvrClient().GetUserInfo(ctx, &pb.GetUserInfoRequest{Id: comment.Comment.UserId})
	}()
	go func() {
		defer wg.Done()
		ReplyUserInfo, err = utils.GetUserSvrClient().GetUserInfo(ctx, &pb.GetUserInfoRequest{Id: comment.ReplyUserId})
	}()
	if err != nil {
		return nil, err
	}
	wg.Wait()
	return &pb.ReplyComment{
		Id:         comment.Comment.Id,
		User:       Userinfo.UserInfo,
		ReplyUser:  ReplyUserInfo.UserInfo,
		Content:    comment.Comment.CommentText,
		CreateDate: comment.Comment.CreateTime.Format(constant.DefaultTime),
		Like:       comment.Comment.FavoriteNum,
	}, nil
}

func ChangeToPbComment(info *pb.UserInfo, comment *respository.Comment) *pb.Comment {
	return &pb.Comment{
		Id:         comment.Id,
		User:       info,
		Content:    comment.CommentText,
		CreateDate: comment.CreateTime.Format(constant.DefaultTime),
		Like:       comment.FavoriteNum,
	}
}
