package model

import "errors"

var (
	ErrLoginUserNotExist = errors.New("用户不存在")
	// 经验值不足
	ErrExperienceNotEnough = errors.New("经验值不足")
	// 用户不满足当前装备的穿戴条件
	ErrUserNotSatisfyGear           = errors.New("用户不满足当前装备的穿戴条件")
	ErrUserTaskStatusNotDoing       = errors.New("用户任务状态不是进行中")
	ErrUserTaskTimeNotStart         = errors.New("用户任务时间未开始")
	ErrUserTaskTimeEnd              = errors.New("用户任务时间已结束")
	ErrTaskNotTimeStart             = errors.New("任务时间未开始")
	ErrTaskTimeEnd                  = errors.New("任务时间已结束")
	ErrTaskMaxCount                 = errors.New("任务次数已达上限")
	ErrTaskTypeNotSupport           = errors.New("任务类型不支持")
	ErrUserTaskStatusNotAllowGiveUp = errors.New("用户任务状态不允许放弃")
	ErrUserTaskUserIdNoPermission   = errors.New("用户无此任务权限")
	ErrUserTaskHadReceived          = errors.New("用户任务奖励已领取")
	ErrUserTaskNotAllowReceiveAward = errors.New("用户任务奖励不允许领取")
	ErrUserNameEmpty                = errors.New("用户名不能为空")
	// 服务异常，请稍后重试
	ErrServerBusy    = errors.New("服务异常，请稍后重试")
	ErrUserNameExist = errors.New("该昵称已存在")
	ErrPropType      = errors.New("道具类型错误")
)
