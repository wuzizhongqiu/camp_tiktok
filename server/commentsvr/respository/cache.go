package respository

import (
	"commentsvr/constant"
	"commentsvr/log"
	"commentsvr/middlerware/cache"
	"golang.org/x/net/context"
)

// Exist 判读key值是否存在
func Exist(key string) (error, bool) {
	DB := cache.GetRdb()
	videoKey := videoKeyPrefix + key
	flag, err := DB.Exists(context.Background(), videoKey).Result()
	if err != nil {
		log.Errorf("Exist %s err==%v ", videoKeyPrefix, err)
		return err, false
	}
	if flag == constant.ExistInRedis {
		return nil, true
	}
	return nil, false
}

// VideoCommentNumAddByCache 缓存记录视频的评论数
func VideoCommentNumAddByCache(vid string, commentId string) error {
	DB := cache.GetRdb()
	VideoKey := videoKeyPrefix + vid
	_, err := DB.SAdd(context.Background(), VideoKey, commentId).Result()
	if err != nil {
		log.Errorf("VideoCommentNumAddByCache err==%v", err)
		return err
	}
	return nil
}

// DelFormCache 如果出错时，从缓存删除
func DelFormCache(vid string) error {
	DB := cache.GetRdb()
	VideoKey := videoKeyPrefix + vid
	_, err := DB.Del(context.Background(), VideoKey).Result()
	if err != nil {
		log.Errorf(" DelFormCache del %s err=%v", vid, err)
		return err
	}
	return nil
}

// CacheGetCommentNum 获取视频的评论数
func CacheGetCommentNum(vid string) (int64, error) {
	DB := cache.GetRdb()
	VideoKey := videoKeyPrefix + vid
	num, err := DB.SCard(context.Background(), VideoKey).Result()
	if err != nil {
		log.Errorf("CacheGetCommentNum err=%v", err)
		return 0, err
	}
	return num, nil
}

func DelMemberFromCache(vid string, commentId string) error {
	DB := cache.GetRdb()
	VideoKey := videoKeyPrefix + vid
	_, err := DB.SRem(context.Background(), VideoKey, commentId).Result()
	if err != nil {
		log.Errorf("DelMemberFromCache err==%v", err)
		return err
	}
	return nil
}
