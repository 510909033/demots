package model

type AlchemyType int64

const (
	// 1 修为
	AlchemyType_XiuWei AlchemyType = 1
)

// 阿尔克米
type AlchemyEntity struct {
	// 名称
	Name string `json:"name"`
	// 描述
	Desc string `json:"desc"`
	// 类型、功效
	Type   AlchemyType `json:"type"`
	Reward Reward      `json:"reward"`
}
