package models

import (
	"time"
)

/**
 * @author: zhashut
 * Date: 2025/4/19
 * Date: 10:16
 * Description: 数据专利-模型层
 */

// DataPatent 专利数据表
type DataPatent struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`          // ID
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`     // 专利名称
	Applicant string    `gorm:"type:varchar(255);not null" json:"applicant"` // 申请人
	Date      time.Time `gorm:"type:date;not null" json:"date"`              // 申请日期
	Field     string    `gorm:"type:varchar(255);not null" json:"field"`     // 技术领域
}

// DataPatentListReq 专利列表专用请求参数
type DataPatentListReq struct {
	PageRequest
	Keyword string `form:"keyword"` // 模糊搜索关键词
	Field   string `form:"field"`   // 技术领域筛选
}

// DataPatentVo 专利列表展示VO
type DataPatentVo struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Applicant string `json:"applicant"`
	Date      string `json:"date"`
	Field     string `json:"field"`
}

// ToVo 转换为展示VO
func (d *DataPatent) ToVo() *DataPatentVo {
	return &DataPatentVo{
		ID:        d.ID,
		Title:     d.Title,
		Applicant: d.Applicant,
		Date:      d.Date.Format("2006-01-02"),
		Field:     d.Field,
	}
}

// DataPatentCreateReq 创建专利请求
type DataPatentCreateReq struct {
	ID      int64  `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Applier string `json:"applier" binding:"required"`
	Date    string `json:"date" binding:"required"` // 格式: 2006-01-02
	Field   string `json:"field" binding:"required"`
}

// DataPatentUpdateReq 更新专利请求
type DataPatentUpdateReq struct {
	ID        int64  `json:"id" binding:"required"`
	Title     string `json:"title"`
	Applicant string `json:"applicant"`
	Date      string `json:"date"` // 格式: 2006-01-02
	Field     string `json:"field"`
}
