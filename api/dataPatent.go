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
 * Time: 10:28
 * Description: 数据专利-服务层
 */

var (
	dataPatentService = service.NewDataPatentService()
)

func CreateDataPatent(c *gin.Context) {
	var req models.DataPatentCreateReq
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	patent, err := dataPatentService.Create(&req)
	if err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	global.ResponseSuccess(c, patent.ToVo())
}

func UpdateDataPatent(c *gin.Context) {
	var req models.DataPatentUpdateReq
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	patent, err := dataPatentService.Update(&req)
	if err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	global.ResponseSuccess(c, patent.ToVo())
}

func DeleteDataPatent(c *gin.Context) {
	var req models.OnlyIDRequest
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	if err := dataPatentService.DeleteByID(req.ID); err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	global.ResponseSuccess(c, nil)
}

func ListDataPatent(c *gin.Context) {
	var req models.DataPatentListReq
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	list, total, err := dataPatentService.List(&req)
	if err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	result := make([]*models.DataPatentVo, 0, len(list))
	for _, p := range list {
		result = append(result, p.ToVo())
	}

	global.ResponseSuccess(c, gin.H{
		"list":  result,
		"total": total,
	})
}
