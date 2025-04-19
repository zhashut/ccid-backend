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
 * Date: 2025/4/19
 * Time: 0:03
 * Description: 数据任务-服务层
 */

type DataTaskService struct {
	DB *db.DataTaskDao
}

func NewDataTaskService() *DataTaskService {
	return &DataTaskService{DB: &db.DataTaskDao{}}
}

// Create 创建数据任务
func (s *DataTaskService) Create(task *models.DataTaskCreateReq) (*models.DataTask, error) {
	if task.Name == "" {
		return nil, errors.New("任务名称不能为空")
	}
	if task.Manager == "" {
		return nil, errors.New("负责人不能为空")
	}
	result, err := s.DB.SaveChange(task)
	if err != nil {
		return nil, err
	}
	return result.(*models.DataTask), nil
}

// DeleteByID 删除数据任务
func (s *DataTaskService) DeleteByID(id int64) error {
	if id <= 0 {
		return errors.New("无效的ID")
	}
	return s.DB.DeleteOneByWhere(&models.DataTask{}, fmt.Sprintf("id = %v", id))
}

// Update 更新数据任务
func (s *DataTaskService) Update(task *models.DataTaskUpdateReq) (*models.DataTask, error) {
	if task.ID == 0 {
		return nil, errors.New("更新需要指定ID")
	}
	result, err := s.DB.SaveChange(task)
	if err != nil {
		return nil, err
	}
	return result.(*models.DataTask), nil
}

// GetByID 获取单条任务
func (s *DataTaskService) GetByID(id int64) (*models.DataTask, error) {
	if id <= 0 {
		return nil, errors.New("无效的ID")
	}
	result, err := s.DB.GetOneByID(&models.DataTask{}, int(id))
	if err != nil {
		return nil, err
	}
	return result.(*models.DataTask), nil
}

// List 带条件的分页列表
func (s *DataTaskService) List(req *models.DataTaskListReq) ([]*models.DataTask, int64, error) {
	query := global.DB.Model(&models.DataTask{})

	// 构建查询条件
	if req.Type != "" {
		query = query.Where("type = ?", req.Type)
	}
	if req.Status != 0 {
		query = query.Where("status = ?", req.Status)
	}
	if req.Manager != "" {
		query = query.Where("manager = ?", req.Manager)
	}

	// 模糊搜索条件
	if req.Keyword != "" {
		keyword := "%" + req.Keyword + "%"
		query = query.Where("name LIKE ? OR manager LIKE ?", keyword, keyword)
	}

	// 获取总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	var list []*models.DataTask
	err := query.Scopes(s.DB.Paginate(int(req.Pages), int(req.PageSize))).
		Order("id DESC").
		Find(&list).Error

	return list, total, err
}

// UpdateStatus 更新任务状态
func (s *DataTaskService) UpdateStatus(id int64, status int8) error {
	if id <= 0 {
		return errors.New("无效的ID")
	}
	if status < 1 || status > 2 {
		return errors.New("无效的状态值")
	}
	return global.DB.Model(&models.DataTask{}).
		Where("id = ?", id).
		Update("status", status).Error
}
