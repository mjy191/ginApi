package Service

import (
	"encoding/json"
	"errors"
	"ginApi/Common/Enum"
	"ginApi/Common/Tools"
	"ginApi/Models"
)

type UserService struct {
}

type UserCopy struct {
	Models.User
}

type customerUser struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type IdParam struct {
	Id int `form:"id" json:"id" binding:"required" msg:"id必填"`
}

type UserNameParam struct {
	UserName string `form:"username" json:"username" binding:"required,min=1" msg:"用户名不能为空"`
}

type PasswordParam struct {
	Password string `form:"password" json:"password" binding:"required,min=1" msg:"密码不能为空"`
}

type AddParam struct {
	Age  int    `form:"age" json:"age"`
	Name string `form:"name" json:"name" binding:"required"`
	Page int    `form:"page" json:"page"`
	UserNameParam
	PasswordParam
	Phone string `form:"phone" json:"phone" binding:"required" msg:"手机不能为空"`
}

type EditParam struct {
	AddParam
	IdParam
}

type DelParam struct {
	IdParam
}

type LoginParam struct {
	UserNameParam
	PasswordParam
}

func (u *UserCopy) MarshalJSON() ([]byte, error) {
	user := customerUser{
		Id:   u.Id,
		Name: u.Name,
	}
	return json.Marshal(user)
}

func (this UserService) Lists(userParam *AddParam) ([]*UserCopy, error) {
	var users []*Models.User
	err := Models.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	list := make([]*UserCopy, 0)
	for _, user := range users {
		list = append(list, &UserCopy{*user})
	}
	return list, nil
}

func (this UserService) Add(param *AddParam) int {
	var user Models.User
	var userCopy Models.User
	var phone Models.Phone
	user.Name = param.Name
	user.Age = param.Age
	user.UserName = param.UserName
	user.Password = Tools.Sha1(param.Password)

	phone.Phone = param.Phone

	//事务提交
	tx := Models.DB.Begin()
	tx.Model(&userCopy).Where("username=?", user.UserName).First(&userCopy)
	if userCopy.Id != 0 {
		tx.Rollback()
		panic(map[string]interface{}{
			"code": Enum.CodeParamError,
			"msg":  "用户名已经注册",
		})
	}
	tx.Create(&user)
	phone.UserId = user.Id
	tx.Create(&phone)
	tx.Commit()
	return user.Id
}

func (this UserService) Edit(param *EditParam) {
	var user Models.User
	var phone Models.Phone
	user.Name = param.Name
	user.UserName = param.UserName
	user.Password = Tools.Sha1(param.Password)
	user.Age = param.Age
	phone.Phone = param.Phone
	// 开启事务
	tx := Models.DB.Begin()
	tx.Model(&user).Where("id=?", param.Id).Updates(&user)
	tx.Model(&phone).Where("userId=?", param.Id).Updates(&phone)
	tx.Commit()
}

func (this UserService) Del(param *DelParam) {
	tx := Models.DB.Begin()
	tx.Delete(&Models.User{}, param.Id)
	tx.Where("userId=?", param.Id).Delete(&Models.Phone{})
	tx.Commit()
}

func (this UserService) Login(param *LoginParam) (Models.User, error) {
	var user Models.User
	Models.DB.Model(&user).Where("username=?", param.UserName).First(&user)
	if Tools.Sha1(param.Password) != user.Password {
		panic(map[string]interface{}{
			"code": Enum.CodeParamError,
			"msg":  "密码错误",
		})
		return user, errors.New("密码错误")
	}
	return user, nil
}
