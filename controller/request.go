package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

var (
	ErrorUserNotLogin = errors.New("用户未登录")
)

// 获取当前登录用户的id
func getCurrentUser(ctx *gin.Context) (userID int64, err error) {
	uid, ok := ctx.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}

	return
}

// getPageInfo 获取分页参数
func getPageInfo(ctx *gin.Context) (int64, int64) {
	pageStr := ctx.Query("page")
	sizeStr := ctx.Query("size")

	var (
		page int64
		size int64
		err  error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil || page < 1 {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil || size < 1 {
		size = 10
	} else if size > 100 {
		size = 100
	}

	return page, size
}
