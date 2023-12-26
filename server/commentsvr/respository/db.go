package respository

import (
	"commentsvr/constant"
	"commentsvr/log"
	"commentsvr/middlerware/db"
	"gorm.io/gorm"
	"time"
)

func GetCommentIdsByVideoId(videoId int64) ([]int64, error) {
	Db := db.GetDb()
	var ids []int64
	err := Db.Model(&Comment{}).Where("video_id=?", videoId).Pluck("id", &ids).Error
	if err != nil {
		log.Errorf("GetCommentIdsByVideoId err==%v", err)
		return nil, err
	}
	return ids, nil
}

func GetCommentSumByDB(videoId int64) (int64, error) {
	Db := db.GetDb()
	var count int64
	if err := Db.Model(&Comment{}).Where("video_id=?", videoId).Count(&count).Error; err != nil {
		log.Errorf("GetCommentSumByDB err==%v", err)
		return 0, err
	}
	return count, nil
}

// GetUserId 获取uid
func GetUserId(commentId int64) (int64, error) {
	Db := db.GetDb()
	var comment Comment
	if err := Db.Where("id=?", commentId).First(&comment).Error; err != nil {
		log.Errorf(" GetUserId err====%v", err)
		return 0, err
	}
	return comment.UserId, nil
}

// GetReplyId 获取到拉去二级评论的恢复id
func GetReplyId(commentId int64) (int64, error) {
	Db := db.GetDb()
	var reply CommentReply
	if err := Db.Where("comment_id=?", commentId).First(&reply).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		log.Errorf(" GetReplyId err====%v", err)
		return 0, err
	}
	return reply.ReplyId, nil
}

func GetOtherCommentListByDB(vid, parentId int64, offset int) ([]*Comment, error) {
	var comments []*Comment
	Db := db.GetDb()
	//先获取到评论，避免深分页
	//注意 preload传入的是要进行预查询的定义结构体中的字段名，不是数据表的名字
	//err := Db.Model(&Comment{}).
	//	Where("id (IN) ?", Db.Select("id").Where(&Comment{VideoId: vid, ParentId: parentId, CommentLevel: constant.OtherLevel}).Offset(offset).Limit(constant.DefaultLimit).
	//		Order("favorite_num desc,create_time desc")).
	//	Find(&comments).Error
	var commentIDs []int
	err := Db.Model(&Comment{}).
		Select("id").
		Where(&Comment{VideoId: vid, ParentId: parentId, CommentLevel: constant.OtherLevel}).
		Offset(offset).
		Limit(constant.DefaultLimit).
		Order("favorite_num desc, create_time desc").
		Pluck("id", &commentIDs).Error
	if err != nil {
		log.Errorf("GetOtherCommentListByDB err===%v", err)
		return nil, err
	}
	err = Db.Model(&Comment{}).
		Where("id IN (?)", commentIDs).
		Find(&comments).
		Error
	if err != nil {
		log.Errorf("GetOtherCommentListByDB err===%v", err)
		return nil, err
	}
	return comments, nil
}

func GetTopCommentListByDB(vid int64, offset int) ([]*Comment, error) {
	var comments []*Comment
	Db := db.GetDb()
	//根据level和vid获取到评论,避免深分页,使用子查询（子查询获取到id）版本不支持子查询limit
	//err := Db.Model(&Comment{}).
	//	Where("id IN (?)", Db.Table("t_comment").
	//		Select("id").Order("favorite_num desc,create_time desc").Where(&Comment{VideoId: vid, CommentLevel: constant.FirstLevel}).Offset(offset).Limit(constant.DefaultLimit)).
	//	Find(&comments).Error
	var tempComments []*Comment
	err := Db.Table("t_comment").
		Select("id").
		Order("favorite_num desc,create_time desc").
		Where(&Comment{VideoId: vid, CommentLevel: constant.FirstLevel}).
		Offset(offset).
		Limit(constant.DefaultLimit).
		Scan(&tempComments).Error
	if err != nil {
		log.Errorf("GetTopCommentList err===%v", err)
		return nil, err
	}
	var ids []int64
	for _, comment := range tempComments {
		ids = append(ids, comment.Id)
	}
	err = Db.Model(&Comment{}).
		Where("id IN (?)", ids).
		Find(&comments).Error
	if err != nil {
		log.Errorf("GetTopCommentList err===%v", err)
		return nil, err
	}
	return comments, nil
}

