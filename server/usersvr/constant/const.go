package constant

import "time"

const (
	SuccessCode = 0
	ErrorCode   = 1
	KeyExpire   = 5 * time.Minute //5分钟过期
)

const (
	SuccessMsg = "success"
)

const (
	ReqUuid           = "uuid"
	SessionKeyPrefix  = "session_"
	CommentInfoPrefix = "comment_info_"
	VideoInfoPrefix   = "video_info_"
)
