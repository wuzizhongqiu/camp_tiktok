package respository

import (
	"gorm.io/gorm"
	"strconv"
)

func GetCommentLikeNum(cid int64) (int64, error) {
	flag, err := ExistCommentKey(strconv.FormatInt(cid, 10))
	if err != nil {
		return 0, err
	}
	var sum int64
	//如果不存在，去数据库查
	if !flag {
		ids, err := GetUserIdListByDB(cid)
		if err != nil {
			return 0, err
		}
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		for _, uid := range ids {
			err = CommentLikeNumAddByCache(strconv.FormatInt(cid, 10), strconv.FormatInt(uid, 10))
			if err != nil {
				err = DelCommentFromCache(strconv.FormatInt(cid, 10))
				return 0, err
			}
		}
		sum = int64(len(ids))
	} else {
		sum, err = CacheGetCommentLikeNum(strconv.FormatInt(cid, 10))
		if err != nil {
			err = DelCommentFromCache(strconv.FormatInt(cid, 10))
			return 0, err
		}
	}
	return sum, nil
}

func IsFavoriteComment(uid, cid int64) (bool, error) {
	flag, err := ExistCommentKey(strconv.FormatInt(cid, 10))
	if err != nil {
		return false, err
	}
	//不存在更新缓存
	if !flag {
		ids, err := GetUserIdListByDB(cid)
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		if err != nil {
			return false, err
		}
		for _, id := range ids {
			err = CommentLikeNumAddByCache(strconv.FormatInt(cid, 10), strconv.FormatInt(id, 10))
			if err != nil {
				err = DelCommentFromCache(strconv.FormatInt(uid, 10))
				return false, err
			}
		}
		//查询数据库
		exist, err := IsUserLikCommentCheckByDB(cid, uid)
		if err != nil {
			return false, err
		}
		return exist, nil
	} else {
		exist, err := IsUserLikeCommentCheckByCache(strconv.FormatInt(uid, 10), strconv.FormatInt(cid, 10))
		if err != nil {
			return false, err
		}
		return exist, nil
	}
}

func CommentUnLikeAction(cid, uid int64) error {
	err := DeleteCommentLike(cid, uid)
	if err != nil {
		return err
	}
	err = CommentUnLikeCache(cid, uid)
	if err != nil {
		return err
	}
	return nil
}

func CommentUnLikeCache(cid, uid int64) error {
	flag, err := ExistCommentKey(strconv.FormatInt(cid, 10))
	if err != nil {
		return err
	}
	//如果不存在
	if !flag {
		//更新到缓存中
		ids, err := GetUserIdListByDB(cid)
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		if err != nil {
			return err
		}
		for _, id := range ids {
			err = CommentLikeNumAddByCache(strconv.FormatInt(cid, 10), strconv.FormatInt(id, 10))
			if err != nil {
				err = DelCommentFromCache(strconv.FormatInt(cid, 10))
				return err
			}
		}
	} else {
		err = DelCommentMemberFromCache(strconv.FormatInt(cid, 10), strconv.FormatInt(uid, 10))
		if err != nil {
			err = DelCommentFromCache(strconv.FormatInt(cid, 10))
			return err
		}
	}
	return nil
}

func CommentLikeAction(cid, uid int64) error {
	err := InsertCommentLike(cid, uid)
	if err != nil {
		return err
	}
	err = CommentLikeCache(cid, uid)
	if err != nil {
		return err
	}
	return nil
}

func CommentLikeCache(cid, uid int64) error {
	flag, err := ExistCommentKey(strconv.FormatInt(cid, 10))
	if err != nil {
		return err
	}
	//如果不存在
	if !flag {
		//更新到缓存中
		ids, err := GetUserIdListByDB(cid)
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		if err != nil {
			return err
		}
		for _, id := range ids {
			err = CommentLikeNumAddByCache(strconv.FormatInt(cid, 10), strconv.FormatInt(id, 10))
			if err != nil {
				err = DelCommentFromCache(strconv.FormatInt(cid, 10))
				return err
			}
		}
	} else {
		err = CommentLikeNumAddByCache(strconv.FormatInt(cid, 10), strconv.FormatInt(uid, 10))
		if err != nil {
			err = DelCommentFromCache(strconv.FormatInt(cid, 10))
			return err
		}
	}
	return nil
}

