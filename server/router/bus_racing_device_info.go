package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRacingReveiveMessageRouter(Router *gin.RouterGroup) {
	RacingReveiveMessageRouter := Router.Group("racingReveiveMessage").Use(middleware.OperationRecord())
	{
		RacingReveiveMessageRouter.POST("createRacingReveiveMessage", v1.CreateRacingReveiveMessage)   // 新建RacingReveiveMessage
		RacingReveiveMessageRouter.DELETE("deleteRacingReveiveMessage", v1.DeleteRacingReveiveMessage) // 删除RacingReveiveMessage
		RacingReveiveMessageRouter.DELETE("deleteRacingReveiveMessageByIds", v1.DeleteRacingReveiveMessageByIds) // 批量删除RacingReveiveMessage
		RacingReveiveMessageRouter.PUT("updateRacingReveiveMessage", v1.UpdateRacingReveiveMessage)    // 更新RacingReveiveMessage
		RacingReveiveMessageRouter.GET("findRacingReveiveMessage", v1.FindRacingReveiveMessage)        // 根据ID获取RacingReveiveMessage
		RacingReveiveMessageRouter.GET("getRacingReveiveMessageList", v1.GetRacingReveiveMessageList)  // 获取RacingReveiveMessage列表
	}
}
