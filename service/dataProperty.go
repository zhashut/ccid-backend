package service

import (
	"errors"
	"fmt"
	"go-init/global"
	"go-init/models"
	"go-init/respository/db"
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
func (s *DataPropertyService) Create(property *models.DataPropertyCreateReq) (*models.DataProperty, error) {
	if property.Title == "" {
		return nil, errors.New("标题不能为空")
	}
	result, err := s.DB.SaveChange(property)
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
func (s *DataPropertyService) Update(property *models.DataPropertyUpdateReq) (*models.DataProperty, error) {
	if property.ID == 0 {
		return nil, errors.New("更新需要指定ID")
	}
	result, err := s.DB.SaveChange(property)
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
		query = query.Where("type = ?", req.Type)
	}
	if req.Status != 0 {
		query = query.Where("status = ?", req.Status)
	}
	if req.TechnicalField != "" {
		query = query.Where("technical_field LIKE ?", "%"+req.TechnicalField+"%")
	}
	if req.Applier != "" {
		query = query.Where("applier = ?", req.Applier)
	}
	if req.PatentID != "" {
		query = query.Where("patent_id = ?", req.PatentID)
	}

	if req.Keyword != "" {
		keyword := "%" + req.Keyword + "%"
		query = query.Where("title LIKE ? OR `source` LIKE ? LIKE ?",
			keyword, keyword)
	}

	// 获取总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	var list []*models.DataProperty
	err := query.Scopes(s.DB.Paginate(int(req.Pages), int(req.PageSize))).
		Order("apply_time DESC").
		Find(&list).Error

	return list, total, err
}

// PatentList 专利专用列表查询
func (s *DataPropertyService) PatentList(req *models.PatentListReq) ([]*models.DataProperty, int64, error) {
	query := global.DB.Model(&models.DataProperty{}).
		Select("patent_id, title, applier, apply_time, technical_field").
		Where("type = ?", "专利")

	// 模糊搜索条件
	if req.Keyword != "" {
		keyword := "%" + req.Keyword + "%"
		query = query.Where(
			"patent_id LIKE ? OR title LIKE ? OR applier LIKE ? OR technical_field LIKE ?",
			keyword, keyword, keyword, keyword,
		)
	}

	// 技术领域筛选
	if req.TechnicalField != "" {
		query = query.Where("technical_field = ?", req.TechnicalField)
	}

	// 获取总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	var list []*models.DataProperty
	err := query.Scopes(s.DB.Paginate(int(req.Pages), int(req.PageSize))).
		Order("apply_time DESC").
		Find(&list).Error

	return list, total, err
}
