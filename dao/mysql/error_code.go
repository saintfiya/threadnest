package mysql

import "errors"

var (
	ErrorUserExist         = errors.New("用户已存在")
	ErrorUserNotExist      = errors.New("用户不存在")
	ErrorInvalidPassword   = errors.New("用户名或密码错误")
	ErrorPostNotExist      = errors.New("帖子不存在")
	ErrorCommunityNotExist = errors.New("社区不存在")
)
