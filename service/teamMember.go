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
 * Time: 9:39
 * Description: 团队成员-服务层
 */

type TeamMemberService struct {
	DB *db.TeamMemberDao
}

func NewTeamMemberService() *TeamMemberService {
	return &TeamMemberService{DB: &db.TeamMemberDao{}}
}

// Create 创建成员
func (s *TeamMemberService) Create(member *models.TeamMemberCreateReq) (*models.TeamMember, error) {
	if member.Name == "" {
		return nil, errors.New("成员姓名不能为空")
	}
	if member.Role < 1 || member.Role > 3 {
		return nil, errors.New("无效的角色值")
	}
	result, err := s.DB.SaveChange(member)
	if err != nil {
		return nil, err
	}
	return result.(*models.TeamMember), nil
}

// DeleteByID 删除成员
func (s *TeamMemberService) DeleteByID(id int64) error {
	if id <= 0 {
		return errors.New("无效的ID")
	}
	return s.DB.DeleteOneByWhere(&models.TeamMember{}, fmt.Sprintf("id = %v", id))
}

// Update 更新成员
func (s *TeamMemberService) Update(member *models.TeamMemberUpdateReq) (*models.TeamMember, error) {
	if member.ID == 0 {
		return nil, errors.New("更新需要指定ID")
	}
	result, err := s.DB.SaveChange(member)
	if err != nil {
		return nil, err
	}
	return result.(*models.TeamMember), nil
}

// GetByID 获取单个成员
func (s *TeamMemberService) GetByID(id int64) (*models.TeamMember, error) {
	if id <= 0 {
		return nil, errors.New("无效的ID")
	}
	result, err := s.DB.GetOneByID(&models.TeamMember{}, int(id))
	if err != nil {
		return nil, err
	}
	return result.(*models.TeamMember), nil
}

// List 带分页的成员列表
func (s *TeamMemberService) List(req *models.TeamMemberListReq) ([]*models.TeamMember, int64, error) {
	query := global.DB.Model(&models.TeamMember{})

	if req.Name != "" {
		query = query.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Role != 0 {
		query = query.Where("role = ?", req.Role)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var list []*models.TeamMember
	err := query.Scopes(s.DB.Paginate(int(req.Page), int(req.PageSize))).
		Order("id DESC").
		Find(&list).Error

	return list, total, err
}

// UpdateRole 更新成员角色
func (s *TeamMemberService) UpdateRole(id int64, role int) error {
	if id <= 0 {
		return errors.New("无效的ID")
	}
	if role < 1 || role > 3 {
		return errors.New("无效的角色值")
	}
	return global.DB.Model(&models.TeamMember{}).
		Where("id = ?", id).
		Update("role", role).Error
}
