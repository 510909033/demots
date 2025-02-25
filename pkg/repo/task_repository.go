package repo

import (
	"api/pkg/model"
	"context"
	"database/sql/driver"
	"encoding/json"
)

func NewTask() model.Task {
	return &Task{}
}

type Task struct {
	Id       int64          `json:"id" gorm:"primaryKey"`
	TaskType model.TaskType `json:"task_type" gorm:"column:task_type;type:bigint;not null;default:0"`
	UserId   int64          `json:"user_id" gorm:"column:user_id;type:bigint;not null;default:0"`
	MinLevel int64          `json:"min_level" gorm:"column:min_level;type:bigint;not null;default:0"`
	MaxLevel int64          `json:"max_level" gorm:"column:max_level;type:bigint;not null;default:0"`
	CreateTs int64          `json:"create_ts" gorm:"column:create_ts;type:bigint;not null;default:0"`
	// json 格式
	Content TaskContentInterface `json:"content" gorm:"column:content;type:json;not null;default:''"`
}

// GetTaskId implements model.Task.
func (t *Task) GetTaskId(ctx context.Context) int64 {
	panic("unimplemented")
}

type TaskContentInterface interface {
	// 获取经验奖励数值
	GetExperienceReward() int64
}

var _ TaskContentInterface = (*TaskContentSignIn)(nil)

// 签到
type TaskContentSignIn struct {
	ExperienceReward int64 `json:"experience_reward"`
}

func (t TaskContentSignIn) GetExperienceReward() int64 {
	return t.ExperienceReward
}
func (t TaskContentSignIn) Value() (driver.Value, error) {
	return json.Marshal(t)
}
func (t *TaskContentSignIn) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), t)
}
