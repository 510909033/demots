package model

import (
	"context"

	"gorm.io/gorm"
)

// 任务类型，1每天任务，2每周任务， 3每月任务
type TaskType int64

const (
	EnumTaskTypeDay   TaskType = 1 // 每天任务
	EnumTaskTypeWeek  TaskType = 2 // 每周任务
	EnumTaskTypeMonth TaskType = 3 // 每月任务
)

// type Task struct {
// }

// type TaskRepository interface {
// 	GetById(ctx context.Context, tx *gorm.DB, taskId int64) (*Task, error)
// }

// 一个任务的接口定义
// type Tasker interface {
// 	// 获取任务id
// 	GetTaskId(ctx context.Context) int64
// 	// 任务类型，1每天任务，2每周任务， 3每月任务
// 	GetTaskType(ctx context.Context) int64
// 	// 获取当前任务类型，任务周期内最大可完成任务数
// 	GetMaxCount(ctx context.Context) int64
// 	// 检查给定的时间戳是否在任务有效期内
// 	CheckTsInValid(ctx context.Context, ts int64) bool
// }

type Tasker interface {
	// 获取任务id
	GetTaskId() int64
	// 任务类型，1每天任务，2每周任务， 3每月任务
	GetTaskType() TaskType
	// 获取当前任务类型，任务周期内最大可完成任务数, 0 表示不限制
	GetMaxCount() int64
	// 任务开始时间戳【包含】
	GetStartTs() int64
	// 任务结束时间戳【不包含】
	GetEndTs() int64
	// 任务期限，单位秒，0表示不限
	GetDeadline() int64
	// 任务配置的发放奖励
	GetAward() TaskRewarder
}

type TaskUseCase interface {
	// 创建一个任务
	CreateTask(ctx context.Context, task Tasker)
	// 获取任务详情
	GetTask(ctx context.Context, taskId int64) Tasker
}

type TaskRepository interface {
	Create(ctx context.Context, tx *gorm.DB, task Tasker)
	Get(ctx context.Context, tx *gorm.DB, taskId int64) Tasker
}

// 任务配置的发放奖励
type TaskRewarder interface {
	GetProps() []RewardOner
}

type RewardOner interface {
	// 获取 propId
	GetPropId() int64
	// 获取数量
	GetCount() int64
}
