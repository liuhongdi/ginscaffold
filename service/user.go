package service

import (
"ginscaffold/dao"
"ginscaffold/model"
)
//得到一篇文章的详情
func GetOneUser(userName string) (*model.User, error) {
	return dao.SelectOneUser(userName)
}
