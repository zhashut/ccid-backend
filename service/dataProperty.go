package service

import (
	"errors"
	"fmt"
	"go-init/global"
	"go-init/models"
	"go-init/respository/db"
	"time"
)

/**
 * @author: zhashut
 * Date: 2025/4/18
 * Time: 23:42
 * Description: 知识产权信息-服务层
 */

type DataPropertyService struct {
	DB *db.DataPropertyDao
}

func NewDataPropertyService() *DataPropertyService {
	return &DataPropertyService{DB: &db.DataPropertyDao{}}
}

// Create 创建知识产权
func (s *DataPropertyService) Create(req *models.DataPropertyCreateReq) (*models.DataProperty, error) {
	if req.Title == "" {
		return nil, errors.New("标题不能为空")
	}

	dataProperty := models.DataProperty{
		Type:   req.Type,
		Title:  req.Title,
		Source: req.Source,
		Time:   time.Now(),
		Status: req.Status,
	}

	result, err := s.DB.SaveChange(&dataProperty)
	if err != nil {
		return nil, err
	}
	return result.(*models.DataProperty), nil
}

// DeleteByID 删除知识产权
func (s *DataPropertyService) DeleteByID(id int64) error {
	if id <= 0 {
		return errors.New("无效的ID")
	}
	return s.DB.DeleteOneByWhere(&models.DataProperty{}, fmt.Sprintf("id = %v", id))
}

// Update 更新知识产权
func (s *DataPropertyService) Update(req *models.DataPropertyUpdateReq) (*models.DataProperty, error) {
	if req.ID == 0 {
		return nil, errors.New("更新需要指定ID")
	}

	dataProperty := models.DataProperty{
		Type:   req.Type,
		Title:  req.Title,
		Source: req.Source,
		Status: req.Status,
	}

	result, err := s.DB.SaveChange(&dataProperty)
	if err != nil {
		return nil, err
	}
	return result.(*models.DataProperty), nil
}

// GetByID 获取单条数据
func (s *DataPropertyService) GetByID(id int64) (*models.DataProperty, error) {
	if id <= 0 {
		return nil, errors.New("无效的ID")
	}
	result, err := s.DB.GetOneByID(&models.DataProperty{}, int(id))
	if err != nil {
		return nil, err
	}
	return result.(*models.DataProperty), nil
}

// List 带条件的分页列表
func (s *DataPropertyService) List(req *models.DataPropertyListReq) ([]*models.DataProperty, int64, error) {
	query := global.DB.Model(&models.DataProperty{})

	// 构建查询条件
	if req.Type != "" {
		query = query.Where("`type` = ?", req.Type)
	}
	if req.Status != 0 {
		query = query.Where("`status` = ?", req.Status)
	}

	if req.Keyword != "" {
		keyword := "%" + req.Keyword + "%"
		query = query.Where("title LIKE ? OR `source` LIKE ?",
			keyword, keyword)
	}

	// 获取总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	var list []*models.DataProperty
	err := query.Scopes(s.DB.Paginate(int(req.Page), int(req.PageSize))).
		Order("`time` DESC").
		Find(&list).Error

	return list, total, err
}
