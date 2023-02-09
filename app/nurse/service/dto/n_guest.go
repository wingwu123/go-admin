package dto

import (
	"go-admin/app/nurse/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type NGuestGetPageReq struct {
	dto.Pagination `search:"-"`
	Id             int    `form:"id" search:"type:exact;column:id;table:n_guest" comment:""`
	UserName       string `form:"user_name" search:"type:contains;column:username;table:n_guest" comment:""`
	NickName       string `form:"nick_name" search:"type:contains;column:nickname;table:n_guest" comment:""`
	Phone          string `form:"phone" search:"type:contains;column:phone;table:n_guest" comment:""`
}

func (m *NGuestGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type NGuestInsertReq struct {
	UserName string `json:"user_name" comment:""`
	Password string `json:"password" comment:""`
	Phone    string `json:"phone" comment:""`
	common.ControlBy
}

func (s *NGuestInsertReq) Generate(model *models.NGuest) {
	model.UserName = s.UserName
	model.Password = s.Password
	model.Phone = s.Phone
}

type NGuestUpdateReq struct {
	Id    int    `uri:"Id" comment:""`
	Phone string `json:"phone" comment:""`
	common.ControlBy
}

func (s *NGuestUpdateReq) Generate(model *models.NGuest) {
	model.Id = s.Id
	model.Phone = s.Phone
}

func (s *NGuestUpdateReq) GetId() interface{} {
	return s.Id
}

type NGuestGetReq struct {
	Id int `uri:"dictCode"`
}

func (s *NGuestGetReq) GetId() interface{} {
	return s.Id
}

type NGuestDeleteReq struct {
	Ids              []int `json:"ids"`
	common.ControlBy `json:"-"`
}

func (s *NGuestDeleteReq) GetId() interface{} {
	return s.Ids
}