func GetTopCommentReplyNumByDB(commentId int64) (int64, error) {
	Db := db.GetDb()
	var count int64
	err := Db.Model(&Comment{}).Where("parent_id = ?", commentId).Count(&count).Error
	if err != nil {
		log.Errorf("GetTopCommentReplyNumByDB err===%v", err)
		return 0, err
	}
	return count, nil
}

func DeleteOtherCommentByDB(commentId int64) error {
	//在comment表软删除
	Db := db.GetDb()
	err := Db.Transaction(func(tx *gorm.DB) error {
		//现在comment表软删除
		err := tx.Model(&Comment{}).Where("id=?", commentId).Update("delete", constant.Delete).Error
		if err != nil {
			log.Errorf(" DeleteOtherCommentByDB comment del err====%v", err)
			return err
		}
		//查询commentId对应的reply_id然后删除
		var ids []int64
		err = tx.Model(&CommentReply{}).Select("reply_id").Where("comment_id=?", commentId).Scan(&ids).Error
		if err != nil {
			log.Errorf(" DeleteOtherCommentByDB reply_id  find err====%v", err)
			return err
		}
		commentReply := CommentReply{}
		err = tx.Model(&CommentReply{}).Where("comment_id=?", commentId).Delete(&commentReply).Error
		if err != nil {
			log.Errorf(" DeleteOtherCommentByDB  delete reply_id err====%v", err)
			return err
		}
		//批量删除reply_id对应的comment,软删除
		err = tx.Model(&Comment{}).Where("id in ?", ids).Update("delete", constant.Delete).Error
		if err != nil {
			log.Errorf(" DeleteOtherCommentByDB delete commentOther err===%v", err)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil

}

func CreateOtherCommentByDB(vid, uid, parentId, replyId, parentUid int64, text string) (*Comment, error) {
	Db := db.GetDb()
	comment := Comment{
		UserId:       uid,
		VideoId:      vid,
		ParentId:     parentId,
		ParentUserId: parentUid,
		CommentLevel: constant.OtherLevel,
		Delete:       0,
		CreateTime:   time.Now(),
		CommentText:  text,
	}
	commentReply := CommentReply{
		ReplyId: replyId,
	}
	err := Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&comment).Error
		if err != nil {
			log.Errorf("CreateOtherCommentByDB create comment err==%v", err)
			return err
		}
		commentReply.CommentId = comment.Id
		err = tx.Create(&commentReply).Error
		if err != nil {
			log.Errorf("CreateOtherCommentByDB create comment_reply err==%v", err)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func CreateTopCommentByDB(uid, vid int64, text string) (*Comment, error) {
	Db := db.GetDb()
	comment := Comment{
		UserId:       uid,
		VideoId:      vid,
		ParentId:     0,
		ParentUserId: 0,
		CommentLevel: constant.FirstLevel,
		Delete:       0, //默认不删除
		CreateTime:   time.Now(),
		CommentText:  text,
	}
	err := Db.Model(&Comment{}).Create(&comment).Error
	if err != nil {
		log.Errorf("CreateTopCommentByDB err===%v", err)
		return nil, err
	}
	return &comment, nil
}

// DeleteTopCommentByDB 软删除
func DeleteTopCommentByDB(commentId int64) error {
	Db := db.GetDb()
	err := Db.Model(&Comment{}).Where("id=?", commentId).Update("delete", constant.Delete).Error
	if err != nil {
		log.Errorf("DeleteTopCommentByDB err==%v", err)
		return err
	}
	return nil
}
