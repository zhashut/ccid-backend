package models

/**
 * @author: zhashut
 * Date: 2025/4/18
 * Time: 23:39
 * Description: 数据任务-模型层
 */

type DataTask struct {
	ID      int64  `gorm:"primaryKey;autoIncrement" json:"id"`            // 任务ID
	Name    string `gorm:"type:varchar(255);not null" json:"name"`        // 任务名称（如专利数据清洗）
	Type    string `gorm:"type:varchar(255);not null" json:"type"`        // 数据类型（专利/论文）
	Manager string `gorm:"type:varchar(255);not null" json:"manager"`     // 负责人
	Status  int    `gorm:"type:tinyint;not null;default:1" json:"status"` // 状态（1: 进行中, 2: 已完成）
}

type DataTaskCreateReq struct {
	Name    string `json:"name" binding:"required"`
	Type    string `json:"type" binding:"required"`
	Manager string `json:"manager" binding:"required"`
	Status  int    `json:"status"`
}

type DataTaskUpdateReq struct {
	ID int64 `json:"id" binding:"required"`
	DataTaskCreateReq
}

type DataTaskStatusReq struct {
	OnlyIDRequest
	Status int `json:"status" binding:"required"`
}

type DataTaskListReq struct {
	PageRequest
	Type    string `form:"type"`
	Status  int    `form:"status"`
	Manager string `form:"manager"`
	Keyword string `form:"keyword"`
}

type DataTaskVo struct {
	*DataTask
}
