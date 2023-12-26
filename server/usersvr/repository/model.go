package repository

// redis user key prefix
var userKeyPrefix = "tiktok:user:"

type User struct {
	//gorm.Model
	Id              int64
	UserName        string
	Password        string
	FollowCount     int64
	FollowerCount   int64
	Avatar          string
	BackgroundImage string
	Signature       string
	TotalFavorited  int64 //被喜欢的视频数量
	FavoriteCount   int64 //该用户点赞的视频数量
}

func (r User) TableName() string {
	return "t_user"
}
