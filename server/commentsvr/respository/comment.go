package respository

import (
	"commentsvr/log"
	"golang.org/x/sync/singleflight"
	"strconv"
)

// ReplyComment 回复评论的详细内容,在获取到所有二级评论后对其中的回复评论进行封装
type ReplyComment struct {
	Comment     Comment
	ReplyId     int64
	ReplyUserId int64
}

var gsf singleflight.Group

func CommentLikeAdd(cid int64, num int64) error {
	err := CommentLikeAddByDB(cid, num)
	return err
}

func GetCommentSum(vid int64) (int64, error) {
	sVid := strconv.FormatInt(vid, 10)
	err, flag := Exist(sVid)
	if err != nil {
		return 0, err
	}
	//如果不存在，则从数据库更新
	if !flag {
		ids, err, _ := gsf.Do(strconv.FormatInt(vid, 10), func() (interface{}, error) {
			return GetCommentIdsByVideoId(vid)
		})
		if err != nil {
			return 0, err
		}
		for _, id := range ids.([]int64) {
			err := VideoCommentNumAddByCache(sVid, strconv.FormatInt(id, 10))
			if err != nil {
				err = DelFormCache(sVid)
				return 0, err
			}
		}
		//更新后，从cache拿数据
		sum, err := CacheGetCommentNum(sVid)
		if err != nil {
			return 0, err
		}
		return sum, nil
	} else {
		sum, err := CacheGetCommentNum(sVid)
		if err != nil {
			return 0, err
		}
		return sum, nil
	}
}

func GetOtherCommentList(vid, parentId int64, page, size int) ([]*Comment, []*ReplyComment, error) {
	//先根据分页获取到所有二级评论
	offset := (page - 1) * size
	allComments, err := GetOtherCommentListByDB(vid, parentId, offset)
	if err != nil {
		return nil, nil, err
	}
	normalComments := make([]*Comment, 0)
	replyComments := make([]*ReplyComment, 0)
	//分别封装
	for _, comment := range allComments {
		//根据id去reply表中
		replyId, err := GetReplyId(comment.Id)
		if err != nil {
			return nil, nil, err
		}
		log.Debugf("get reply id =%d", replyId)
		if replyId == 0 {
			normalComments = append(normalComments, comment)
		} else {
			userId, err := GetUserId(replyId)
			log.Debugf("get reply user_id =%d", userId)
			if err != nil {
				return nil, nil, err
			}
			replyComment := &ReplyComment{
				Comment:     *comment,
				ReplyId:     replyId,
				ReplyUserId: userId,
			}
			replyComments = append(replyComments, replyComment)
			log.Debugf("get replyComment=%v", replyComment)
		}
	}
	return normalComments, replyComments, nil
}

func GetTopCommentList(vid int64, page, size int) ([]*Comment, error) {
	offset := (page - 1) * size
	comments, err := GetTopCommentListByDB(vid, offset)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func GetTopCommentReplyNum(commentId int64) (int64, error) {
	num, err := GetTopCommentReplyNumByDB(commentId)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func CreateTopComment(uid, vid int64, text string) (*Comment, error) {
	//先创建到数据库
	comment, err := CreateTopCommentByDB(uid, vid, text)
	if err != nil {
		return nil, err
	}
	sVid := strconv.FormatInt(vid, 10)
	err, flag := Exist(sVid)
	if err != nil {
		return nil, err
	}
	//如果不存在，则从数据库更新
	if !flag {
		ids, err := GetCommentIdsByVideoId(vid)
		if err != nil {
			return nil, err
		}
		for _, id := range ids {
			err := VideoCommentNumAddByCache(sVid, strconv.FormatInt(id, 10))
			if err != nil {
				err = DelFormCache(sVid)
				return nil, err
			}
		}
	} else {
		err = VideoCommentNumAddByCache(sVid, strconv.FormatInt(comment.Id, 10))
		if err != nil {
			err = DelFormCache(sVid)
			return nil, err
		}
	}
	return comment, nil
}

func DeleteTopComment(vid, commentId int64) error {
	err := DeleteTopCommentByDB(commentId)
	if err != nil {
		return err
	}
	sVid := strconv.FormatInt(vid, 10)
	err, flag := Exist(sVid)
	if err != nil {
		return err
	}
	//如果不存在，则从数据库更新,然后删除
	if !flag {
		ids, err := GetCommentIdsByVideoId(vid)
		if err != nil {
			return err
		}
		for _, id := range ids {
			err := VideoCommentNumAddByCache(sVid, strconv.FormatInt(id, 10))
			if err != nil {
				err = DelFormCache(sVid)
				return err
			}
		}
	} else {
		err = DelMemberFromCache(sVid, strconv.FormatInt(commentId, 10))
		if err != nil {
			err = DelFormCache(sVid)
			return err
		}
	}
	return nil
}

func CreateOtherComment(vid, uid, parentId, replyId, parentUid int64, text string) (*Comment, error) {
	//先创建数据到数据库
	comment, err := CreateOtherCommentByDB(vid, uid, parentId, replyId, parentUid, text)
	if err != nil {
		return nil, err
	}
	sVid := strconv.FormatInt(vid, 10)
	err, flag := Exist(sVid)
	if err != nil {
		return nil, err
	}
	//如果不存在，则从数据库更新
	if !flag {
		ids, err := GetCommentIdsByVideoId(vid)
		if err != nil {
			return nil, err
		}
		for _, id := range ids {
			err := VideoCommentNumAddByCache(sVid, strconv.FormatInt(id, 10))
			if err != nil {
				err = DelFormCache(sVid)
				return nil, err
			}
		}
	} else {
		err = VideoCommentNumAddByCache(sVid, strconv.FormatInt(comment.Id, 10))
		if err != nil {
			err = DelFormCache(sVid)
			return nil, err
		}
	}
	return comment, nil
}

func DeleteOtherComment(vid, commentId int64) error {
	err := DeleteOtherCommentByDB(commentId)
	if err != nil {
		return err
	}
	sVid := strconv.FormatInt(vid, 10)
	err, flag := Exist(sVid)
	if err != nil {
		return err
	}
	//如果不存在，则从数据库更新
	if !flag {
		ids, err := GetCommentIdsByVideoId(vid)
		if err != nil {
			return err
		}
		for _, id := range ids {
			err := VideoCommentNumAddByCache(sVid, strconv.FormatInt(id, 10))
			if err != nil {
				err = DelFormCache(sVid)
				return err
			}
		}
	} else {
		err = DelMemberFromCache(sVid, strconv.FormatInt(commentId, 10))
		if err != nil {
			err = DelFormCache(sVid)
			return err
		}
	}
	return nil
}
