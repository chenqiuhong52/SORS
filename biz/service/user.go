package service

import (
	"SORS/biz/model"
	"SORS/dal/db"
	"SORS/pkg/util"
)

func Login(username string, password string) (*model.LoginResult, error) {
	// 查询用户是否存在，密码是否正确
	u, err := db.SelectUserByUsernameAndPassword(username, password)
	if err != nil {
		return nil, err
	}
	// 把需要的信息，放入Token （UserID）
	token, err := util.SignToken(u.ID)
	if err != nil {
		return nil, err
	}
	// 返回结果
	return &model.LoginResult{Token: token}, nil
}

func GetUserInfo(userID uint) (*model.GetUserInfoResult, error) {
	// 查询用户数据
	u, err := db.SelectUserByID(userID)
	if err != nil {
		return nil, err
	}
	// 查询用户权限
	level, err := db.SelectRoleLevelByID(u.RoleID)
	if err != nil {
		return nil, err
	}
	// 返回
	return &model.GetUserInfoResult{
		Username: u.Username,
		Role:     level,
	}, nil
}
