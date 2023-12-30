package respository

import "time"

type CommentOther struct {
	Comment        Comment
	CommentReplies []CommentReply
}

var videoKeyPrefix = "tiktok:video:comment"

type Comment struct {
	Id           int64     //评论id,自增主键
	UserId       int64     //评论发布用户id
	VideoId      int64     //评论视频的id
	ParentId     int64     //父评论的id 一级评论的父id为0
	ParentUserId int64     //父评论的用户id，可以通过parent_id查询到，为了在查询时少一个sql就加了一个冗余字段
	CommentLevel int       //评论的等级 1为一级，2为多级
	Delete       int       //是否删除 0 不删除 1删除
	CreateTime   time.Time //评论发布的时间
	FavoriteNum  int64     //喜爱数
	CommentText  string    //评论的内容
}

type CommentReply struct {
	Id        int64 //主键
	CommentId int64 //评论的id与Comment的id
	ReplyId   int64 //回复评论的id
}

func (c Comment) TableName() string {
	return "t_comment"
}

func (r CommentReply) TableName() string {
	return "t_comment_replies"
}
