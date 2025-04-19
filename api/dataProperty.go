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
 * Time: 0:19
 * Description: 数据任务-服务层
 */

var (
	dataPropertyService = service.NewDataPropertyService()
)

// CreateDataProperty 创建数据
func CreateDataProperty(c *gin.Context) {
	var req models.DataPropertyCreateReq
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	result, err := dataPropertyService.Create(&req)
	if err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	global.ResponseSuccess(c, result.ToVo())
}

// ListDataProperty 获取列表
func ListDataProperty(c *gin.Context) {
	var req models.DataPropertyListReq
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	list, total, err := dataPropertyService.List(&req)
	if err != nil {
		global.ResponseFailed(c, global.ErrorPERATION)
		return
	}

	result := make([]*models.DataPropertyVo, 0, len(list))
	for _, item := range list {
		result = append(result, item.ToVo())
	}

	global.ResponseSuccess(c, gin.H{
		"list":  result,
		"total": total,
	})
}

// UpdateDataProperty 更新数据
func UpdateDataProperty(c *gin.Context) {
	var req models.DataPropertyUpdateReq
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	result, err := dataPropertyService.Update(&req)
	if err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	global.ResponseSuccess(c, result.ToVo())
}

// DeleteDataProperty 删除数据
func DeleteDataProperty(c *gin.Context) {
	var req models.OnlyIDRequest
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	if err := dataPropertyService.DeleteByID(req.ID); err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	global.ResponseSuccess(c, nil)
}
