package repository

// redis user key prefix
var userKeyPrefix = "tiktok:user:"

type User struct {
	//gorm.Model
	Id              int64  `gorm:"column:id,primary_key;auto_increment"`
	Name            string `gorm:"column:user_name"`
	Password        string `gorm:"column:password"`
	Follow          int64  `gorm:"column:follow_count"`
	Follower        int64  `gorm:"column:follower_count"`
	Avatar          string `gorm:"column:avatar"`
	BackgroundImage string `gorm:"column:background_image"`
	Signature       string `gorm:"column:signature"`
	TotalFav        int64  `gorm:"column:total_favorited"` //被喜欢的视频数量
	FavCount        int64  `gorm:"column:favorite_count"`  //该用户点赞的视频数量
}

func (r User) TableName() string {
	return "t_user"
}