package api

import (
	"github.com/gin-gonic/gin"
	"go-init/global"
	"go-init/models"
	"go-init/service"
)

/**
 * @author: zhashut
 * Date: 2025/4/19
 * Time: 0:38
 * Description: 数据任务-接口层
 */

var (
	dataTaskService = service.NewDataTaskService()
)

// CreateDataTask 创建任务
func CreateDataTask(c *gin.Context) {
	var req models.DataTaskCreateReq
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	result, err := dataTaskService.Create(&req)
	if err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	global.ResponseSuccess(c, result)
}

// ListDataTask 任务列表
func ListDataTask(c *gin.Context) {
	var req models.DataTaskListReq
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	list, total, err := dataTaskService.List(&req)
	if err != nil {
		global.ResponseFailed(c, global.ErrorPERATION)
		return
	}

	global.ResponseSuccess(c, gin.H{
		"list":  list,
		"total": total,
	})
}

// UpdateTaskStatus 更新状态
func UpdateTaskStatus(c *gin.Context) {
	var req models.DataTaskStatusReq
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	if err := dataTaskService.UpdateStatus(req.ID, req.Status); err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	global.ResponseSuccess(c, nil)
}

// UpdateDataTask 更新任务
func UpdateDataTask(c *gin.Context) {
	var req models.DataTaskUpdateReq
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	result, err := dataTaskService.Update(&req)
	if err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	global.ResponseSuccess(c, result)
}

// DeleteDataTask 删除任务
func DeleteDataTask(c *gin.Context) {
	var req models.OnlyIDRequest
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	if err := dataTaskService.DeleteByID(req.ID); err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	global.ResponseSuccess(c, nil)
}
