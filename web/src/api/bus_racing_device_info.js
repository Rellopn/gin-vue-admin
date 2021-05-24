import service from '@/utils/request'

// @Tags RacingReveiveMessage
// @Summary 创建RacingReveiveMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RacingReveiveMessage true "创建RacingReveiveMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /racingReveiveMessage/createRacingReveiveMessage [post]
export const createRacingReveiveMessage = (data) => {
     return service({
         url: "/racingReveiveMessage/createRacingReveiveMessage",
         method: 'post',
         data
     })
 }


// @Tags RacingReveiveMessage
// @Summary 删除RacingReveiveMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RacingReveiveMessage true "删除RacingReveiveMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /racingReveiveMessage/deleteRacingReveiveMessage [delete]
 export const deleteRacingReveiveMessage = (data) => {
     return service({
         url: "/racingReveiveMessage/deleteRacingReveiveMessage",
         method: 'delete',
         data
     })
 }

// @Tags RacingReveiveMessage
// @Summary 删除RacingReveiveMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RacingReveiveMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /racingReveiveMessage/deleteRacingReveiveMessage [delete]
 export const deleteRacingReveiveMessageByIds = (data) => {
     return service({
         url: "/racingReveiveMessage/deleteRacingReveiveMessageByIds",
         method: 'delete',
         data
     })
 }

// @Tags RacingReveiveMessage
// @Summary 更新RacingReveiveMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RacingReveiveMessage true "更新RacingReveiveMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /racingReveiveMessage/updateRacingReveiveMessage [put]
 export const updateRacingReveiveMessage = (data) => {
     return service({
         url: "/racingReveiveMessage/updateRacingReveiveMessage",
         method: 'put',
         data
     })
 }


// @Tags RacingReveiveMessage
// @Summary 用id查询RacingReveiveMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RacingReveiveMessage true "用id查询RacingReveiveMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /racingReveiveMessage/findRacingReveiveMessage [get]
 export const findRacingReveiveMessage = (params) => {
     return service({
         url: "/racingReveiveMessage/findRacingReveiveMessage",
         method: 'get',
         params
     })
 }


// @Tags RacingReveiveMessage
// @Summary 分页获取RacingReveiveMessage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取RacingReveiveMessage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /racingReveiveMessage/getRacingReveiveMessageList [get]
 export const getRacingReveiveMessageList = (params) => {
     return service({
         url: "/racingReveiveMessage/getRacingReveiveMessageList",
         method: 'get',
         params
     })
 }