func IsFavoriteVideo(uid, vid int64) (bool, error) {
	err, flag := ExistUserKey(strconv.FormatInt(uid, 10))
	if err != nil {
		return false, err
	}
	//不存在更新缓存
	if !flag {
		ids, err := GetUserLikeList(uid)
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		if err != nil {
			return false, err
		}
		for _, id := range ids {
			err = UserLikeAddByCache(strconv.FormatInt(id, 10), strconv.FormatInt(uid, 10))
			if err != nil {
				err = DelUserFromCache(strconv.FormatInt(uid, 10))
				return false, err
			}
		}
		//查询数据库
		exist, err := IsUserLikeVideoCheckByDB(vid, uid)
		if err != nil {
			return false, err
		}
		return exist, nil
	} else {
		exist, err := IsUserLikeVideoCheckByCache(strconv.FormatInt(uid, 10), strconv.FormatInt(vid, 10))
		if err != nil {
			return false, err
		}
		return exist, nil
	}
}

func GetUserLikeVideos(uid int64) ([]int64, error) {
	err, flag := ExistUserKey(strconv.FormatInt(uid, 10))
	if err != nil {
		return nil, err
	}
	if !flag {
		ids, err := GetUserLikeList(uid)
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
		for _, id := range ids {
			err = UserLikeAddByCache(strconv.FormatInt(id, 10), strconv.FormatInt(uid, 10))
			if err != nil {
				err = DelUserFromCache(strconv.FormatInt(uid, 10))
				return nil, err
			}
		}
		return ids, nil
	} else {
		ids, err := GetUserLikeListFromCache(strconv.FormatInt(uid, 10))
		if err != nil {
			err = DelUserFromCache(strconv.FormatInt(uid, 10))
			return nil, err
		}
		return ids, nil
	}
}

func VideoLikeNum(vid int64) (int64, error) {
	sum, err := CacheGetVLikeNum(vid)
	if err != nil {
		return 0, err
	}
	return sum, nil
}

func CacheGetVLikeNum(vid int64) (int64, error) {
	err, flag := Exist(strconv.FormatInt(vid, 10))
	if err != nil {
		return 0, err
	}
	var sum int64
	//如果不存在，去数据库查
	if !flag {
		ids, err := GetVideoLikeList(vid)
		if err != nil {
			return 0, err
		}
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		for _, uid := range ids {
			err = VideoLikeNumAddByCache(strconv.FormatInt(vid, 10), strconv.FormatInt(uid, 10))
			if err != nil {
				err = DelFromCache(strconv.FormatInt(vid, 10))
				return 0, err
			}
		}
		sum = int64(len(ids))
	} else {
		sum, err = CacheGetVideoLikeNum(strconv.FormatInt(vid, 10))
		if err != nil {
			err = DelFromCache(strconv.FormatInt(vid, 10))
			return 0, err
		}
	}
	return sum, nil
}

func VideoLikeAction(vid, uid int64) error {
	err := InsertVideoLike(vid, uid)
	if err != nil {
		return err
	}
	//更新到视频缓存
	err = VideoLikeUpdateToCache(vid, uid)
	if err != nil {
		return err
	}
	//更新用户喜爱缓存
	err = UserLikeVideoToCache(vid, uid)
	if err != nil {
		return err
	}
	return nil
}

