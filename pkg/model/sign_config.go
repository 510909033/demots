package model

import (
	"context"

	"gorm.io/gorm"
)

type SignConfigEntity struct {
	// 签到配置id
	Id int64
	// 场景
	Scene int64
	// 签到类型 0中断后继续（不重置）， 1中断后重置
	SignType int64
	// 签到单位，0每天，1每自然周，2每自然月
	SignUnit int64
	// 有效期 开始时间戳，0表示不限
	StartTs int64
	// 有效期 结束时间戳，0表示不限
	EndTs int64
	// 签到奖励配置
	SignRewardConfig []SignRewardConfig
}

// 签到道具奖励
type SignRewardConfig struct {
	// 道具信息
	PropEntity
	// 道具发放数量
	RewardPropNum int64
}

type SignConfigRepository interface {
	// 根据场景获取签到配置
	GetByScent(ctx context.Context, tx *gorm.DB, scene int64) *SignConfigEntity
	// 创建签到配置
	Create(ctx context.Context, tx *gorm.DB, signConfig *SignConfigEntity)
}
