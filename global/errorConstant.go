package global

import (
	"ginscaffold/pkg/result"
)
var (
	// OK
	OK = result.NewError(0, "OK")

	//参数模块
	ErrParam = result.NewError(400, "参数不合法")

	//文章模块报错
	ErrArticleNot = result.NewError(10001, "文章不存在")
	ErrArticleS = result.NewError(10002, "文章查询出错")

	//用户模块
	ErrUserNot = result.NewError(20001, "用户不存在")
	ErrUserNoHeader = result.NewError(20002, "缺少请求头")
	ErrUserNoAuth = result.NewError(20003, "请求头中auth为空")
	ErrUserToken = result.NewError(20004, "无效的Token")
	// ...
)
