package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type UserSignTool interface {
	// 用户签到
	UserSign(ctx context.Context, user UserRepository, sign SignConfigInterface, ts time.Time) *UserSignResp
	// 用户在当前时间，当前签到配置是否已经签到了
	UserIsSigned(ctx context.Context, user UserRepository, sign SignConfigInterface, ts time.Time) bool
}

type UserSignInterface interface {
	GetId() int64
	GetUserId() int64
	GetSignConfigId() int64
	GetSignTime() time.Time
}

type UserSignRepository interface {
	Create(ctx context.Context, tx *gorm.DB, userSign UserSignInterface, t time.Time)
}

type UserSignResp struct {
	// 是否已经签到过， true已经签到过，false未签到
	IsSigned bool `json:"is_signed"`
}
