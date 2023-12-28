package db

import "errors"

type User struct {
	ID       uint
	Username string
	Password string
	RoleID   uint
}

func (User) TableName() string {
	return "user"
}

type Role struct {
	ID    uint
	Level string
}

func (Role) TableName() string {
	return "role"
}

func SelectUserByUsernameAndPassword(username string, password string) (*User, error) {
	res := &User{}
	dbRes := DB.Where("`username`  = ? AND `password` = ?", username, password).Find(res)
	if dbRes.RowsAffected != 1 {
		return nil, errors.New("用户名或者密码错误")
	}
	if dbRes.Error != nil {
		return nil, dbRes.Error
	}
	return res, nil
}

func SelectUserByID(userID uint) (*User, error) {
	res := &User{}
	err := DB.Where("`id` = ?", userID).Find(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func SelectRoleLevelByID(roleID uint) (string, error) {
	var res string
	err := DB.Select("`level`").Where("`id` = ?", roleID).Model(&Role{}).Find(&res).Error
	if err != nil {
		return "", err
	}
	return res, nil
}
