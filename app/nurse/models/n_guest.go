package models

import (
	"go-admin/common/models"
)

type NGuest struct {
	Id       int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	UserName string `json:"username" gorm:"size:128;comment:用户名"`
	Password string `json:"password" gorm:"size:128;comment:密码"`
	Salt     string `json:"salt" gorm:"size:128;comment:密码-混淆符"`
	NickName string `json:"nickname" gorm:"size:128;comment:昵称"`
	Phone    string `json:"phone" gorm:"size:16;comment:手机号"`
	models.ModelTime
	models.ControlBy
}

func (NGuest) TableName() string {
	return "n_guest"
}

func (e *NGuest) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *NGuest) GetId() interface{} {
	return e.Id
}
