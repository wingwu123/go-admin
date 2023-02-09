package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/nurse/models"
	"go-admin/app/nurse/service/dto"
	cDto "go-admin/common/dto"
)

type NGuest struct {
	service.Service
}

// GetPage 获取列表
func (e *NGuest) GetPage(c *dto.NGuestGetPageReq, list *[]models.NGuest, count *int64) error {
	var err error
	var data models.NGuest

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Get 获取对象
func (e *NGuest) Get(d *dto.NGuestGetReq, model *models.NGuest) error {
	var err error
	var data models.NGuest

	db := e.Orm.Model(&data).
		First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if err = db.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Insert 创建对象
func (e *NGuest) Insert(c *dto.NGuestInsertReq) (int, error) {
	var err error
	var data = new(models.NGuest)
	c.Generate(data)
	err = e.Orm.Create(data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return 0, err
	}
	return data.Id, nil
}

// Update 修改对象
func (e *NGuest) Update(c *dto.NGuestUpdateReq) error {
	var err error
	var model = models.NGuest{}
	e.Orm.First(&model, c.GetId())
	c.Generate(&model)
	db := e.Orm.Save(model)
	if err = db.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	return nil
}

// Remove 删除
func (e *NGuest) Remove(c *dto.NGuestDeleteReq) error {
	var err error
	var data models.NGuest

	db := e.Orm.Delete(&data, c.GetId())
	if err = db.Error; err != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}

// GetAll 获取所有
func (e *NGuest) GetAll(c *dto.NGuestGetPageReq, list *[]models.NGuest) error {
	var err error
	var data models.NGuest

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}
