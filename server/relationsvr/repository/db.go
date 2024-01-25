package repository

import (
	"gorm.io/gorm"
	"relationsvr/constant"
	"relationsvr/log"
	"relationsvr/middlerware/db"
)

func DeleteRelationByDb(uid, target int64) error {
	DB := db.GetDb()
	err := DB.Model(&Relation{}).Where(&Relation{UserId: target, FollowerId: uid, Delete: constant.NotDelete}).Update("delete", constant.Delete).Error
	if err != nil {
		log.Errorf("DeleteRelationByDb err==%v", err)
	}
	return err
}

func InsertRelationToDb(uid, target int64) error {
	DB := db.GetDb()
	err := DB.Transaction(func(tx *gorm.DB) error {
		//先查询有无关注关系
		var relation Relation
		err := tx.Where(&Relation{UserId: target, FollowerId: uid, Delete: constant.Delete}).First(&relation).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				//不存在就进行插入操作
				relation.UserId = target
				relation.FollowerId = uid
				relation.Delete = constant.NotDelete
				err = tx.Create(&relation).Error
				if err != nil {
					log.Errorf("InsertRelationToDb err===%v", err)
					return err
				}
				return nil
			}
			log.Errorf("InsertRelationToDb err===%v", err)
			return err
		}
		//有就进行更新操作
		err = tx.Model(&Relation{}).Where(&Relation{UserId: target, FollowerId: uid, Delete: constant.Delete}).Update("delete", constant.NotDelete).Error
		if err != nil {
			log.Errorf("InsertRelationToDb err===%v", err)
		}
		return err
	})
	return err
}

func GetFollowListByDB(uid int64) ([]int64, error) {
	var ids []int64
	DB := db.GetDb()
	err := DB.Model(&Relation{}).Where(&Relation{FollowerId: uid, Delete: constant.NotDelete}).Pluck("user_id", &ids).Error
	if err != nil {
		log.Errorf(" GetFollowListByDB err==%v", err)
		return nil, err
	}
	return ids, nil
}

func GetFollowerListByDB(uid int64) ([]int64, error) {
	var ids []int64
	DB := db.GetDb()
	err := DB.Model(&Relation{}).Where(&Relation{UserId: uid, Delete: constant.NotDelete}).Pluck("follower_id", &ids).Error
	if err != nil {
		log.Errorf("GetFollowerListByDB err==%v", err)
		return nil, err
	}
	return ids, nil
}
