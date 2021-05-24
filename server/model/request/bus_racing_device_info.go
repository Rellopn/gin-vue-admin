package request

import "gin-vue-admin/model"

type RacingReveiveMessageSearch struct{
    model.RacingReveiveMessage
    PageInfo
}