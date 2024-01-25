package service

import (
	"github.com/Keqing-win/camp_tiktok/pkg/pb"
	"golang.org/x/net/context"
	"likesvr/constant"
	"likesvr/respository"
	"strconv"
)

type LikeService struct {
	pb.UnimplementedFavoriteServiceServer
}

func (l *LikeService) GetCommentLikeSum(ctx context.Context, in *pb.CommentLikeSumReq) (*pb.CommentLikeSumRsq, error) {
	//这里因为刚git提交完拉去不到最新的修改，应该是in，commentId
	sum, err := respository.GetCommentLikeNum(in.CommentId)
	if err != nil {
		return nil, err
	}
	return &pb.CommentLikeSumRsq{LikeNums: sum}, nil
}

func (l *LikeService) IsFavoriteCommentDict(ctx context.Context, in *pb.IsFavoriteCommentDictReq) (*pb.IsFavoriteCommentDictRsp, error) {
	isFavoriteDict := make(map[string]bool)
	for _, v := range in.FavoriteUnitList {
		exist, err := respository.IsFavoriteComment(v.UserId, v.CommentId)
		if err != nil {
			return nil, err
		}
		isFavoriteKey := strconv.FormatInt(v.UserId, 10) + "_" + strconv.FormatInt(v.CommentId, 10)
		isFavoriteDict[isFavoriteKey] = exist
	}
	return &pb.IsFavoriteCommentDictRsp{IsFavoriteDict: isFavoriteDict}, nil
}

func (l *LikeService) FavoriteCommentAction(ctx context.Context, in *pb.FavoriteCommentActionReq) (*pb.FavoriteCommentActionRsp, error) {
	if in.ActionType == constant.Like {
		err := respository.CommentLikeAction(in.CommentId, in.UserId)
		if err != nil {
			return nil, err
		}
	} else {
		err := respository.CommentUnLikeAction(in.CommentId, in.UserId)
		if err != nil {
			return nil, err
		}
	}
	return &pb.FavoriteCommentActionRsp{}, nil
}

func (l *LikeService) IsFavoriteVideoDict(ctx context.Context, in *pb.IsFavoriteVideoDictReq) (*pb.IsFavoriteVideoDictRsp, error) {
	isFavoriteDict := make(map[string]bool)
	for _, v := range in.FavoriteUnitList {
		exist, err := respository.IsFavoriteVideo(v.UserId, v.VideoId)
		if err != nil {
			return nil, err
		}
		isFavoriteKey := strconv.FormatInt(v.UserId, 10) + "_" + strconv.FormatInt(v.VideoId, 10)
		isFavoriteDict[isFavoriteKey] = exist
	}
	return &pb.IsFavoriteVideoDictRsp{IsFavoriteDict: isFavoriteDict}, nil
}

func (l *LikeService) FavoriteVideoAction(ctx context.Context, in *pb.FavoriteVideoActionReq) (*pb.FavoriteVideoActionRsp, error) {
	if in.ActionType == constant.Like {
		err := respository.VideoLikeAction(in.VideoId, in.UserId)
		if err != nil {
			return nil, err
		}
		return &pb.FavoriteVideoActionRsp{}, nil
	} else {
		err := respository.VideoUnLikeAction(in.VideoId, in.UserId)
		if err != nil {
			return nil, err
		}
		return &pb.FavoriteVideoActionRsp{}, nil
	}
}

func (l *LikeService) GetVideoLikeSum(ctx context.Context, in *pb.VideoLikeSumReq) (*pb.VideoLikeSumRsq, error) {
	sum, err := respository.VideoLikeNum(in.VideoId)
	if err != nil {
		return nil, err
	}
	return &pb.VideoLikeSumRsq{LikeNums: sum}, nil
}

func (l *LikeService) GetFavoriteVideoIdList(ctx context.Context, in *pb.GetFavoriteVideoIdListReq) (*pb.GetFavoriteVideoIdListRsp, error) {
	ids, err := respository.GetUserLikeVideos(in.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GetFavoriteVideoIdListRsp{VideoIdList: ids}, nil
}
