package model

import (
	"context"

	"gorm.io/gorm"
)

type UserPropEntity struct {
	// 场景
	Scene int64
	// 用户id
	UserId int64
	// 创建时间戳
	CreateTs int64
	// 道具id
	PropId int64
	// 道具数量
	PropNum int64
	// 用户签到id, 非用户签到为0
	UserSignId int64
}

type UserPropRepository interface {
	Create(ctx context.Context, tx *gorm.DB, userProp *UserPropEntity)
}
