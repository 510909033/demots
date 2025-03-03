package model

import "context"

type GetTaskIder interface {
	GetTaskId() int64
}

type GetUserTaskIder interface {
	GetUserTaskId() int64
}

type EntityUseCase interface {
	// 创建一个用户
	CreateUser(ctx context.Context, user UserRepository) (UserRepository, error)
	// 登录并获取用户所有状态数据
	Login(ctx context.Context, param *LoginReq) (*LoginResp, error)
	// 断线重连
	Reconnect(ctx context.Context, param *LoginReq) (*LoginResp, error)
	// 用户领取一个任务, 并直接开始
	UserCreateTask(ctx context.Context, user *UserEntity, taskId GetTaskIder) (*UserCreateTaskResp, error)
	// 完成一个任务
	CompleteTask(ctx context.Context, user *UserEntity, userTaskId GetUserTaskIder) (*CompleteTaskResp, error)
	// 用户放弃一个任务
	UserGiveUpUserTask(ctx context.Context, user *UserEntity, userTaskId GetUserTaskIder) error
	// 用户领取任务奖励， 同时设置任务已领取
	UserReceiveUserTaskReward(ctx context.Context, user *UserEntity, userTaskId GetUserTaskIder) (*UserReceiveUserTaskResp, error)

	//
	InitSignConfig(ctx context.Context)
	// 初始化装备数据
	InitGearConfig(ctx context.Context)

	Debug(ctx context.Context)

	// 直接增加经验值
	CompleteUserAction(ctx context.Context, user *UserEntity, param *UserActionReq) (*CompleteAddExperienceResp, error)
	// 更换装备
	CompleteUserChangeGear(ctx context.Context, user *UserEntity, param *UserChangeGearReq) (*UserChangeGearResp, error)
	// 检查用户是否可以更换某个装备
	CheckUserCanChangeGear(ctx context.Context, user *UserEntity, param *UserChangeGearReq) bool
}

type LoginReq struct {
	Id int64
}
type LoginResp struct {
	User *UserEntity
}

type CompleteTaskResp struct{}

type CommonBigLevelResp struct {
	NewLevelList []Level `json:"new_level_list"`
}

type CommonSecondLevelResp struct {
	NewLevelList []SecondLevel `json:"new_level_list"`
}

type CompleteAddExperienceResp struct {
	User *UserEntity  `json:"user"`
	Cfg  *LevelConfig `json:"cfg"`
	// 提升的等级列表
	LevelList []int64 `json:"level_list"`
	// 剩余经验值
	LeaveExperience int64 `json:"leave_experience"`
}

type UserActionReq struct {
	// 只增加经验，负值为减少，0不操作
	AddExperience int64 `json:"add_experience"`
	// 等级 + 1 同时减少经验
	AddLevel_1 bool `json:"add_level"`
	// 将全部经验值转换为提升等级, 可能会提升多次等级
	ConvertAllExperienceToLevel bool `json:"convert_all_experience_to_level"`
	// 分配属性（未分配的属性点需要充足）
	AddAttribute *AddAttributeReq `json:"add_attribute"`
	// 增加未分配属性点（如果是减少，分配的属性点需要充足）
	AddUnallocatedAttribute *AddAttributeReq `json:"add_unallocated_attribute"`

	// 用户背包增加装备
	AddUserBagGear AddUserBagGear `json:"add_user_bag_gear"`
}

type UserChangeGearReq struct {
	PropId   int64 `json:"gear_id"`
	Position int64 `json:"position"`
}

type UserChangeGearResp struct {
	User *UserEntity `json:"user"`
}

type AddUserBagGear struct {
	// 增加的装备列表
	Gears []*GearEntity `json:"gears"`
}

type UserCreateTaskResp struct {
}

type UserReceiveUserTaskResp struct {
}
