package model

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"reflect"

	"gorm.io/gorm"
)

// 任务类型，1每天任务，2每周任务， 3每月任务
type TaskType int64

const (
	EnumTaskTypeDay   TaskType = 1 // 每天任务
	EnumTaskTypeWeek  TaskType = 2 // 每周任务
	EnumTaskTypeMonth TaskType = 3 // 每月任务
)

type TaskEntity struct {
	Id       int64
	TaskType TaskType
	MaxCount int64
	StartTs  int64
	EndTs    int64
	Deadline int64
	Award    *TaskReward `gorm:"column:award;type:json;not null"`
}

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
	GetAward() *TaskReward

	Equal(other Tasker) bool
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

type Reward struct {
	// 获取 propId
	PropId int64
	// 获取数量
	Count int64
}

// GetCount implements RewardOner.
func (r *Reward) GetCount() int64 {
	return r.Count
}

// GetPropId implements RewardOner.
func (r *Reward) GetPropId() int64 {
	return r.PropId
}

// 任务配置的发放奖励
type TaskReward struct {
	Props []Reward
}

// 实现gorm json
func (t *TaskReward) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), t)
}
func (t TaskReward) Value() (driver.Value, error) {
	return json.Marshal(t)
}

// GetProps implements TaskRewarder.
func (t *TaskReward) GetProps() []Reward {
	return t.Props
}

func (t *TaskEntity) GetTaskId() int64 {
	return t.Id
}

func (t *TaskEntity) GetTaskType() TaskType {
	return t.TaskType
}

func (t *TaskEntity) GetMaxCount() int64 {
	return t.MaxCount
}

func (t *TaskEntity) GetStartTs() int64 {
	return t.StartTs
}

func (t *TaskEntity) GetEndTs() int64 {
	return t.EndTs
}

func (t *TaskEntity) GetDeadline() int64 {
	return t.Deadline
}

func (t *TaskEntity) GetAward() *TaskReward {
	return t.Award
}

func (t *TaskEntity) Equal(other Tasker) bool {
	// 比较全部字段
	if t.Id != other.GetTaskId() {
		return false
	}
	if t.TaskType != other.GetTaskType() {
		return false
	}
	if t.MaxCount != other.GetMaxCount() {
		return false
	}
	if t.StartTs != other.GetStartTs() {
		return false
	}
	if t.EndTs != other.GetEndTs() {
		return false
	}
	if t.Deadline != other.GetDeadline() {
		return false
	}
	if !reflect.DeepEqual(t.Award, other.GetAward()) {
		return false
	}

	return true
}
