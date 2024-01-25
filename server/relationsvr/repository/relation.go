package repository

import (
	"gorm.io/gorm"
	"relationsvr/log"
	"strconv"
)

func FollowerNum(uid int64) (int64, error) {
	err, flag := ExistFollowerKey(strconv.FormatInt(uid, 10))
	if err != nil {
		return 0, err
	}
	if !flag {
		//不存在去数据库查，但是数据库可能不存在，但是不是出现错误，只是这个人没有粉丝而已
		ids, err := GetFollowerListByDB(uid)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return 0, nil
			}
			return 0, err
		}
		//更新到缓存
		for _, id := range ids {
			err = FollowerNumAddByCache(strconv.FormatInt(uid, 10), strconv.FormatInt(id, 10))
			if err != nil {
				DelFollowerFromCache(strconv.FormatInt(uid, 10))
				return 0, err
			}
		}
		return int64(len(ids)), nil
	} else {
		sum, err := CacheGetFollowerNum(strconv.FormatInt(uid, 10))
		if err != nil {
			return sum, err
		}
		return sum, nil
	}
}

func FollowNum(uid int64) (int64, error) {
	err, flag := ExistFollowKey(strconv.FormatInt(uid, 10))
	if err != nil {
		return 0, err
	}
	if !flag {
		//不存在去数据库查，但是数据库可能不存在，但是不是出现错误，只是这个人没有关注者而已
		ids, err := GetFollowListByDB(uid)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return 0, nil
			}
			return 0, err
		}
		//更新到缓存
		for _, id := range ids {
			err = FollowAddByCache(strconv.FormatInt(uid, 10), strconv.FormatInt(id, 10))
			if err != nil {
				DelFollowFromCache(strconv.FormatInt(uid, 10))
				return 0, err
			}
		}
		return int64(len(ids)), nil
	} else {
		sum, err := CacheGetFollowNum(strconv.FormatInt(uid, 10))
		if err != nil {
			return 0, err
		}
		return sum, nil
	}

}

func RelationIsFollow(uid, target int64) (bool, error) {
	err, flag := ExistFollowKey(strconv.FormatInt(uid, 10))
	if err != nil {
		return false, err
	}
	if !flag {
		//不在就查数据库然后更新缓存
		ids, err := GetFollowListByDB(uid)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return false, nil
			}
			return false, err
		}
		//更新到缓存
		for _, id := range ids {
			err = FollowAddByCache(strconv.FormatInt(uid, 10), strconv.FormatInt(id, 10))
			if err != nil {
				DelFollowFromCache(strconv.FormatInt(uid, 10))
				return false, err
			}
		}
		check, err := IsFollowCheckByCache(strconv.FormatInt(uid, 10), strconv.FormatInt(target, 10))
		if err != nil {
			return false, err
		}
		return check, nil
	} else {
		check, err := IsFollowCheckByCache(strconv.FormatInt(uid, 10), strconv.FormatInt(target, 10))
		if err != nil {
			return false, err
		}
		return check, nil
	}
}

func RelationFollowList(uid int64) ([]int64, error) {
	//先查缓存，不在去数据库查询并更新
	err, flag := ExistFollowKey(strconv.FormatInt(uid, 10))
	if err != nil {
		return nil, err
	}
	if !flag {
		//不存在去数据库查，但是数据库可能不存在，但是不是出现错误，只是这个人没有关注者而已
		ids, err := GetFollowListByDB(uid)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, err
		}
		//更新到缓存
		for _, id := range ids {
			err = FollowAddByCache(strconv.FormatInt(uid, 10), strconv.FormatInt(id, 10))
			if err != nil {
				DelFollowFromCache(strconv.FormatInt(uid, 10))
				return nil, err
			}
		}
		return ids, nil
	} else {
		ids, err := GetFollowListFromCache(strconv.FormatInt(uid, 10))
		if err != nil {
			return nil, err
		}
		return ids, nil
	}
}

func RelationFollowerList(uid int64) ([]int64, error) {
	//先查缓存，不在去数据库查询并更新
	err, flag := ExistFollowerKey(strconv.FormatInt(uid, 10))
	if err != nil {
		return nil, err
	}
	if !flag {
		//不存在去数据库查，但是数据库可能不存在，但是不是出现错误，只是这个人没有粉丝而已
		ids, err := GetFollowerListByDB(uid)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, err
		}
		//更新到缓存
		for _, id := range ids {
			err = FollowerNumAddByCache(strconv.FormatInt(uid, 10), strconv.FormatInt(id, 10))
			if err != nil {
				DelFollowerFromCache(strconv.FormatInt(uid, 10))
				return nil, err
			}
		}
		return ids, nil
	} else {
		ids, err := GetFollowerListByCache(strconv.FormatInt(uid, 10))
		if err != nil {
			return nil, err
		}
		return ids, nil
	}
}

