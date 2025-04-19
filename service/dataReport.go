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
 * Date: 2025/4/19
 * Time: 0:12
 * Description: 数据报告-服务层
 */

type DataReportService struct {
	DB *db.DataReportDao
}

func NewDataReportService() *DataReportService {
	return &DataReportService{DB: &db.DataReportDao{}}
}

// Create 创建数据报告
func (s *DataReportService) Create(req *models.DataReportCreateReq) (*models.DataReport, error) {
	if req.Name == "" {
		return nil, errors.New("报告名称不能为空")
	}
	if req.Type == "" {
		return nil, errors.New("报告类型不能为空")
	}
	report := models.DataReport{
		Name:    req.Name,
		Type:    req.Type,
		Status:  req.Status,
		Time:    time.Now(),
		Content: req.Content,
	}
	result, err := s.DB.SaveChange(&report)
	if err != nil {
		return nil, err
	}
	return result.(*models.DataReport), nil
}

// DeleteByID 删除数据报告
func (s *DataReportService) DeleteByID(id int64) error {
	if id <= 0 {
		return errors.New("无效的ID")
	}
	return s.DB.DeleteOneByWhere(&models.DataReport{}, fmt.Sprintf("id = %v", id))
}

// Update 更新数据报告
func (s *DataReportService) Update(req *models.DataReportUpdateReq) (*models.DataReport, error) {
	if req.ID == 0 {
		return nil, errors.New("更新需要指定ID")
	}
	report := models.DataReport{
		Name:    req.Name,
		Type:    req.Type,
		Status:  req.Status,
		Content: req.Content,
	}
	result, err := s.DB.SaveChange(&report)
	if err != nil {
		return nil, err
	}
	return result.(*models.DataReport), nil
}

// GetByID 获取单条报告
func (s *DataReportService) GetByID(id int64) (*models.DataReport, error) {
	if id <= 0 {
		return nil, errors.New("无效的ID")
	}
	result, err := s.DB.GetOneByID(&models.DataReport{}, int(id))
	if err != nil {
		return nil, err
	}
	return result.(*models.DataReport), nil
}

// List 带条件的分页列表（支持时间范围）
func (s *DataReportService) List(req *models.DataReportListReq) ([]*models.DataReport, int64, error) {
	query := global.DB.Model(&models.DataReport{})

	// 精确匹配条件
	if req.Type != "" {
		query = query.Where("`type` = ?", req.Type)
	}
	if req.Status != 0 {
		query = query.Where("`status` = ?", req.Status)
	}

	// 时间范围条件
	if req.StartTime != "" {
		query = query.Where("`time` >= ?", req.StartTime)
	}
	if req.EndTime != "" {
		query = query.Where("`time` <= ?", req.EndTime)
	}

	// 模糊搜索条件
	if req.Keyword != "" {
		keyword := "%" + req.Keyword + "%"
		query = query.Where("name LIKE ? OR `type` LIKE ?", keyword, keyword)
	}

	// 获取总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	var list []*models.DataReport
	err := query.Scopes(s.DB.Paginate(int(req.Pages), int(req.PageSize))).
		Order("`time` DESC").
		Find(&list).Error

	return list, total, err
}

// UpdateStatus 更新报告状态
func (s *DataReportService) UpdateStatus(id int64, status int) error {
	if id <= 0 {
		return errors.New("无效的ID")
	}
	if status < 1 || status > 2 {
		return errors.New("无效的状态值")
	}
	return global.DB.Model(&models.DataReport{}).
		Where("id = ?", id).
		Update("`status`", status).Error
}
