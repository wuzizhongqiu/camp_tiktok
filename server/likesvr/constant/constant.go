package constant

import "time"

const (
	KeyExist  = 1
	Delete    = 1
	NotDelete = 0
	Like      = 1
	UnLike    = 0
	KeyExpire = 3 * time.Minute //设计3分钟过期时间
)
