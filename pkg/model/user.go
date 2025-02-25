package model

import (
	"context"

	"gorm.io/gorm"
)

// // 用户
// type User struct {
// 	// 用户id
// 	Id int64
// 	// 智力
// 	Intellect int64
// 	// 体质
// 	Physique int64
// 	// 耐力
// 	Endurance int64
// 	// 经验
// 	Experience int64
// 	// 头像
// 	Avatar string
// 	// 姓名
// 	Name string
// 	// 年龄
// 	Age int64
// }

// // * **境界等级:**  炼气、筑基、金丹、元婴、化神、炼虚、合体、大乘、渡劫。
type Level int64
type SecondLevel int64

// type UserRepository interface {
// 	// 通过id获取用户
// 	GetById(ctx context.Context, tx *gorm.DB, id int64) (*User, error)
// }

func demo(ctx context.Context) {
	db, _ := gorm.Open(nil, nil)
	db.Transaction(func(tx *gorm.DB) error {

		return nil
	})

}

// type UserUseCase interface {
// 	// 用户完成一个任务
// 	CompleteTask(ctx context.Context, userId, taskId int64) error
// }

// 用户
type UserRepository interface {
	// 获取用户id
	GetId(ctx context.Context) int64
	// 增加经验
	AddExperience(ctx context.Context, experience int64)

	// 设置用户名
	SetName(ctx context.Context, name string)

	Save(ctx context.Context, tx *gorm.DB) error
}

// 一个任务
type Task interface {
	GetTaskId(ctx context.Context) int64
}
