package model

import (
	"api/pkg/utils/fn"
	"context"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

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

func NewProper(t PropType, data any) *PropEntity {
	return &PropEntity{
		Type:   t,
		Entity: data,
	}
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

// NewPropId implements Proper.
func (e *PropEntity) NewPropId() int64 {
	return e.Id
}

type PropRepository interface {
	Create(ctx context.Context, tx *gorm.DB, prop *PropEntity)
	Get(ctx context.Context, tx *gorm.DB, propId int64) *PropEntity

	// 获取一个随机道具
	GetRandomProp(ctx context.Context, tx *gorm.DB) *PropEntity
}

func (e *PropEntity) ConvertExperiencer() *ExperienceEntity {
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
	SendReward(ctx context.Context, tx *gorm.DB, user *UserEntity, now int64, taskRewarder *TaskReward)
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

func (e *PropEntity) SetContent(content string) {
	e.Content = content
}

func (e *PropEntity) SetEndTs(endTs int64) {
	e.EndTs = endTs
}
func (e *PropEntity) SetStartTs(startTs int64) {
	e.StartTs = startTs
}

// GetContent implements  *PropEntity.
func (e *PropEntity) GetContent() string {
	return e.Content
}

// GetEndTs implements  *PropEntity.
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
func (e *PropEntity) GetEntity() any {
	return e.Entity
}
func (e *PropEntity) Equal(other *PropEntity) bool {
	// 比较所有字段
	if e.Id != other.NewPropId() {
		return false
	}
	if e.Type != other.GetType() {
		return false
	}
	if e.Content != other.GetContent() {
		return false
	}
	if e.StartTs != other.GetStartTs() {
		return false
	}
	if e.EndTs != other.GetEndTs() {
		return false
	}
	return true
}
