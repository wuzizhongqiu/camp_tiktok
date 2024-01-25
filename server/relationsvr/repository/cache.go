package repository

import (
	"golang.org/x/net/context"
	"relationsvr/constant"
	"relationsvr/log"
	"relationsvr/middlerware/cache"
	"strconv"
)

func setExpire(key string) error {
	rdb := cache.GetRdb()
	_, err := rdb.Expire(context.Background(), key, constant.KeyExpire).Result()
	if err != nil {
		log.Infof("setExpire %s err==%v", key, err)
		return err
	}
	return nil
}

// ExistFollowerKey 检查记录粉丝的key是否存在
func ExistFollowerKey(key string) (error, bool) {
	keyFollower := key + followerKeyPrefix
	rdb := cache.GetRdb()
	flag, err := rdb.Exists(context.Background(), keyFollower).Result()
	if err != nil {
		log.Error("ExistFollowerKey key==%s err=%v", key, err)
		return err, false
	}
	if flag == constant.KeyExist {
		return nil, true
	}
	return nil, false
}

// FollowerNumAddByCache 缓存 粉丝关系
func FollowerNumAddByCache(uid string, followerId string) error {
	DB := cache.GetRdb()
	keyFollower := uid + followerKeyPrefix
	err := setExpire(keyFollower)
	if err != nil {
		return err
	}
	_, err = DB.SAdd(context.Background(), keyFollower, followerId).Result()
	if err != nil {
		log.Errorf("FollowerNumAddByCache err==%v", err)
		return err
	}
	return nil
}

// DelFollowerFromCache 如果出错时，从缓存删除
func DelFollowerFromCache(uid string) {
	DB := cache.GetRdb()
	keyFollower := uid + followerKeyPrefix
	_, err := DB.Del(context.Background(), keyFollower).Result()
	if err != nil {
		log.Errorf(" DelFollowerFromCache del %s err=%v", uid, err)
	}
}

func GetFollowerListByCache(uid string) ([]int64, error) {
	DB := cache.GetRdb()
	keyFollower := uid + followerKeyPrefix
	data, err := DB.SMembers(context.Background(), keyFollower).Result()
	if err != nil {
		log.Errorf("GetFollowerListByCache %s err==%v", uid, err)
		return nil, err
	}
	ids := make([]int64, 0)
	for _, s := range data {
		vid, _ := strconv.ParseInt(s, 10, 64)
		ids = append(ids, vid)
	}
	return ids, nil
}

func CacheGetFollowerNum(uid string) (int64, error) {
	DB := cache.GetRdb()
	keyFollower := uid + followerKeyPrefix
	err := setExpire(keyFollower)
	if err != nil {
		return 0, err
	}
	num, err := DB.SCard(context.Background(), keyFollower).Result()
	if err != nil {
		log.Errorf("CacheGetFollowerNum err=%v", err)
		return 0, err
	}
	return num, nil
}

func DelFollowerMemberFromCache(uid string, followerId string) error {
	DB := cache.GetRdb()
	keyFollower := uid + followerKeyPrefix
	err := setExpire(keyFollower)
	if err != nil {
		return err
	}
	_, err = DB.SRem(context.Background(), keyFollower, followerId).Result()
	if err != nil {
		log.Errorf("DelFollowerMemberFromCache err==%v", err)
		return err
	}
	return nil
}

// ExistFollowKey 缓存关注关系
func ExistFollowKey(uid string) (error, bool) {
	keyFollow := uid + followKeyPrefix
	rdb := cache.GetRdb()
	err := setExpire(keyFollow)
	if err != nil {
		return err, false
	}
	flag, err := rdb.Exists(context.Background(), keyFollow).Result()
	if err != nil {
		log.Error("ExistFollowKey key==%s err=%v", uid, err)
		return err, false
	}
	if flag == constant.KeyExist {
		return nil, true
	}
	return nil, false
}

func FollowAddByCache(uid string, target string) error {
	DB := cache.GetRdb()
	keyFollow := uid + followKeyPrefix
	err := setExpire(keyFollow)
	if err != nil {
		return err
	}
	_, err = DB.SAdd(context.Background(), keyFollow, target).Result()
	if err != nil {
		log.Errorf("FollowAddByCache err==%v", err)
		return err
	}
	return nil
}

func DelFollowMemberFromCache(vid string, uid string) error {
	DB := cache.GetRdb()
	keyFollow := uid + followKeyPrefix
	err := setExpire(keyFollow)
	if err != nil {
		return err
	}
	_, err = DB.SRem(context.Background(), keyFollow, vid).Result()
	if err != nil {
		log.Errorf(" DelFollowMemberFromCache err==%v", err)
		return err
	}
	return nil
}

func DelFollowFromCache(uid string) {
	DB := cache.GetRdb()
	keyFollow := uid + followKeyPrefix
	_, err := DB.Del(context.Background(), keyFollow).Result()
	if err != nil {
		log.Errorf(" DelFollowFromCache del %s err=%v", keyFollow, err)
	}
}

func GetFollowListFromCache(uid string) ([]int64, error) {
	DB := cache.GetRdb()
	keyFollow := uid + followKeyPrefix
	err := setExpire(keyFollow)
	if err != nil {
		return nil, err
	}
	data, err := DB.SMembers(context.Background(), keyFollow).Result()
	if err != nil {
		log.Errorf(" GetFollowListFromCache %s err=%v", keyFollow, err)
		return nil, err
	}
	ids := make([]int64, 0)
	for _, s := range data {
		vid, _ := strconv.ParseInt(s, 10, 64)
		ids = append(ids, vid)
	}
	return ids, nil
}

func CacheGetFollowNum(uid string) (int64, error) {
	DB := cache.GetRdb()
	keyFollow := uid + followKeyPrefix
	err := setExpire(keyFollow)
	if err != nil {
		return 0, err
	}
	num, err := DB.SCard(context.Background(), keyFollow).Result()
	if err != nil {
		log.Errorf("CacheGetFollowerNum err=%v", err)
		return 0, err
	}
	return num, nil
}

func IsFollowCheckByCache(uid string, target string) (bool, error) {
	DB := cache.GetRdb()
	keyFollow := uid + followKeyPrefix
	err := setExpire(keyFollow)
	if err != nil {
		return false, err
	}
	flag, err := DB.SIsMember(context.Background(), keyFollow, target).Result()
	if err != nil {
		log.Errorf("IsUserLikeVideoCheckByCache err=%v", err)
		return flag, err
	}
	return flag, nil
}
