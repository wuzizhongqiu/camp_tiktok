package respository

import (
	"golang.org/x/net/context"
	"likesvr/constant"
	"likesvr/log"
	"likesvr/middlerware/cache"
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

func Exist(key string) (error, bool) {
	keyVideo := key + videoKeyPrefix
	rdb := cache.GetRdb()
	flag, err := rdb.Exists(context.Background(), keyVideo).Result()
	if err != nil {
		log.Error("Exist key==%s err=%v", key, err)
		return err, false
	}
	if flag == constant.KeyExist {
		return nil, true
	}
	return nil, false
}

// VideoLikeNumAddByCache  缓存视频的喜爱数
func VideoLikeNumAddByCache(vid string, uid string) error {
	DB := cache.GetRdb()
	VideoKey := videoKeyPrefix + vid
	err := setExpire(VideoKey)
	if err != nil {
		return err
	}
	_, err = DB.SAdd(context.Background(), VideoKey, uid).Result()
	if err != nil {
		log.Errorf("VideoCommentNumAddByCache err==%v", err)
		return err
	}
	return nil
}

// DelFromCache 如果出错时，从缓存删除
func DelFromCache(vid string) error {
	DB := cache.GetRdb()
	VideoKey := videoKeyPrefix + vid
	_, err := DB.Del(context.Background(), VideoKey).Result()
	if err != nil {
		log.Errorf(" DelFormCache del %s err=%v", vid, err)
		return err
	}
	return nil
}

func CacheGetVideoLikeNum(vid string) (int64, error) {
	DB := cache.GetRdb()
	VideoKey := videoKeyPrefix + vid
	err := setExpire(VideoKey)
	if err != nil {
		return 0, err
	}
	num, err := DB.SCard(context.Background(), VideoKey).Result()
	if err != nil {
		log.Errorf("CacheGetCommentNum err=%v", err)
		return 0, err
	}
	return num, nil
}

func DelMemberFromCache(vid string, uid string) error {
	DB := cache.GetRdb()
	VideoKey := videoKeyPrefix + vid
	err := setExpire(VideoKey)
	if err != nil {
		return err
	}
	_, err = DB.SRem(context.Background(), VideoKey, uid).Result()
	if err != nil {
		log.Errorf("DelMemberFromCache err==%v", err)
		return err
	}
	return nil
}

func ExistUserKey(uid string) (error, bool) {
	keyUser := uid + videoKeyPrefix
	rdb := cache.GetRdb()
	err := setExpire(keyUser)
	if err != nil {
		return err, false
	}
	flag, err := rdb.Exists(context.Background(), keyUser).Result()
	if err != nil {
		log.Error("Exist key==%s err=%v", uid, err)
		return err, false
	}
	if flag == constant.KeyExist {
		return nil, true
	}
	return nil, false
}

func UserLikeAddByCache(vid string, uid string) error {
	DB := cache.GetRdb()
	userKey := userKeyPrefix + uid
	err := setExpire(userKey)
	if err != nil {
		return err
	}
	_, err = DB.SAdd(context.Background(), userKey, vid).Result()
	if err != nil {
		log.Errorf("VideoCommentNumAddByCache err==%v", err)
		return err
	}
	return nil
}

func DelUserMemberFromCache(vid string, uid string) error {
	DB := cache.GetRdb()
	userKey := userKeyPrefix + uid
	err := setExpire(userKey)
	if err != nil {
		return err
	}
	_, err = DB.SRem(context.Background(), userKey, vid).Result()
	if err != nil {
		log.Errorf("DelMemberFromCache err==%v", err)
		return err
	}
	return nil
}

func DelUserFromCache(uid string) error {
	DB := cache.GetRdb()
	userKey := userKeyPrefix + uid
	_, err := DB.Del(context.Background(), userKey).Result()
	if err != nil {
		log.Errorf("  DelUserFromCache del %s err=%v", userKey, err)
		return err
	}
	return nil
}

func GetUserLikeListFromCache(uid string) ([]int64, error) {
	DB := cache.GetRdb()
	userKey := userKeyPrefix + uid
	err := setExpire(userKey)
	if err != nil {
		return nil, err
	}
	data, err := DB.SMembers(context.Background(), userKey).Result()
	if err != nil {
		log.Errorf(" GetUserLikeListFromCache %s err=%v", userKey, err)
		return nil, err
	}
	ids := make([]int64, 0)
	for _, s := range data {
		vid, _ := strconv.ParseInt(s, 10, 64)
		ids = append(ids, vid)
	}
	return ids, nil
}

func IsUserLikeVideoCheckByCache(uid string, vid string) (bool, error) {
	DB := cache.GetRdb()
	userKey := userKeyPrefix + uid
	err := setExpire(userKey)
	if err != nil {
		return false, err
	}
	flag, err := DB.SIsMember(context.Background(), userKey, vid).Result()
	if err != nil {
		return flag, err
	}
	return flag, nil
}

func ExistCommentKey(cid string) (bool, error) {
	DB := cache.GetRdb()
	commentKey := commentKeyPrefix + cid
	flag, err := DB.Exists(context.Background(), commentKey).Result()
	if err != nil {
		return false, err
	}
	if flag == constant.KeyExist {
		return true, nil
	}
	return false, nil
}

func CommentLikeNumAddByCache(cid string, uid string) error {
	DB := cache.GetRdb()
	commentKey := commentKeyPrefix + cid
	err := setExpire(commentKey)
	if err != nil {
		return err
	}
	_, err = DB.SAdd(context.Background(), commentKey, uid).Result()
	if err != nil {
		log.Errorf(" CommentLikeNumAddByCache err==%v", err)
		return err
	}
	return nil
}

func DelCommentFromCache(cid string) error {
	DB := cache.GetRdb()
	commentKey := commentKeyPrefix + cid
	_, err := DB.Del(context.Background(), commentKey).Result()
	if err != nil {
		log.Errorf("  DelCommentFromCache del %s err=%v", cid, err)
		return err
	}
	return nil
}

func CacheGetCommentLikeNum(cid string) (int64, error) {
	DB := cache.GetRdb()
	commentKey := commentKeyPrefix + cid
	err := setExpire(commentKey)
	if err != nil {
		return 0, err
	}
	num, err := DB.SCard(context.Background(), commentKey).Result()
	if err != nil {
		log.Errorf("CacheGetCommentNum err=%v", err)
		return 0, err
	}
	return num, nil
}

func DelCommentMemberFromCache(cid string, uid string) error {
	DB := cache.GetRdb()
	commentKey := commentKeyPrefix + cid
	err := setExpire(commentKey)
	if err != nil {
		return err
	}
	_, err = DB.SRem(context.Background(), commentKey, uid).Result()
	if err != nil {
		log.Errorf("DelCommentMemberFromCache err==%v", err)
		return err
	}
	return nil
}

func IsUserLikeCommentCheckByCache(uid string, cid string) (bool, error) {
	DB := cache.GetRdb()
	commentKey := commentKeyPrefix + cid
	err := setExpire(commentKey)
	if err != nil {
		return false, err
	}
	flag, err := DB.SIsMember(context.Background(), commentKey, uid).Result()
	if err != nil {
		return flag, err
	}
	return flag, nil
}
