package models

import "time"

/**
 * @author: zhashut
 * Date: 2025/4/18
 * Time: 23:40
 * Description: 数据报告-模型层
 */

type DataReport struct {
	ID      int64     `gorm:"primaryKey;autoIncrement" json:"id"`            // 报告ID
	Name    string    `gorm:"type:varchar(255);not null" json:"name"`        // 报告名称
	Type    string    `gorm:"type:varchar(255);not null" json:"type"`        // 报告类型（专利分析/技术分析/市场分析）
	Status  int       `gorm:"type:tinyint;not null;default:1" json:"status"` // 状态（1: 进行中, 2: 已完成）
	Time    time.Time `gorm:"type:date;not null" json:"time"`                // 生成时间
	Content string    `gorm:"type:text;not null" json:"content"`             // 报告内容
}

type DataReportCreateReq struct {
	Name    string `json:"name" binding:"required"`
	Type    string `json:"type" binding:"required"`
	Status  int    `json:"status"`
	Content string `json:"content" binding:"required"`
}

type DataReportUpdateReq struct {
	ID int64 `json:"id" binding:"required"`
	DataReportCreateReq
}

type DataReportStatusReq struct {
	OnlyIDRequest
	Status int `json:"status" binding:"required"`
}

type DataReportListReq struct {
	PageRequest
	Type      string `form:"type"`
	Status    int    `form:"status"`
	Keyword   string `form:"keyword"`
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
}

type DataReportVo struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Status  string `json:"status"`
	Content string `json:"content"`
	Time    string `json:"time"`
}

var dataReportStatusMap = map[int]string{
	1: "进行中",
	2: "已完成",
}

func (d *DataReport) ToVo() *DataReportVo {
	return &DataReportVo{
		ID:      d.ID,
		Name:    d.Name,
		Type:    d.Type,
		Status:  dataReportStatusMap[d.Status],
		Content: d.Content,
		Time:    d.Time.Format("2006-01-02"),
	}
}
