package models

/**
 * @author: zhashut
 * Date: 2025/4/19
 * Time: 9:34
 * Description: 团队成员-模型层
 */

type TeamMember struct {
	ID   int64  `gorm:"primaryKey;autoIncrement" json:"id"`          // 成员ID
	Name string `gorm:"type:varchar(50);not null" json:"name"`       // 姓名
	Role int    `gorm:"type:tinyint;not null;default:1" json:"role"` // 角色（1:查看,2:编辑,3:管理）
}

type TeamMemberCreateReq struct {
	Name string `json:"name" binding:"required,min=2,max=50"`
	Role int    `json:"role" binding:"required,min=1,max=3"`
}

type TeamMemberUpdateReq struct {
	ID   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required,min=2,max=50"`
	Role int    `json:"role" binding:"required,min=1,max=3"`
}

type TeamMemberListReq struct {
	PageRequest
	Name string `form:"name"`
	Role int    `form:"role"`
}

type TeamMemberVo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Role int    `json:"role"` // 返回角色文字描述
}

var roleMap = map[int]string{
	1: "查看",
	2: "编辑",
	3: "管理",
}

func (t *TeamMember) ToVo() *TeamMemberVo {
	return &TeamMemberVo{
		ID:   t.ID,
		Name: t.Name,
		Role: t.Role,
	}
}