func RelationAdd(uid, target int64) error {
	//先将关系插入到数据库中，再更新到缓存中
	err := InsertRelationToDb(uid, target)
	if err != nil {
		return err
	}
	//去缓存中更新
	err = RelationAddToCache(uid, target)
	if err != nil {
		return err
	}
	return nil
}

func RelationAddToCache(uid, target int64) error {
	//先更新粉丝缓存
	err, flag := ExistFollowerKey(strconv.FormatInt(target, 10))
	if err != nil {
		return err
	}
	//如果不存在就去数据库拉去,更新到缓存中
	if !flag {
		ids, err := GetFollowerListByDB(target)
		if err != nil {
			return err
		}
		for _, id := range ids {
			err = FollowerNumAddByCache(strconv.FormatInt(target, 10), strconv.FormatInt(id, 10))
			if err != nil {
				DelFollowerFromCache(strconv.FormatInt(target, 10))
				return err
			}
		}
	} else {
		//如果缓存没过期，就直接更新
		err = FollowerNumAddByCache(strconv.FormatInt(target, 10), strconv.FormatInt(uid, 10))
		if err != nil {
			DelFollowerFromCache(strconv.FormatInt(target, 10))
			return err
		}
	}
	//更新关注者缓存
	err, flag = ExistFollowKey(strconv.FormatInt(uid, 10))
	if err != nil {
		return err
	}
	//如果不存在就去数据库拉去,更新到缓存中
	if !flag {
		ids, err := GetFollowListByDB(uid)
		if err != nil {
			return err
		}
		for _, id := range ids {
			err = FollowAddByCache(strconv.FormatInt(uid, 10), strconv.FormatInt(id, 10))
			if err != nil {
				DelFollowFromCache(strconv.FormatInt(uid, 10))
				return err
			}
		}
	} else {
		//如果缓存没过期，就直接更新
		err = FollowAddByCache(strconv.FormatInt(uid, 10), strconv.FormatInt(target, 10))
		if err != nil {
			DelFollowerFromCache(strconv.FormatInt(uid, 10))
			return err
		}
	}
	return nil
}

func RelationDel(uid, target int64) error {
	//先从缓存中删除，再去删除数据库
	//先更新粉丝缓存
	err, flag := ExistFollowerKey(strconv.FormatInt(target, 10))
	if err != nil {
		return err
	}
	//如果不存在就去数据库拉去,更新到缓存中
	if !flag {
		ids, err := GetFollowerListByDB(target)
		if err != nil {
			return err
		}
		for _, id := range ids {
			err = FollowerNumAddByCache(strconv.FormatInt(target, 10), strconv.FormatInt(id, 10))
			if err != nil {
				DelFollowerFromCache(strconv.FormatInt(target, 10))
				return err
			}
		}
		err = DelFollowerMemberFromCache(strconv.FormatInt(target, 10), strconv.FormatInt(uid, 10))
		if err != nil {
			DelFollowerFromCache(strconv.FormatInt(target, 10))
			return err
		}
	} else {
		//如果缓存没过期，就直接更新
		log.Info("DelFollowerMemberFromCache init")
		err = DelFollowerMemberFromCache(strconv.FormatInt(target, 10), strconv.FormatInt(uid, 10))
		if err != nil {
			DelFollowerFromCache(strconv.FormatInt(target, 10))
			return err
		}
	}
	//更新关注者缓存
	err, flag = ExistFollowKey(strconv.FormatInt(uid, 10))
	if err != nil {
		return err
	}
	//如果不存在就去数据库拉去,更新到缓存中
	if !flag {
		ids, err := GetFollowListByDB(uid)
		if err != nil {
			return err
		}
		for _, id := range ids {
			err = FollowAddByCache(strconv.FormatInt(uid, 10), strconv.FormatInt(id, 10))
			if err != nil {
				DelFollowFromCache(strconv.FormatInt(uid, 10))
				return err
			}
		}
		err = DelFollowMemberFromCache(strconv.FormatInt(uid, 10), strconv.FormatInt(target, 10))
		if err != nil {
			DelFollowFromCache(strconv.FormatInt(uid, 10))
			return err
		}
	} else {
		//如果缓存没过期，就直接更新
		log.Info(" DelFollowMemberFromCache init")
		err = DelFollowMemberFromCache(strconv.FormatInt(uid, 10), strconv.FormatInt(target, 10))
		if err != nil {
			DelFollowFromCache(strconv.FormatInt(uid, 10))
			return err
		}
	}
	err = DeleteRelationByDb(uid, target)
	if err != nil {
		return err
	}
	return nil
}
