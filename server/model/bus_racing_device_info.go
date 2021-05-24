// 自动生成模板RacingReveiveMessage
package model

import (
	"gin-vue-admin/global"
      "time"
)

// 如果含有time.Time 请自行import time包
type RacingReveiveMessage struct {
      global.GVA_MODEL
      Imsi  string `json:"imsi" form:"imsi" gorm:"column:imsi;comment:;type:varchar(50);size:50;"`
      Topic  string `json:"topic" form:"topic" gorm:"column:topic;comment:;type:varchar(200);size:200;"`
      Payload  string `json:"payload" form:"payload" gorm:"column:payload;comment:;type:varchar(255);size:255;"`
      MessageData  time.Time `json:"messageData" form:"messageData" gorm:"column:message_data;comment:;type:datetime;"`
      ReceiveDate  time.Time `json:"receiveDate" form:"receiveDate" gorm:"column:receive_date;comment:;type:datetime;"`
      RacePassId  int `json:"racePassId" form:"racePassId" gorm:"column:race_pass_id;comment:;type:int;size:11;"`
      RaceId  int `json:"raceId" form:"raceId" gorm:"column:race_id;comment:;type:int;size:11;"`
}


