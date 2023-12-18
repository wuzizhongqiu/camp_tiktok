package service

import (
	"errors"
	"github.com/Keqing-win/camp_tiktok/pkg/pb"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"strconv"
	"usersvr/middlerware/lock"
	"usersvr/repository"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (u *UserService) UpdateUserFollowerCount(ctx context.Context, in *pb.UpdateUserFollowerCountReq) (*pb.UpdateUserFollowerCountRsp, error) {
	err := repository.UpdateUserFollowerNum(in.UserId, in.ActionType)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserFollowerCountRsp{}, nil
}

func (u *UserService) UpdateUserFollowCount(ctx context.Context, in *pb.UpdateUserFollowCountReq) (*pb.UpdateUserFollowCountRsp, error) {
	err := repository.UpdateUserFollowNum(in.UserId, in.ActionType)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserFollowCountRsp{}, nil
}

func (u *UserService) UpdateUserFavoriteCount(ctx context.Context, in *pb.UpdateUserFavoriteCountReq) (*pb.UpdateUserFavoriteCountRsp, error) {
	err := repository.UpdateUserFavoriteNum(in.UserId, in.ActionType)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserFavoriteCountRsp{}, nil
}

func (u *UserService) UpdateUserFavoritedCount(ctx context.Context, in *pb.UpdateUserFavoritedCountReq) (*pb.UpdateUserFavoritedCountRsp, error) {
	err := repository.UpdateUserFavoritedNum(in.UserId, in.ActionType)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserFavoritedCountRsp{}, nil
}

func (u *UserService) CacheChangeUserCount(ctx context.Context, in *pb.CacheChangeUserCountReq) (*pb.CacheChangeUserCountRsp, error) {
	uid := strconv.FormatInt(in.UserId, 10)
	//保证同一时间只有一个操作
	mutex := lock.GetRedSync("user_" + uid)
	defer lock.Unlock(mutex)
	user, err := repository.CacheGetUserInfo(in.UserId)
	if err != nil {
		return nil, err
	}
	switch in.CountType {
	case "follow":
		user.Follow += in.Op
	case "follower":
		user.Follower += in.Op
	case "liked":
		user.TotalFav += in.Op
	case "like":
		user.FavCount += in.Op
	}
	repository.CacheSetUserInfo(user)
	return &pb.CacheChangeUserCountRsp{}, nil
}

func (u *UserService) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	flag, err := repository.UserIsExist(in.Username)
	if err != nil {
		return nil, err
	}
	if flag {
		return nil, errors.New("user exist")
	}
	var pass string
	pass, err = repository.PassWordHash(in.Password)
	if err != nil {
		return nil, err
	}
	var user repository.User
	user, err = repository.InsertUser(in.Username, pass)
	if err != nil {
		return nil, err
	}
	var token string
	token, err = repository.GenerateToken(user.Id, user.Name)
	if err != nil {
		return nil, err
	}
	resp := &pb.RegisterResponse{
		UserId: user.Id,
		Token:  token,
	}
	return resp, nil
}

func (u *UserService) CheckPassWord(ctx context.Context, in *pb.CheckPassWordRequest) (*pb.CheckPassWordResponse, error) {
	user, err := repository.GetUserInfo(in.Username)
	if err != nil {
		return nil, err
	}
	//校验密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password))
	if err != nil {
		return nil, errors.New("password error")
	}
	//生成token
	token, err := repository.GenerateToken(user.Id, user.Name)
	if err != nil {
		return nil, err
	}
	resp := &pb.CheckPassWordResponse{
		UserId: user.Id,
		Token:  token,
	}
	return resp, nil
}

func (u *UserService) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	user, err := repository.GetUserInfo(req.Id)
	if err != nil {
		return nil, err
	}

	resp := &pb.GetUserInfoResponse{UserInfo: UserToUserInfo(user)}
	return resp, nil
}

func (u *UserService) GetUserInfoDict(ctx context.Context, in *pb.GetUserInfoDictRequest) (*pb.GetUserInfoDictResponse, error) {
	users, err := repository.GetUserInfoList(in.UserIdList)
	if err != nil {
		return nil, err
	}
	tmp := make(map[int64]*pb.UserInfo)
	for _, user := range users {
		tmp[user.Id] = &pb.UserInfo{
			Id:              user.Id,
			Name:            user.Name,
			FollowCount:     user.Follow,
			FollowerCount:   user.Follower,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFav,
			FavoriteCount:   user.FavCount,
		}
	}
	resp := &pb.GetUserInfoDictResponse{UserInfoDict: tmp}
	return resp, nil
}

func UserToUserInfo(info repository.User) *pb.UserInfo {
	return &pb.UserInfo{
		Id:              info.Id,
		Name:            info.Name,
		FollowCount:     info.Follow,
		FollowerCount:   info.Follower,
		IsFollow:        false,
		Avatar:          info.Avatar,
		BackgroundImage: info.BackgroundImage,
		Signature:       info.Signature,
		TotalFavorited:  info.TotalFav,
		FavoriteCount:   info.FavCount,
	}
}
