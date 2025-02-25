package model

type SignConfigInterface interface {
	SceneInterface
	// 签到类型 0中断后继续（不重置）， 1中断后重置
	GetSignType() int64
	// 签到单位，0每天，1每自然周，2每自然月
	GetSignUnit() int64
	// 有效期 开始时间戳，0表示不限
	GetStartTs() int64
	// 有效期 结束时间戳，0表示不限
	GetEndTs() int64
	// 签到奖励配置
	GetSignRewardConfig() []SignRewardConfig
}

// 签到道具奖励
type SignRewardConfig struct {
	// 道具信息
	PropEntity
	// 道具发放数量
	RewardPropNum int64
}
