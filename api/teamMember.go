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
 * Time: 9:52
 * Description: 团队成员-接口层
 */

var (
	teamMemberService = service.NewTeamMemberService()
)

// CreateTeamMember 创建成员
func CreateTeamMember(c *gin.Context) {
	var req models.TeamMemberCreateReq
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	result, err := teamMemberService.Create(&req)
	if err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	global.ResponseSuccess(c, result.ToVo())
}

// UpdateTeamMember 更新成员
func UpdateTeamMember(c *gin.Context) {
	var req models.TeamMemberUpdateReq
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	result, err := teamMemberService.Update(&req)
	if err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	global.ResponseSuccess(c, result.ToVo())
}

// UpdateMemberRole 更新角色
func UpdateMemberRole(c *gin.Context) {
	var req models.TeamMemberUpdateReq
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	if err := teamMemberService.UpdateRole(req.ID, req.Role); err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	global.ResponseSuccess(c, nil)
}

// DeleteTeamMember 删除成员
func DeleteTeamMember(c *gin.Context) {
	var req models.OnlyIDRequest
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	if err := teamMemberService.DeleteByID(req.ID); err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	global.ResponseSuccess(c, nil)
}

// ListTeamMember 成员列表
func ListTeamMember(c *gin.Context) {
	var req models.TeamMemberListReq
	if err := c.ShouldBind(&req); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	list, total, err := teamMemberService.List(&req)
	if err != nil {
		global.ResponseErrorWithMsg(c, global.ErrorPERATION, err.Error())
		return
	}

	result := make([]*models.TeamMemberVo, 0, len(list))
	for _, item := range list {
		result = append(result, item.ToVo())
	}

	global.ResponseSuccess(c, gin.H{
		"list":  result,
		"total": total,
	})
}
