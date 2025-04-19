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
 * Time: 0:40
 * Description: 数据报告-接口层
 */

var (
	dataReportService = service.NewDataReportService()
)

// CreateDataReport 生成报告
func CreateDataReport(c *gin.Context) {
	var req models.DataReportCreateReq
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	result, err := dataReportService.Create(&req)
	if err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	global.ResponseSuccess(c, result.ToVo())
}

// ListDataReport 报告列表
func ListDataReport(c *gin.Context) {
	var req models.DataReportListReq
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	list, total, err := dataReportService.List(&req)
	if err != nil {
		global.ResponseFailed(c, global.ErrorPERATION)
		return
	}
	result := make([]*models.DataReportVo, 0, len(list))
	for _, item := range list {
		result = append(result, item.ToVo())
	}

	global.ResponseSuccess(c, gin.H{
		"list":  result,
		"total": total,
	})
}

// UpdateDataReport 更新报告
func UpdateDataReport(c *gin.Context) {
	var req models.DataReportUpdateReq
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	result, err := dataReportService.Update(&req)
	if err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	global.ResponseSuccess(c, result.ToVo())
}

// DeleteDataReport 删除报告
func DeleteDataReport(c *gin.Context) {
	var req models.OnlyIDRequest
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	if err := dataReportService.DeleteByID(req.ID); err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	global.ResponseSuccess(c, nil)
}

// UpdateReportStatus 更新报告状态
func UpdateReportStatus(c *gin.Context) {
	var req models.DataReportStatusReq
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	if err := dataReportService.UpdateStatus(req.ID, req.Status); err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	global.ResponseSuccess(c, nil)
}
