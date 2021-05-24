package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateRacingReveiveMessage
//@description: 创建RacingReveiveMessage记录
//@param: racingReveiveMessage model.RacingReveiveMessage
//@return: err error

func CreateRacingReveiveMessage(racingReveiveMessage model.RacingReveiveMessage) (err error) {
	err = global.GVA_DB.Create(&racingReveiveMessage).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRacingReveiveMessage
//@description: 删除RacingReveiveMessage记录
//@param: racingReveiveMessage model.RacingReveiveMessage
//@return: err error

func DeleteRacingReveiveMessage(racingReveiveMessage model.RacingReveiveMessage) (err error) {
	err = global.GVA_DB.Delete(&racingReveiveMessage).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRacingReveiveMessageByIds
//@description: 批量删除RacingReveiveMessage记录
//@param: ids request.IdsReq
//@return: err error

func DeleteRacingReveiveMessageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.RacingReveiveMessage{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateRacingReveiveMessage
//@description: 更新RacingReveiveMessage记录
//@param: racingReveiveMessage *model.RacingReveiveMessage
//@return: err error

func UpdateRacingReveiveMessage(racingReveiveMessage model.RacingReveiveMessage) (err error) {
	err = global.GVA_DB.Save(&racingReveiveMessage).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRacingReveiveMessage
//@description: 根据id获取RacingReveiveMessage记录
//@param: id uint
//@return: err error, racingReveiveMessage model.RacingReveiveMessage

func GetRacingReveiveMessage(id uint) (err error, racingReveiveMessage model.RacingReveiveMessage) {
	err = global.GVA_DB.Where("id = ?", id).First(&racingReveiveMessage).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRacingReveiveMessageInfoList
//@description: 分页获取RacingReveiveMessage记录
//@param: info request.RacingReveiveMessageSearch
//@return: err error, list interface{}, total int64

func GetRacingReveiveMessageInfoList(info request.RacingReveiveMessageSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.RacingReveiveMessage{})
    var racingReveiveMessages []model.RacingReveiveMessage
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Imsi != "" {
        db = db.Where("`imsi` LIKE ?","%"+ info.Imsi+"%")
    }
    if info.Topic != "" {
        db = db.Where("`topic` LIKE ?","%"+ info.Topic+"%")
    }
    if info.Payload != "" {
        db = db.Where("`payload` LIKE ?","%"+ info.Payload+"%")
    }
    if !info.MessageData.IsZero() {
         db = db.Where("`message_data` <> ?",info.MessageData)
    }
    if !info.ReceiveDate.IsZero() {
         db = db.Where("`receive_date` <> ?",info.ReceiveDate)
    }
    if info.RacePassId != 0 {
        db = db.Where("`race_pass_id` = ?",info.RacePassId)
    }
    if info.RaceId != 0 {
        db = db.Where("`race_id` = ?",info.RaceId)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&racingReveiveMessages).Error
	return err, racingReveiveMessages, total
}