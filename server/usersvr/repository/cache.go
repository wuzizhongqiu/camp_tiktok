package repository

import (
	"github.com/go-redis/redis"
	"golang.org/x/net/context"
	"strconv"
	"usersvr/log"
	"usersvr/middlerware/cache"
)

// CacheCheckUser 检查key是否失效
func CacheCheckUser(uid int64) (error, bool) {
	rdb := cache.GetRdb()
	userKey := userKeyPrefix + strconv.FormatInt(uid, 10)
	_, err := rdb.HGetAll(context.Background(), userKey).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, false
		}
		log.Errorf("CacheCheckUser err==%v", err)
		return err, false
	}
	return nil, true
}

func CacheUpdateFollowerNum(uid int64, num int64) error {
	rdb := cache.GetRdb()
	userKey := userKeyPrefix + strconv.FormatInt(uid, 10)
	err := rdb.HIncrBy(context.Background(), userKey, "follower_count", num).Err()
	if err != nil {
		log.Errorf("CacheUpdateFollowerNum err====%v ", err)
		return err
	}
	return nil
}

func CacheUpdateFollowNum(uid int64, num int64) error {
	rdb := cache.GetRdb()
	userKey := userKeyPrefix + strconv.FormatInt(uid, 10)
	err := rdb.HIncrBy(context.Background(), userKey, "follow_count", num).Err()
	if err != nil {
		log.Errorf("CacheUpdateFollowNum err====%v ", err)
		return err
	}
	return nil
}

func CacheUpdateFavoriteNum(uid int64, num int64) error {
	rdb := cache.GetRdb()
	userKey := userKeyPrefix + strconv.FormatInt(uid, 10)
	err := rdb.HIncrBy(context.Background(), userKey, "favorite_count", num).Err()
	if err != nil {
		log.Errorf("CacheUpdateFavoriteNum err====%v ", err)
		return err
	}
	return nil
}

func CacheUpdateFavoritedNum(uid int64, num int64) error {
	rdb := cache.GetRdb()
	userKey := userKeyPrefix + strconv.FormatInt(uid, 10)
	err := rdb.HIncrBy(context.Background(), userKey, "total_favorited", num).Err()
	if err != nil {
		log.Errorf("CacheUpdateFavoritedNum err====%v ", err)
		return err
	}
	return nil
}

func CacheGetUserInfo(uid int64) (User, error) {
	rdb := cache.GetRdb()
	userKey := userKeyPrefix + strconv.FormatInt(uid, 10)
	var user User
	data, err := rdb.HGetAll(context.Background(), userKey).Result()
	if err != nil {
		log.Errorf("cache GetUserInfo err:%v", err)
		return user, err
	}
	if len(data) == 0 {
		log.Errorf("cache GetUserInfo err:%v", err)
		return user, err
	}
	user.Id, _ = strconv.ParseInt(data["id"], 10, 64)
	user.Name = data["user_name"]
	user.Password = data["password"]
	user.Follow, _ = strconv.ParseInt(data["follow_count"], 10, 64)
	user.Follower, _ = strconv.ParseInt(data["follower_count"], 10, 64)
	user.Avatar = data["avatar"]
	user.BackgroundImage = data["background_image"]
	user.Signature = data["signature"]
	user.TotalFav, _ = strconv.ParseInt(data["total_favorited"], 10, 64)
	user.FavCount, _ = strconv.ParseInt(data["favorite_count"], 10, 64)
	return user, nil
}

func CacheSetUserInfo(u User) {
	rdb := cache.GetRdb()
	userKey := userKeyPrefix + strconv.FormatInt(u.Id, 10)
	if err := rdb.HSet(context.Background(), userKey, map[string]interface{}{
		"id":               u.Id,
		"user_name":        u.Name,
		"password":         u.Password,
		"follow_count":     u.Follow,
		"follower_count":   u.Follower,
		"avatar":           u.Avatar,
		"background_image": u.BackgroundImage,
		"signature":        u.Signature,
		"total_favorited":  u.TotalFav,
		"favorite_count":   u.FavCount,
	}).Err(); err != nil {
		log.Errorf("CacheSetUserInfo err===%v", err)
		return
	}
}
