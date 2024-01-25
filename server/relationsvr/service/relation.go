package service

import (
	"github.com/Keqing-win/camp_tiktok/pkg/pb"
	"golang.org/x/net/context"
	"relationsvr/constant"
	"relationsvr/repository"
	"strconv"
)

type RelationService struct {
	pb.UnimplementedRelationServiceServer
}

func (r *RelationService) RelationAction(ctx context.Context, in *pb.RelationActionReq) (*pb.RelationActionRsp, error) {
	if in.ActionType == constant.Follow {
		err := repository.RelationAdd(in.SelfUserId, in.ToUserId)
		if err != nil {
			return nil, err
		}
	} else {
		err := repository.RelationDel(in.SelfUserId, in.ToUserId)
		if err != nil {
			return nil, err
		}
	}
	return &pb.RelationActionRsp{}, nil
}

func (r *RelationService) GetRelationFollowList(ctx context.Context, in *pb.GetRelationFollowListReq) (*pb.GetRelationFollowListRsp, error) {
	ids, err := repository.RelationFollowList(in.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GetRelationFollowListRsp{FollowList: ids}, nil
}

func (r *RelationService) GetRelationFollowerList(ctx context.Context, in *pb.GetRelationFollowerListReq) (*pb.GetRelationFollowerListRsp, error) {
	ids, err := repository.RelationFollowerList(in.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GetRelationFollowerListRsp{FollowerList: ids}, nil
}

func (r *RelationService) IsFollowDict(ctx context.Context, in *pb.IsFollowDictReq) (*pb.IsFollowDictRsp, error) {
	isFollowDict := make(map[string]bool)
	for _, v := range in.FollowUintList {
		flag, err := repository.RelationIsFollow(v.SelfUserId, v.UserIdList)
		if err != nil {
			return nil, err
		}
		isFollowKey := strconv.FormatInt(v.SelfUserId, 10) + "_" + strconv.FormatInt(v.UserIdList, 10)
		isFollowDict[isFollowKey] = flag
	}
	return &pb.IsFollowDictRsp{IsFollowDict: isFollowDict}, nil
}

func (r *RelationService) GetUserFollowerNum(ctx context.Context, in *pb.UserFollowerNumReq) (*pb.UserFollowerNumRsp, error) {
	sum, err := repository.FollowerNum(in.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.UserFollowerNumRsp{Sum: sum}, err
}

func (r *RelationService) GetUserFollowNum(ctx context.Context, in *pb.UserFollowNumReq) (*pb.UserFollowNumRsp, error) {
	sum, err := repository.FollowNum(in.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.UserFollowNumRsp{Sum: sum}, nil
}
