package logic

import (
	"errors"
	"reddit/dao/mysql"
	"reddit/models"
	"reddit/pkg/gen"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户存不存在
	exists, err := mysql.CheckUserExists(p.UserName)
	if err != nil {
		return err // 数据库查询出错
	}
	if exists {
		return errors.New("user already exists")
	}
	// 生成 uid
	userID, err := gen.GenID()
	if err != nil {
		return errors.New("id generation failed")
	}
	// 构造一个 user 实例
	u := models.User{
		UserID:   int8(userID),
		UserName: p.UserName,
		Password: p.Password,
	}

	// 保存进 数据库
	return mysql.InsertUser(&u)
}

func Login(p *models.ParamLogin) (err error) {
	user := &models.User{
		UserName: p.UserName,
		Password: p.Password,
	}

	if err := mysql.Login(user); err != nil {
		return errors.New("login failed: " + err.Error())
	}

	return nil
}
