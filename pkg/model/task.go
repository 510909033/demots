package model

import (
	"context"
)

// // 任务类型
type TaskType int64

// type Task struct {
// }

// type TaskRepository interface {
// 	GetById(ctx context.Context, tx *gorm.DB, taskId int64) (*Task, error)
// }

type EntityUseCase interface {
	// 创建一个用户
	CreateUser(ctx context.Context, user UserRepository) (UserRepository, error)
	// 登录并获取用户所有状态数据
	Login(ctx context.Context, param *LoginReq) (LoginResp, error)
	// 完成一个任务
	CompleteTask(ctx context.Context, user UserRepository, task Task) (*CompleteTaskResp, error)
	//
	InitSignConfig(ctx context.Context)

	Debug(ctx context.Context)
}

type LoginReq struct{}
type LoginResp struct {
	CommonBigLevelResp    CommonBigLevelResp    `json:"common_big_level_resp"`
	CommonSecondLevelResp CommonSecondLevelResp `json:"common_second_level_resp"`
}

type CompleteTaskResp struct{}

type CommonBigLevelResp struct {
	NewLevelList []Level `json:"new_level_list"`
}

type CommonSecondLevelResp struct {
	NewLevelList []SecondLevel `json:"new_level_list"`
}
