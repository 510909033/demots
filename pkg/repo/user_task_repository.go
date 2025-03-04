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

func (r *userTaskRepo) Create(ctx context.Context, tx *gorm.DB, userTask *model.UserTaskEntity) {
	err := tx.WithContext(ctx).Create(userTask).Error
	fn.PanicErr(err)
}

func (r *userTaskRepo) GetDoingUserTask(ctx context.Context, tx *gorm.DB, user *model.UserEntity) []*model.UserTaskEntity {
	var userTasks []*model.UserTaskEntity
	err := tx.WithContext(ctx).Where("user_id = ? AND status = ?", user.Id, model.EnumUserTaskStatusDoing).Find(&userTasks).Error
	fn.PanicErr(err)

	return userTasks
}

func (r *userTaskRepo) Get(ctx context.Context, tx *gorm.DB, userTaskId int64) *model.UserTaskEntity {
	var userTask model.UserTaskEntity
	err := tx.First(&userTask, userTaskId).Error
	fn.PanicErr(err)
	return &userTask
}

func (r *userTaskRepo) Save(ctx context.Context, tx *gorm.DB, userTask *model.UserTaskEntity) {
	err := tx.Save(userTask).Error
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

// DebugGetUserTaskList implements model.UserTaskRepository.
func (r *userTaskRepo) DebugGetUserTaskList(ctx context.Context, tx *gorm.DB, user *model.UserEntity) []*model.UserTaskEntity {
	var userTasks []*model.UserTaskEntity
	err := tx.WithContext(ctx).Where("user_id = ? ", user.Id).Find(&userTasks).Error
	fn.PanicErr(err)

	return userTasks
}
