package repository

var followerKeyPrefix = "tiktok:follower:"
var followKeyPrefix = "tiktok:follow"

type Relation struct {
	Id         int64
	UserId     int64 //用户id
	FollowerId int64 //关注者id
	Delete     int   //是否取消关注关系
}

func (r Relation) TableName() string {
	return "t_relation"
}
