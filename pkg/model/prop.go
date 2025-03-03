package model

import (
	"api/pkg/utils/fn"
	"context"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

var _ Proper = &PropEntity{}

type NewPropId interface {
	NewPropId() int64
}

type PropType int64

const (
	// 1-装备
	PropTypeGear PropType = 1
	// 2-经验, 直接发放
	PropTypeExperience PropType = 2
)

func _() {
	// expi
}

type Proper interface {
	NewPropId
	GetPropId() int64
	GetType() PropType
	GetContent() string
	ConvertGearEntity() *GearEntity
	ConvertExperiencer() Experiencer
	// 有效期，开始时间戳， 0 表示立即生效， 左闭
	GetStartTs() int64
	// 有效期，结束时间戳，0 表示永久有效,  右开
	GetEndTs() int64
}

type PropEntity struct {
	// 道具id
	Id int64
	// 道具类型,
	Type    PropType
	StartTs int64
	EndTs   int64
	Content string

	Entity any `gorm:"-" json:"-"`
}

// ConvertGearEntity implements Proper.
func (e *PropEntity) ConvertGearEntity() *GearEntity {
	switch e.Type {
	case PropTypeGear:
		var entity = GearEntity{}
		err := json.Unmarshal([]byte(e.Content), &entity)
		fn.PanicErr(err)
		e.Entity = &entity
		return &entity
	}
	fn.PanicErr(fmt.Errorf("未知的类型 %d", e.Type))
	return nil
}

// GetContent implements Proper.
func (e *PropEntity) GetContent() string {
	return e.Content
}

// GetEndTs implements Proper.
func (e *PropEntity) GetEndTs() int64 {
	return e.EndTs
}

// GetPropId implements Proper.
func (e *PropEntity) GetPropId() int64 {
	return e.Id
}

// GetStartTs implements Proper.
func (e *PropEntity) GetStartTs() int64 {
	return e.StartTs
}

// GetType implements Proper.
func (e *PropEntity) GetType() PropType {
	return e.Type
}

// NewPropId implements Proper.
func (e *PropEntity) NewPropId() int64 {
	return e.Id
}

type PropRepository interface {
	Create(ctx context.Context, tx *gorm.DB, prop *PropEntity)
	Get(ctx context.Context, tx *gorm.DB, propId int64) Proper
}

func (e *PropEntity) ConvertExperiencer() Experiencer {
	switch e.Type {
	case PropTypeExperience:
		var entity = ExperienceEntity{}
		err := json.Unmarshal([]byte(e.Content), &entity)
		fn.PanicErr(err)
		e.Entity = &entity
		return &entity
	}
	fn.PanicErr(fmt.Errorf("未知的类型 %d", e.Type))
	return nil
}

// func (t PropContent) Value() (driver.Value, error) {
// 	return json.Marshal(t)
// }
// func (t *PropContent) Scan(input interface{}) error {
// 	// t.Data = string(input.([]byte))
// 	return json.Unmarshal(input.([]byte), t)
// 	return nil
// }

func (e *PropEntity) Convert() {
	switch e.Type {
	case PropTypeGear:
		e.ConvertGearEntity()
	case PropTypeExperience:
		e.ConvertExperiencer()
	default:
		fn.PanicErr(fmt.Errorf("未知的类型 %d", e.Type))
	}
}

type PropUseCase interface {
	SendReward(ctx context.Context, tx *gorm.DB, user *UserEntity, now int64, taskRewarder TaskRewarder)
}

// PropType
func (t PropType) String() string {
	switch t {
	case PropTypeGear:
		return "装备"
	case PropTypeExperience:
		return "经验"
	default:
		return "unknown"
	}
}
