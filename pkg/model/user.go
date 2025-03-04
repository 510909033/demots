package model

import (
	"context"
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type NewUserIder interface {
	GetUserId() int64
}

// 用户
type UserEntity struct {
	// 用户id
	Id int64
	// 智力
	Intellect int64
	// 未分配的智力
	UnallocatedIntellect int64
	// 体质
	Physique int64
	// 未分配的体质
	UnallocatedPhysique int64
	// 耐力
	Endurance int64
	// 未分配的耐力
	UnallocatedEndurance int64
	// 经验
	Experience int64
	// 头像
	Avatar string
	// 姓名
	Name string
	// 年龄
	Age int64
	// 等级
	Level int64

	// 用户装备
	UserGear TUserGear `gorm:"column:user_gear;type:json;not null"`
	// 用户背包
	UserBag TUserBag `gorm:"column:user_bag;type:json;not null"`
}

// GetId implements *UserEntity.
func (u *UserEntity) GetId(ctx context.Context) int64 {
	return u.Id
}

// GetUserId implements *UserEntity.
func (u *UserEntity) GetUserId() int64 {
	return u.Id
}

// type TUserGear map[int64]*GearEntity // key 是position
type TUserGear map[int64]*UserGearEntity

func (u *UserEntity) GetUserGear(pos int64) *UserGearEntity {
	if _, ok := u.UserGear[pos]; !ok {
		return &UserGearEntity{
			PropId:   0,
			Position: pos,
		}
	}
	return u.UserGear[pos]
}

// key 是propId
type TUserBag map[int64]*UserBagEntity

// // * **境界等级:**  炼气、筑基、金丹、元婴、化神、炼虚、合体、大乘、渡劫。
type Level int64
type SecondLevel int64

// type UserRepository interface {
// 	// 通过id获取用户
// 	GetById(ctx context.Context, tx *gorm.DB, id int64) (*User, error)
// }

// type UserUseCase interface {
// 	// 用户完成一个任务
// 	CompleteTask(ctx context.Context, userId, taskId int64) error
// }

// 用户
type UserRepository interface {
	// 创建用户
	Create(ctx context.Context, tx *gorm.DB, user *UserEntity)
	// 获取用户id
	Get(ctx context.Context, tx *gorm.DB, userId int64) *UserEntity
	// 增加经验
	AddExperience(ctx context.Context, tx *gorm.DB, user *UserEntity, experience int64)
	// 增加等级
	AddLevel(ctx context.Context, tx *gorm.DB, user *UserEntity, level int64)
	// 增加未分配属性
	AddUnallocatedAttribute(ctx context.Context, tx *gorm.DB, user *UserEntity, cfg *LevelConfig)
	// 增加分配属性
	AddAttribute(ctx context.Context, tx *gorm.DB, user *UserEntity, param *AddAttributeReq)
	// 设置用户某个位置的装备
	// SetUserGear(ctx context.Context, tx *gorm.DB, user *UserEntity, position int64, gearEntity *GearEntity)
	SetUserGear(ctx context.Context, tx *gorm.DB, user *UserEntity, position int64, gearId int64, check func(ctx context.Context, user *UserEntity, param *UserChangeGearReq) bool)
}

type AddAttributeReq struct {
	// 智力
	Intellect int64
	// 体质
	Physique int64
	//  耐力
	Endurance int64
}

func (t TUserGear) Value() (driver.Value, error) {
	return json.Marshal(t)
}
func (t *TUserGear) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), t)
}
func (t TUserBag) Value() (driver.Value, error) {
	return json.Marshal(t)
}
func (t *TUserBag) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), t)
}

// 拷贝，值拷贝整数、字符串、bool 简单数据结构
func (u *UserEntity) CopyEntity() *UserEntity {
	return &UserEntity{
		Id:                   u.Id,
		Intellect:            u.Intellect,
		UnallocatedIntellect: u.UnallocatedIntellect,
		Physique:             u.Physique,
		UnallocatedPhysique:  u.UnallocatedPhysique,
		Endurance:            u.Endurance,
		UnallocatedEndurance: u.UnallocatedEndurance,
		Experience:           u.Experience,
		Avatar:               u.Avatar,
		Name:                 u.Name,
		Age:                  u.Age,
		Level:                u.Level,
		UserGear:             map[int64]*UserGearEntity{},
		UserBag:              map[int64]*UserBagEntity{},
	}
}
