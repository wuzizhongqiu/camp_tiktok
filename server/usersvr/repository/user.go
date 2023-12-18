package repository

import (
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"usersvr/log"
)

var (
	Secret = []byte("tiktok")
)

// JWTClaims 自定义claim
type JWTClaims struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.Claims
}

func UpdateUserFollowerNum(uid int64, actionType int64) error {
	var num int64
	if actionType == 1 {
		num = 1
	} else {
		num = -1
	}
	//先更新数据库
	var err error
	err = UpdateFollowerNumByDB(uid, num)
	if err != nil {
		return err
	}
	//查缓存是否存在
	var flag bool
	err, flag = CacheCheckUser(uid)
	if err != nil {
		return err
	}
	//如果不存在
	if !flag {
		user, err := GetUserInfoFromDB(uid)
		if err != nil {
			return err
		}
		//异步去更新缓存
		go CacheSetUserInfo(user)
		return nil
	}
	//如果存在
	err = CacheUpdateFollowerNum(uid, num)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserFollowNum(uid int64, actionType int64) error {
	var num int64
	if actionType == 1 {
		num = 1
	} else {
		num = -1
	}
	//先更新数据库
	var err error
	err = UpdateFollowNumByDB(uid, num)
	if err != nil {
		return err
	}
	//查缓存是否存在
	var flag bool
	err, flag = CacheCheckUser(uid)
	if err != nil {
		return err
	}
	//如果不存在
	if !flag {
		user, err := GetUserInfoFromDB(uid)
		if err != nil {
			return err
		}
		//异步去更新缓存
		go CacheSetUserInfo(user)
		return nil
	}
	//如果存在
	err = CacheUpdateFollowNum(uid, num)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserFavoriteNum(uid int64, actionType int64) error {
	var num int64
	if actionType == 1 {
		num = 1
	} else {
		num = -1
	}
	//先更新数据库
	var err error
	err = UpdateFavCountNumByDB(uid, num)
	if err != nil {
		return err
	}
	//查缓存是否存在
	var flag bool
	err, flag = CacheCheckUser(uid)
	if err != nil {
		return err
	}
	//如果不存在
	if !flag {
		user, err := GetUserInfoFromDB(uid)
		if err != nil {
			return err
		}
		//异步去更新缓存
		go CacheSetUserInfo(user)
		return nil
	}
	//如果存在
	err = CacheUpdateFavoriteNum(uid, num)
	if err != nil {
		return err
	}
	return nil

}

func UpdateUserFavoritedNum(uid int64, actionType int64) error {
	var num int64
	if actionType == 1 {
		num = 1
	} else {
		num = -1
	}

	var err error
	err = UpdateTotalFav(uid, num)
	if err != nil {
		return err
	}

	//能更新成功说明在数据库中存在
	var flag bool
	err, flag = CacheCheckUser(uid)
	if err != nil {
		return err
	}

	//说明过期了,数据次数在数据库中更新过了
	if !flag {
		user, err := GetUserInfoFromDB(uid)
		if err != nil {
			return err
		}
		go CacheSetUserInfo(user)
		return nil
	}

	//数据没过期,直接去缓存中更新
	err = CacheUpdateFavoritedNum(uid, num)
	if err != nil {
		return err
	}
	return nil
}

func GenerateToken(uid int64, username string) (string, error) {
	//生成claim
	claims := &JWTClaims{
		UserId:   uid,
		UserName: username,
		Claims: jwt.RegisteredClaims{
			Issuer: "server",
		},
	}
	//生成token
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := withClaims.SignedString(Secret)
	if err != nil {
		log.Errorf("GenerateToken err=%v", err)
		return "", err
	}
	return token, nil
}

func GetUserInfo(u interface{}) (user User, err error) {
	switch u.(type) {
	case int64:
		//先从缓存查询
		user, err = CacheGetUserInfo(u.(int64))
		if err == nil {
			return user, nil
		}
		//去db查询
		user, err = GetUserInfoFromDB(u.(int64))
	case string:
		user, err = GetUserInfoByUserName(u.(string))
	}
	return
}

func GetUserInfoList(ids []int64) ([]*User, error) {
	users, err := GetUserListByDB(ids)
	if err != nil {
		return nil, err
	}
	return users, err
}

func InsertUser(username string, password string) (User, error) {
	user, err := UserInsert(username, password)
	if err != nil {
		return User{}, err
	}
	go CacheSetUserInfo(user)
	return user, nil
}

func PassWordHash(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("bcrypt.GenerateFromPassword err==%v", err)
		return "", err
	}
	return string(pass), nil
}
