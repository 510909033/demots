package repo

import (
	"api/pkg/model"
	"api/pkg/utils/fn"
	"context"

	"gorm.io/gorm"
)

type taskRepo struct {
}

func NewTaskRepo(db *gorm.DB) model.TaskRepository {
	return &taskRepo{}
}

func (r *taskRepo) Create(ctx context.Context, tx *gorm.DB, task model.Tasker) {
	err := tx.WithContext(ctx).Create(&task).Error
	fn.PanicErr(err)
}

func (r *taskRepo) Get(ctx context.Context, tx *gorm.DB, taskId int64) model.Tasker {
	var task model.TaskEntity
	err := tx.WithContext(ctx).Where("id = ?", taskId).First(&task).Error
	fn.PanicErr(err)
	return &task
}
