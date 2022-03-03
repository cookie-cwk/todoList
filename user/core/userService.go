package core

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"user/model"
	srv "user/services"
)

func (*UserService) UserLogin(ctx context.Context, req *srv.UserRequest, resp *srv.UserDetailResponse) error {
	var user model.User
	resp.Code = 200
	if err := model.DB.Where("user_name=?", req.UserName).First(&user).Error; err != nil {
		if gorm.ErrRecordNotFound != nil {
			resp.Code = 400
			return nil
		} else {
			resp.Code = 500
			return nil
		}
	}
	log.Println("user service")
	if user.CheckPassword(req.Password) == false {
		resp.Code = 400
		return nil
	}
	resp.UserDetail = BuildUser(user)
	return nil
}

func BuildUser(item model.User) *srv.UserModel {
	userModel := srv.UserModel{
		ID: uint32(item.ID),
		UserName: item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}

func (*UserService) UserRegister(ctx context.Context, req *srv.UserRequest, resp *srv.UserDetailResponse) error {
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码输入不一致")
		return err
	}
	var count int64 = 0
	if err := model.DB.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		err := errors.New("用户名已存在")
		return err
	}
	user := model.User{
		UserName: req.UserName,
	}
	if err := user.SetPassword(req.Password); err != nil {
		return err
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return err
	}
	resp.UserDetail = BuildUser(user)
	return nil
}