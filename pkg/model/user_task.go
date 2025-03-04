package model

import (
	"context"

	"gorm.io/gorm"
)

// 当前用户任务状态, 0未开始，1进行中，2已完成，3已领取奖励，4已放弃
type UserTaskStatus int64

const (
	EnumUserTaskStatusUnStart  UserTaskStatus = 0 // 未开始
	EnumUserTaskStatusDoing    UserTaskStatus = 1 // 进行中
	EnumUserTaskStatusComplete UserTaskStatus = 2 // 已完成
	EnumUserTaskStatusReceive  UserTaskStatus = 3 // 已领取奖励
	EnumUserTaskStatusGiveUp   UserTaskStatus = 4 // 已放弃
)

type UserTaskUseCase interface {
	// 创建用户任务, 返回创建后的用户任务信息
	CreateUserTask(ctx context.Context, user *UserEntity, task Tasker) *UserTaskEntity
	// 获取进行中用户任务
	GetDoingUserTask(ctx context.Context, user *UserEntity) []*UserTaskEntity
	// 放弃任务
	GiveUpUserTask(ctx context.Context, userTask *UserTaskEntity)
	//完成任务
	CompleteUserTask(ctx context.Context, userTask *UserTaskEntity)
}

type UserTaskRepository interface {
	// 创建用户任务
	Create(ctx context.Context, tx *gorm.DB, userTask *UserTaskEntity)
	// 获取进行中用户任务
	GetDoingUserTask(ctx context.Context, tx *gorm.DB, user *UserEntity) []*UserTaskEntity
	// 根据用户任务id获取信息
	Get(ctx context.Context, tx *gorm.DB, userTaskId int64) *UserTaskEntity
	Save(ctx context.Context, tx *gorm.DB, userTask *UserTaskEntity)
	// 获取用户某段时间内已领取的任务数 [左闭右开]
	GetUserTaskCount(ctx context.Context, tx *gorm.DB, user *UserEntity, taskId int64, startTs, endTs int64) int64
	// 【调试】获取用户的全部任务列表
	DebugGetUserTaskList(ctx context.Context, tx *gorm.DB, user *UserEntity) []*UserTaskEntity
}

func NewUserTaskEntity(userId, taskId int64) *UserTaskEntity {
	return &UserTaskEntity{
		UserId: userId,
		TaskId: taskId,
	}
}

type UserTaskEntity struct {
	Id         int64
	UserId     int64
	TaskId     int64
	CreateTs   int64
	CompleteTs int64
	Status     UserTaskStatus
	GiveUpTs   int64
	ReceiveTs  int64
	StartTs    int64
	EndTs      int64
}

// UserTaskStatus
func (u UserTaskStatus) Summary() string {
	switch u {
	case EnumUserTaskStatusUnStart:
		return "未开始"
	case EnumUserTaskStatusDoing:
		return "进行中"
	case EnumUserTaskStatusComplete:
		return "已完成"
	case EnumUserTaskStatusReceive:
		return "已领取"
	case EnumUserTaskStatusGiveUp:
		return "已放弃"
	default:
		return "未知"
	}
}
