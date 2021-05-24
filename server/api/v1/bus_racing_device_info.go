package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags RacingReveiveMessage
// @Summary 创建RacingReveiveMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RacingReveiveMessage true "创建RacingReveiveMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /racingReveiveMessage/createRacingReveiveMessage [post]
func CreateRacingReveiveMessage(c *gin.Context) {
	var racingReveiveMessage model.RacingReveiveMessage
	_ = c.ShouldBindJSON(&racingReveiveMessage)
	if err := service.CreateRacingReveiveMessage(racingReveiveMessage); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags RacingReveiveMessage
// @Summary 删除RacingReveiveMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RacingReveiveMessage true "删除RacingReveiveMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /racingReveiveMessage/deleteRacingReveiveMessage [delete]
func DeleteRacingReveiveMessage(c *gin.Context) {
	var racingReveiveMessage model.RacingReveiveMessage
	_ = c.ShouldBindJSON(&racingReveiveMessage)
	if err := service.DeleteRacingReveiveMessage(racingReveiveMessage); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags RacingReveiveMessage
// @Summary 批量删除RacingReveiveMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RacingReveiveMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /racingReveiveMessage/deleteRacingReveiveMessageByIds [delete]
func DeleteRacingReveiveMessageByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteRacingReveiveMessageByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags RacingReveiveMessage
// @Summary 更新RacingReveiveMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RacingReveiveMessage true "更新RacingReveiveMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /racingReveiveMessage/updateRacingReveiveMessage [put]
func UpdateRacingReveiveMessage(c *gin.Context) {
	var racingReveiveMessage model.RacingReveiveMessage
	_ = c.ShouldBindJSON(&racingReveiveMessage)
	if err := service.UpdateRacingReveiveMessage(racingReveiveMessage); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags RacingReveiveMessage
// @Summary 用id查询RacingReveiveMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RacingReveiveMessage true "用id查询RacingReveiveMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /racingReveiveMessage/findRacingReveiveMessage [get]
func FindRacingReveiveMessage(c *gin.Context) {
	var racingReveiveMessage model.RacingReveiveMessage
	_ = c.ShouldBindQuery(&racingReveiveMessage)
	if err, reracingReveiveMessage := service.GetRacingReveiveMessage(racingReveiveMessage.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reracingReveiveMessage": reracingReveiveMessage}, c)
	}
}

// @Tags RacingReveiveMessage
// @Summary 分页获取RacingReveiveMessage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.RacingReveiveMessageSearch true "分页获取RacingReveiveMessage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /racingReveiveMessage/getRacingReveiveMessageList [get]
func GetRacingReveiveMessageList(c *gin.Context) {
	var pageInfo request.RacingReveiveMessageSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetRacingReveiveMessageInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
