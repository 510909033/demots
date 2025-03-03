package repo

import (
	"context"

	"api/pkg/model"
	"api/pkg/utils/fn"

	"gorm.io/gorm"
)

type userTaskRepo struct {
	db *gorm.DB
}

func NewUserTaskRepository() model.UserTaskRepository {
	return &userTaskRepo{}
}

func (r *userTaskRepo) Create(ctx context.Context, tx *gorm.DB, userTask model.UserTasker) {
	err := tx.WithContext(ctx).Create(&userTask).Error
	fn.PanicErr(err)
}

func (r *userTaskRepo) GetDoingUserTask(ctx context.Context, tx *gorm.DB, user *model.UserEntity) []model.UserTasker {
	var userTasks []model.UserTaskEntity
	err := tx.WithContext(ctx).Where("user_id = ? AND status = ?", user.Id, model.EnumUserTaskStatusDoing).Find(&userTasks).Error
	fn.PanicErr(err)

	var result []model.UserTasker
	for _, userTask := range userTasks {
		result = append(result, &userTask)
	}
	return result
}

func (r *userTaskRepo) Get(ctx context.Context, tx *gorm.DB, userTaskId int64) model.UserTasker {
	var userTask model.UserTaskEntity
	err := tx.First(&userTask, userTaskId).Error
	fn.PanicErr(err)
	return &userTask
}

func (r *userTaskRepo) Save(ctx context.Context, tx *gorm.DB, userTask model.UserTasker) {
	err := tx.Save(&userTask).Error
	fn.PanicErr(err)
}

func (r *userTaskRepo) GetUserTaskCount(ctx context.Context, tx *gorm.DB, user *model.UserEntity, taskId int64, startTs, endTs int64) int64 {
	var count int64
	err := tx.Model(&model.UserTaskEntity{}).
		Where("user_id = ? AND task_id = ? AND create_ts >= ? AND create_ts < ?", user.Id, taskId, startTs, endTs).Count(&count).
		Error
	fn.PanicErr(err)
	return count
}
