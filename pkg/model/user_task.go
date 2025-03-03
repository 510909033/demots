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

type UserTasker interface {
	// 获取用户id
	GetUserId() int64
	// 获取任务id
	GetTaskId() int64
	// 获取领取任务时间戳
	GetCreateTs() int64
	// 获取完成任务时间戳
	GetCompleteTs() int64
	// 获取当前任务状态
	GetStatus() UserTaskStatus
	// 获取放弃任务时的时间戳
	GetGiveUpTs() int64
	// 获取领取奖励时的时间戳
	GetReceiveTs() int64
	// 任务开始时间戳【包含】
	GetStartTs() int64
	// 任务结束时间戳【不包含】
	GetEndTs() int64

	SetUserId(userId int64)
	SetTaskId(taskId int64)
	SetCreateTs(createTs int64)
	// 设置完成时间
	SetCompleteTs(completeTs int64)
	// 设置放弃时间
	SetGiveUpTs(giveUpTs int64)
	// 设置领取奖励时间
	SetReceiveTs(receiveTs int64)
	// 设置任务状态
	SetStatus(status UserTaskStatus)
	// 设置任务开始时间
	SetStartTs(startTs int64)
	// 设置任务结束时间, 0 表示不限制
	SetEndTs(endTs int64)
}

type UserTaskUseCase interface {
	// 创建用户任务, 返回创建后的用户任务信息
	CreateUserTask(ctx context.Context, user *UserEntity, task Tasker) UserTasker
	// 获取进行中用户任务
	GetDoingUserTask(ctx context.Context, user *UserEntity) []UserTasker
	// 放弃任务
	GiveUpUserTask(ctx context.Context, userTask UserTasker)
	//完成任务
	CompleteUserTask(ctx context.Context, userTask UserTasker)
}

type UserTaskRepository interface {
	// 创建用户任务
	Create(ctx context.Context, tx *gorm.DB, userTask UserTasker)
	// 获取进行中用户任务
	GetDoingUserTask(ctx context.Context, tx *gorm.DB, user *UserEntity) []UserTasker
	// 根据用户任务id获取信息
	Get(ctx context.Context, tx *gorm.DB, userTaskId int64) UserTasker
	Save(ctx context.Context, tx *gorm.DB, userTask UserTasker)
	// 获取用户某段时间内已领取的任务数 [左闭右开]
	GetUserTaskCount(ctx context.Context, tx *gorm.DB, user *UserEntity, taskId int64, startTs, endTs int64) int64
}

func NewUserTasker(userId, taskId int64) UserTasker {
	return (UserTasker)(nil)
}

type UserTaskEntity struct {
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

func (u *UserTaskEntity) GetUserId() int64 {
	return u.UserId
}

func (u *UserTaskEntity) GetTaskId() int64 {
	return u.TaskId
}

func (u *UserTaskEntity) GetCreateTs() int64 {
	return u.CreateTs
}

func (u *UserTaskEntity) GetCompleteTs() int64 {
	return u.CompleteTs
}

func (u *UserTaskEntity) GetStatus() UserTaskStatus {
	return u.Status
}

func (u *UserTaskEntity) GetGiveUpTs() int64 {
	return u.GiveUpTs
}

func (u *UserTaskEntity) GetReceiveTs() int64 {
	return u.ReceiveTs
}

func (u *UserTaskEntity) GetStartTs() int64 {
	return u.StartTs
}

func (u *UserTaskEntity) GetEndTs() int64 {
	return u.EndTs
}

func (u *UserTaskEntity) SetUserId(userId int64) {
	u.UserId = userId
}

func (u *UserTaskEntity) SetTaskId(taskId int64) {
	u.TaskId = taskId
}

func (u *UserTaskEntity) SetCreateTs(createTs int64) {
	u.CreateTs = createTs
}

func (u *UserTaskEntity) SetCompleteTs(completeTs int64) {
	u.CompleteTs = completeTs
}

func (u *UserTaskEntity) SetGiveUpTs(giveUpTs int64) {
	u.GiveUpTs = giveUpTs
}

func (u *UserTaskEntity) SetReceiveTs(receiveTs int64) {
	u.ReceiveTs = receiveTs
}

func (u *UserTaskEntity) SetStatus(status UserTaskStatus) {
	u.Status = status
}

func (u *UserTaskEntity) SetStartTs(startTs int64) {
	u.StartTs = startTs
}

func (u *UserTaskEntity) SetEndTs(endTs int64) {
	u.EndTs = endTs
}
