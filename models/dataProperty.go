package models

import "time"

/**
 * @author: zhashut
 * Date: 2025/4/18
 * Time: 23:39
 * Description: 知识产权信息-模型层
 */

type DataProperty struct {
	ID     int64     `gorm:"primaryKey;autoIncrement" json:"id"`            // 数据ID
	Type   string    `gorm:"type:varchar(255);not null" json:"type"`        // 数据类型（专利/论文/软件著作权）
	Title  string    `gorm:"type:varchar(255);not null" json:"title"`       // 标题
	Source string    `gorm:"type:varchar(255);not null" json:"source"`      // 来源
	Time   time.Time `gorm:"type:date" json:"time"`                         // 采集时间
	Status int       `gorm:"type:tinyint;not null;default:1" json:"status"` // 状态（1: 进行中, 2: 已完成）
}

type DataPropertyCreateReq struct {
	Type   string `json:"type" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Source string `json:"source" binding:"required"`
	Status int    `json:"status"`
}

type DataPropertyUpdateReq struct {
	ID int64 `json:"id" binding:"required"`
	DataPropertyCreateReq
}

type DataPropertyListReq struct {
	PageRequest
	Keyword string `form:"keyword"`
	Type    string `form:"type"`
	Status  int    `form:"status"`
}

type DataPropertyVo struct {
	ID     int64  `json:"id"`
	Type   string `json:"type"`
	Title  string `json:"title"`
	Source string `json:"source"`
	Time   string `json:"time"` // yyyy-MM-dd 格式
	Status int    `json:"status"`
}

var dataPropertyStatusMap = map[int]string{
	1: "进行中",
	2: "已完成",
}

func (d *DataProperty) ToVo() *DataPropertyVo {
	return &DataPropertyVo{
		ID:     d.ID,
		Type:   d.Type,
		Title:  d.Title,
		Source: d.Source,
		Time:   d.Time.Format("2006-01-02"),
		Status: d.Status,
	}
}
