package repository

import (
	"gorm.io/gorm"
	"usersvr/log"
	"usersvr/middlerware/db"
)

func UpdateFollowerNumByDB(uid int64, num int64) error {
	Db := db.GetDb()
	err := Db.Model(&User{}).Where("id = ?", uid).Exec("follower_count+?", num).Error
	if err != nil {
		log.Errorf("UpdateFollowerNumByDB err===%v", err)
		return err
	}
	return nil
}

func UpdateFollowNumByDB(uid int64, num int64) error {
	Db := db.GetDb()
	err := Db.Model(&User{}).Where("id = ?", uid).Exec("follow_count+?", num).Error
	if err != nil {
		log.Errorf("UpdateFollowNumByDB err===%v", err)
		return err
	}
	return nil
}

func UpdateFavCountNumByDB(uid int64, num int64) error {
	Db := db.GetDb()
	err := Db.Model(&User{}).Where("id = ?", uid).Exec("favorite_count+?", num).Error
	if err != nil {
		log.Errorf("UpdateFavCountNumByDB err===%v", err)
		return err
	}
	return nil
}

func UpdateTotalFav(uid int64, num int64) error {
	Db := db.GetDb()
	err := Db.Model(&User{}).Where("id = ?", uid).Exec("total_favorited+?", num).Error
	if err != nil {
		log.Errorf("UpdateTotalFav err===%v", err)
		return err
	}
	return nil
}

func UserInsert(username string, password string) (User, error) {
	Db := db.GetDb()
	user := User{
		UserName:        username,
		Password:        password,
		FollowCount:     0,
		FollowerCount:   0,
		TotalFavorited:  0,
		FavoriteCount:   0,
		Avatar:          "https://tse1-mm.cn.bing.net/th/id/R-C.d83ded12079fa9e407e9928b8f300802?rik=Gzu6EnSylX9f1Q&riu=http%3a%2f%2fwww.webcarpenter.com%2fpictures%2fGo-gopher-programming-language.jpg&ehk=giVQvdvQiENrabreHFM8x%2fyOU70l%2fy6FOa6RS3viJ24%3d&risl=&pid=ImgRaw&r=0",
		BackgroundImage: "https://tse2-mm.cn.bing.net/th/id/OIP-C.sDoybxmH4DIpvO33-wQEPgHaEq?pid=ImgDet&rs=1",
		Signature:       "test sign",
	}
	if err := Db.Model(&User{}).Create(&user).Error; err != nil {
		log.Errorf("UserInsert err===%v", err)
		return User{}, err
	}
	return user, nil
}

func UserIsExist(username string) (bool, error) {
	Db := db.GetDb()
	var user User
	if err := Db.Where("user_name=?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		log.Errorf("UserIsExist err== %v", err)
		return false, err
	}
	return true, nil
}

func GetUserInfoFromDB(uid int64) (User, error) {
	Db := db.GetDb()
	var user User
	err := Db.Where("id=?", uid).First(&user).Error
	if err != nil {
		log.Errorf("GetUserInfoFromDB err=%v", err)
		return User{}, err
	}
	return user, nil
}

func GetUserInfoByUserName(username string) (User, error) {
	Db := db.GetDb()
	var user User
	err := Db.Where("user_name=?", username).First(&user).Error
	if err != nil {
		log.Errorf(" GetUserInfoByUserName err=%v", err)
		return User{}, err
	}
	return user, nil
}

func GetUserListByDB(ids []int64) ([]*User, error) {
	Db := db.GetDb()
	var users []*User
	if err := Db.Find(&users, ids).Error; err != nil {
		log.Errorf(" GetUserListByDB err=%v", err)
		return nil, err
	}
	return users, nil
}
