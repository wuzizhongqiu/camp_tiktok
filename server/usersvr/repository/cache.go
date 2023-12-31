package repository

import (
	"golang.org/x/net/context"
	"strconv"
	"usersvr/constant"
	"usersvr/log"
	"usersvr/middlerware/cache"
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

// CacheCheckUser 检查key是否失效
func CacheCheckUser(uid int64) (bool, error) {
	rdb := cache.GetRdb()
	userKey := userKeyPrefix + strconv.FormatInt(uid, 10)
	data, err := rdb.HGetAll(context.Background(), userKey).Result()
	if err != nil {
		log.Errorf("CacheCheckUser err: %v", err)
		return false, err
	}
	if len(data) == 0 {
		return false, nil
	}
	return true, nil
}

func CacheUpdateFollowerNum(uid int64, num int64) error {
	rdb := cache.GetRdb()
	userKey := userKeyPrefix + strconv.FormatInt(uid, 10)
	err := setExpire(userKey)
	if err != nil {
		return err
	}
	err = rdb.HIncrBy(context.Background(), userKey, "follower_count", num).Err()
	if err != nil {
		log.Errorf("CacheUpdateFollowerNum err====%v ", err)
		return err
	}
	return nil
}

func CacheUpdateFollowNum(uid int64, num int64) error {
	rdb := cache.GetRdb()
	userKey := userKeyPrefix + strconv.FormatInt(uid, 10)
	err := setExpire(userKey)
	if err != nil {
		return err
	}
	err = rdb.HIncrBy(context.Background(), userKey, "follow_count", num).Err()
	if err != nil {
		log.Errorf("CacheUpdateFollowNum err====%v ", err)
		return err
	}
	return nil
}

func CacheUpdateFavoriteNum(uid int64, num int64) error {
	rdb := cache.GetRdb()
	userKey := userKeyPrefix + strconv.FormatInt(uid, 10)
	err := setExpire(userKey)
	if err != nil {
		return err
	}
	err = rdb.HIncrBy(context.Background(), userKey, "favorite_count", num).Err()
	if err != nil {
		log.Errorf("CacheUpdateFavoriteNum err====%v ", err)
		return err
	}
	return nil
}

func CacheUpdateFavoritedNum(uid int64, num int64) error {
	rdb := cache.GetRdb()
	userKey := userKeyPrefix + strconv.FormatInt(uid, 10)
	err := setExpire(userKey)
	if err != nil {
		return err
	}
	err = rdb.HIncrBy(context.Background(), userKey, "total_favorited", num).Err()
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
	err := setExpire(userKey)
	if err != nil {
		return user, err
	}
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
	user.UserName = data["user_name"]
	user.Password = data["password"]
	user.FollowCount, _ = strconv.ParseInt(data["follow_count"], 10, 64)
	user.FollowerCount, _ = strconv.ParseInt(data["follower_count"], 10, 64)
	user.Avatar = data["avatar"]
	user.BackgroundImage = data["background_image"]
	user.Signature = data["signature"]
	user.TotalFavorited, _ = strconv.ParseInt(data["total_favorited"], 10, 64)
	user.FavoriteCount, _ = strconv.ParseInt(data["favorite_count"], 10, 64)
	return user, nil
}

func CacheSetUserInfo(u User) {
	rdb := cache.GetRdb()
	userKey := userKeyPrefix + strconv.FormatInt(u.Id, 10)
	setExpire(userKey)

	log.Infof("userKey =%s", userKey)
	log.Infof("%+v", u)
	if err := rdb.HMSet(context.Background(), userKey, map[string]interface{}{
		"id":               u.Id,
		"user_name":        u.UserName,
		"password":         u.Password,
		"follow_count":     u.FollowCount,
		"follower_count":   u.FollowerCount,
		"avatar":           u.Avatar,
		"background_image": u.BackgroundImage,
		"signature":        u.Signature,
		"total_favorited":  u.TotalFavorited,
		"favorite_count":   u.FavoriteCount,
	}).Err(); err != nil {
		log.Errorf("CacheSetUserInfo err===%v", err)
		log.Infof("userKey =%s", userKey)
		return
	}
}