func VideoLikeUpdateToCache(vid, uid int64) error {
	err, flag := Exist(strconv.FormatInt(vid, 10))
	if err != nil {
		return err
	}
	//如果不存在，去数据库查
	if !flag {
		ids, err := GetVideoLikeList(vid)
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		if err != nil {
			return err
		}
		for _, uid := range ids {
			err = VideoLikeNumAddByCache(strconv.FormatInt(vid, 10), strconv.FormatInt(uid, 10))
			if err != nil {
				err = DelFromCache(strconv.FormatInt(vid, 10))
				return err
			}
		}
	} else {
		err = VideoLikeNumAddByCache(strconv.FormatInt(vid, 10), strconv.FormatInt(uid, 10))
		if err != nil {
			err = DelFromCache(strconv.FormatInt(vid, 10))
			return err
		}
	}
	return nil
}

func VideoUnLikeAction(vid, uid int64) error {
	//先更新到数据库
	err := DeleteVideoLike(vid, uid)
	if err != nil {
		return err
	}
	err = VideoUnLikeUpdateToCache(vid, uid)
	if err != nil {
		return err
	}
	err = UserUnLikeVideoToCache(vid, uid)
	if err != nil {
		return err
	}
	return nil
}

func VideoUnLikeUpdateToCache(vid, uid int64) error {
	//先查询是否存在
	err, flag := Exist(strconv.FormatInt(vid, 10))
	if err != nil {
		return err
	}
	//如果不存在，去数据库查
	if !flag {
		ids, err := GetVideoLikeList(vid)
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		if err != nil {
			return err
		}
		for _, id := range ids {
			//加入到缓存中,此时不用删除因为在数据库中已经删除了
			err = VideoLikeNumAddByCache(strconv.FormatInt(vid, 10), strconv.FormatInt(id, 10))
			if err != nil {
				err = DelFromCache(strconv.FormatInt(vid, 10))
				return err
			}
		}
	} else {
		//如果缓存没过期，就从缓存中删除
		err = DelMemberFromCache(strconv.FormatInt(vid, 10), strconv.FormatInt(uid, 10))
		if err != nil {
			err = DelFromCache(strconv.FormatInt(vid, 10))
			return err
		}
	}
	return nil
}

func UserLikeVideoToCache(vid, uid int64) error {
	err, flag := ExistUserKey(strconv.FormatInt(uid, 10))
	if err != nil {
		return err
	}
	//如果不存在，就去数据库查询,然后更新缓存
	if !flag {
		ids, err := GetUserLikeList(uid)
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		if err != nil {
			return err
		}
		for _, id := range ids {
			err = UserLikeAddByCache(strconv.FormatInt(id, 10), strconv.FormatInt(uid, 10))
			if err != nil {
				err = DelUserFromCache(strconv.FormatInt(uid, 10))
				return err
			}
		}
	} else {
		err = UserLikeAddByCache(strconv.FormatInt(vid, 10), strconv.FormatInt(uid, 10))
		if err != nil {
			err = DelUserFromCache(strconv.FormatInt(uid, 10))
			return err
		}
	}
	return nil
}

func UserUnLikeVideoToCache(vid, uid int64) error {
	err, flag := ExistUserKey(strconv.FormatInt(uid, 10))
	if err != nil {
		return err
	}
	//如果不存在，就去数据库查询,然后更新缓存
	if !flag {
		ids, err := GetUserLikeList(uid)
		//如果数据库里也没有，不代表有错
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		if err != nil {
			return err
		}
		for _, id := range ids {
			//这里不删除的原因是，先操作了数据库，再更新的缓存，此时记录已经在数据库查询不到了
			err = UserLikeAddByCache(strconv.FormatInt(id, 10), strconv.FormatInt(uid, 10))
			if err != nil {
				err = DelUserFromCache(strconv.FormatInt(uid, 10))
				return err
			}
		}
	} else {
		err = DelUserMemberFromCache(strconv.FormatInt(vid, 10), strconv.FormatInt(uid, 10))
		if err != nil {
			err = DelUserFromCache(strconv.FormatInt(uid, 10))
			return err
		}
	}
	return nil
}
