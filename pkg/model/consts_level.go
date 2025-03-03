package model

// 道具相关
const (
// 道具类型, 0异常， 1经验值， 2经验球
// ---

// PropTypeExpDirect = 1 // 道具类型,1经验值
// PropTypeExpBall   = 2 // 道具类型,2经验球

)

// 等级配置结构
type LevelConfig struct {
	// 升到下一级所需经验
	NeedExperience int64
	// 智力
	Intellect int64
	// 体质
	Physique int64
}

// 等级配置， key是等级
var LevelConfigMap = map[int64]*LevelConfig{}

func initLevelConfig() {
	for i := 0; i <= 1000; i++ {
		LevelConfigMap[int64(i)] = &LevelConfig{
			NeedExperience: int64(i+1) * 100,
			Intellect:      int64(i%5+1) * 10,
			Physique:       int64(i%5+1) * 10,
		}
	}
}
