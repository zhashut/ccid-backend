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
 * Date: 10:23
 * Description: 数据专利-服务层
 */

type DataPatentService struct {
	DB *db.DataPatentDao
}

func NewDataPatentService() *DataPatentService {
	return &DataPatentService{DB: &db.DataPatentDao{}}
}

// Create 创建专利
func (s *DataPatentService) Create(patent *models.DataPatentCreateReq) (*models.DataPatent, error) {
	if patent.Title == "" {
		return nil, errors.New("专利标题不能为空")
	}

	result, err := s.DB.SaveChange(patent)
	if err != nil {
		return nil, err
	}
	return result.(*models.DataPatent), nil
}

// DeleteByID 删除专利
func (s *DataPatentService) DeleteByID(id int64) error {
	if id == 0 {
		return errors.New("专利号不能为空")
	}
	return s.DB.DeleteOneByWhere(&models.DataPatent{}, fmt.Sprintf("id = %v", id))
}

// GetByID 获取专利详情
func (s *DataPatentService) GetByID(id int64) (*models.DataPatent, error) {
	if id == 0 {
		return nil, errors.New("专利号不能为空")
	}
	result, err := s.DB.GetOneByID(&models.DataPatent{}, int(id))
	if err != nil {
		return nil, err
	}
	return result.(*models.DataPatent), nil
}

// Update 更新专利
func (s *DataPatentService) Update(patent *models.DataPatentUpdateReq) (*models.DataPatent, error) {
	if patent.ID == 0 {
		return nil, errors.New("专利号不能为空")
	}
	result, err := s.DB.SaveChange(patent)
	if err != nil {
		return nil, err
	}
	return result.(*models.DataPatent), nil
}

// List 获取专利列表
func (s *DataPatentService) List(req *models.DataPatentListReq) ([]*models.DataPatent, int64, error) {
	query := global.DB.Model(&models.DataPatent{})

	if req.Keyword != "" {
		query = query.Where("title LIKE ? OR applicant LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}
	if req.Field != "" {
		query = query.Where("field = ?", req.Field)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var list []*models.DataPatent
	err := query.Scopes(s.DB.Paginate(int(req.Pages), int(req.PageSize))).
		Order("`date` DESC").
		Find(&list).Error

	return list, total, err
}
