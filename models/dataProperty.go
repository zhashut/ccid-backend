package models

import "time"

/**
 * @author: zhashut
 * Date: 2025/4/18
 * Time: 23:39
 * Description: 知识产权信息-模型层
 */

type DataProperty struct {
	ID             int64     `gorm:"primaryKey;autoIncrement" json:"id"`            // 数据ID
	Type           string    `gorm:"type:varchar(255);not null" json:"type"`        // 数据类型（专利/论文/软件著作权）
	Title          string    `gorm:"type:varchar(255);not null" json:"title"`       // 标题
	Source         string    `gorm:"type:varchar(255);not null" json:"source"`      // 来源
	CollectTime    time.Time `gorm:"type:date" json:"collect_time"`                 // 采集时间
	Status         int       `gorm:"type:tinyint;not null;default:1" json:"status"` // 状态（1: 进行中, 2: 已完成）
	PatentID       string    `gorm:"type:varchar(255);not null" json:"patent_id"`   // 专利号
	Applier        string    `gorm:"type:varchar(255);not null" json:"applier"`     // 申请人
	ApplyTime      time.Time `gorm:"type:date;not null" json:"apply_time"`          // 申请日期
	TechnicalField string    `gorm:"type:varchar(255)" json:"technical_field"`      // 技术领域（如机器学习/计算机视觉）
}

type DataPropertyCreateReq struct {
	Type           string `json:"type" binding:"required"`
	Title          string `json:"title" binding:"required"`
	Source         string `json:"source" binding:"required"`
	CollectTime    string `json:"collect_time"`
	Status         int    `json:"status"`
	PatentID       string `json:"patent_id" binding:"required"`
	Applier        string `json:"applier" binding:"required"`
	ApplyTime      string `json:"apply_time" binding:"required"`
	TechnicalField string `json:"technical_field"`
}

type DataPropertyUpdateReq struct {
	ID int64 `json:"id" binding:"required"`
	DataPropertyCreateReq
}

type DataPropertyListReq struct {
	PageRequest
	Keyword        string `form:"keyword"`
	Type           string `form:"type"`
	Status         int    `form:"status"`
	TechnicalField string `form:"technical_field"`
	Applier        string `form:"applier"`
	PatentID       string `form:"patent_id"`
}

// PatentListReq 专利列表专用请求参数
type PatentListReq struct {
	PageRequest
	Keyword        string `form:"keyword"`         // 模糊搜索关键词
	TechnicalField string `form:"technical_field"` // 技术领域筛选
}

// PatentListVo 专利列表展示VO
type PatentListVo struct {
	PatentID       string `json:"patent_id"`
	Title          string `json:"title"`
	Applier        string `json:"applier"`
	ApplyTime      string `json:"apply_time"`
	TechnicalField string `json:"technical_field"`
}

type DataPropertyVo struct {
	ID             int64  `json:"id"`
	Type           string `json:"type"`
	Title          string `json:"title"`
	Source         string `json:"source"`
	CollectTime    string `json:"collect_time"` // yyyy-MM-dd 格式
	Status         int    `json:"status"`
	PatentID       string `json:"patent_id"`
	Applier        string `json:"applier"`
	ApplyTime      string `json:"apply_time"` // yyyy-MM-dd 格式
	TechnicalField string `json:"technical_field"`
}

func (d *DataProperty) ToVo() *DataPropertyVo {
	return &DataPropertyVo{
		ID:             d.ID,
		Type:           d.Type,
		Title:          d.Title,
		Source:         d.Source,
		CollectTime:    d.CollectTime.Format("2006-01-02"),
		Status:         d.Status,
		PatentID:       d.PatentID,
		Applier:        d.Applier,
		ApplyTime:      d.ApplyTime.Format("2006-01-02"),
		TechnicalField: d.TechnicalField,
	}
}

func (d *DataProperty) ToPatentVo() *PatentListVo {
	return &PatentListVo{
		Title:          d.Title,
		PatentID:       d.PatentID,
		Applier:        d.Applier,
		ApplyTime:      d.ApplyTime.Format("2006-01-02"),
		TechnicalField: d.TechnicalField,
	}
}
