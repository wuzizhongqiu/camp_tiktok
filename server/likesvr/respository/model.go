package respository

var videoKeyPrefix = "tiktok:video:like:"
var userKeyPrefix = "tiktok:user:like:"
var commentKeyPrefix = "tiktok:comment:like:"

type VideoLike struct {
	Id      int64 //主键
	VideoId int64 //视频id
	UserId  int64 //点赞的用户id
	Delete  int   //是否删除 0没删除 1删除了
}

type CommentLike struct {
	Id        int64 //主键
	CommentId int64 //点赞的id
	UserId    int64 //点赞的用户id
	Delete    int   //是否删除 0没删除 1删除了
}

func (v VideoLike) TableName() string {
	return "t_video_like"
}
func (v CommentLike) TableName() string {
	return "t_comment_like"
}
