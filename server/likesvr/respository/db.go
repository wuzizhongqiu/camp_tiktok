package respository

import (
	"gorm.io/gorm"
	"likesvr/constant"
	"likesvr/log"
	"likesvr/middlerware/db"
)

func IsUserLikCommentCheckByDB(cid, uid int64) (bool, error) {
	DB := db.GetDb()
	var like CommentLike
	err := DB.Where(&CommentLike{CommentId: cid, UserId: uid, Delete: constant.NotDelete}).First(&like).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func DeleteCommentLike(cid, uid int64) error {
	DB := db.GetDb()
	err := DB.Model(&CommentLike{}).Where(&CommentLike{CommentId: cid, UserId: uid}).Update("delete", constant.Delete).Error
	if err != nil {
		log.Errorf("DeleteCommentLike err===%v", err)
		return err
	}
	return nil
}

func InsertCommentLike(cid, uid int64) error {
	var comment CommentLike
	comment.CommentId = cid
	comment.UserId = uid
	comment.Delete = constant.NotDelete
	DB := db.GetDb()
	err := DB.Create(&comment).Error
	if err != nil {
		log.Errorf("InsertCommentLike insert err==%v", err)
		return err
	}
	return nil
}

func GetUserIdListByDB(cid int64) ([]int64, error) {
	var ids []int64
	DB := db.GetDb()
	err := DB.Model(&CommentLike{}).Where(&CommentLike{CommentId: cid, Delete: constant.NotDelete}).Pluck("user_id", &ids).Error
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func InsertVideoLike(vid, uid int64) error {
	var like VideoLike
	like.VideoId = vid
	like.UserId = uid
	like.Delete = constant.NotDelete
	DB := db.GetDb()
	err := DB.Create(&like).Error
	if err != nil {
		log.Errorf("VideoLikeAction insert err==%v", err)
		return err
	}
	return nil
}

func IsUserLikeVideoCheckByDB(vid, uid int64) (bool, error) {
	DB := db.GetDb()
	var like VideoLike
	err := DB.Model(&VideoLike{}).Where(&VideoLike{VideoId: vid, UserId: uid, Delete: constant.NotDelete}).First(&like).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func GetVideoLikeList(vid int64) ([]int64, error) {
	DB := db.GetDb()
	var ids []int64
	err := DB.Model(&VideoLike{}).Where(&VideoLike{VideoId: vid, Delete: constant.NotDelete}).Pluck("user_id", &ids).Error
	if err != nil {
		log.Errorf(" GetVideoLikeList err==%v ", err)
		return nil, err
	}
	return ids, nil
}

func GetUserLikeList(uid int64) ([]int64, error) {
	DB := db.GetDb()
	var ids []int64
	err := DB.Model(&VideoLike{}).Where(&VideoLike{UserId: uid, Delete: constant.NotDelete}).Pluck("video_id", &ids).Error
	if err != nil {
		log.Errorf(" GetUserLikeList err==%v ", err)
		return nil, err
	}
	return ids, nil
}

// DeleteVideoLike 软删除
func DeleteVideoLike(vid, uid int64) error {
	DB := db.GetDb()
	err := DB.Model(&VideoLike{}).Where(&VideoLike{VideoId: vid, UserId: uid}).Update("delete", constant.Delete).Error
	if err != nil {
		log.Errorf("DeleteVideoLike err===%v", err)
		return err
	}
	return nil
}
