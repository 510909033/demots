package model

import (
	"context"

	"gorm.io/gorm"
)

type CommonUseCase interface {
	CompleteUserAction(ctx context.Context, tx *gorm.DB, user *UserEntity, param *UserActionReq) (*CompleteAddExperienceResp, error)
	UpdateUser(ctx context.Context, user *UserEntity)
	// 获取用户某天的每日任务列表数据
	GetUserDayTaskList(ctx context.Context, user *UserEntity, now int64) (*UserDayTaskListResp, error)
}

type UserDayTaskListResp struct {
	User *UserEntity `json:"user"`
}
