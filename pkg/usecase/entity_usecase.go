package usecase

import (
	"api/pkg/model"
	"api/pkg/repo"
	"api/pkg/utils/fn"
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _log, _ = zap.NewProduction()
var log = _log.Sugar()
var db *gorm.DB
var lock sync.Mutex

// usecase
var (
	userRepo     model.UserRepository     = repo.NewUserRepository()
	taskRepo     model.TaskRepository     = repo.NewTaskRepository()
	userTaskRepo model.UserTaskRepository = repo.NewUserTaskRepository()
	propRepo     model.PropRepository     = repo.NewPropRepository()
	useBagRepo   model.UserBagRepository  = repo.NewUserBagRepository()

	// taskCase     model.TaskUseCase     = NewTaskUseCase()
	// userTaskCase model.UserTaskUseCase = NewUserTaskUseCase()
	propCase   model.PropUseCase   = NewPropUseCase()
	commonCase model.CommonUseCase = NewCommonUseCase()
	debug                          = NewDebug()
)

// 所有道具
var propAll = make(map[int64]*model.PropEntity, 2000)

// 任务数据
var taskAll = make(map[int64]*model.TaskEntity, 2000)

var UserMap = make(map[int64]*model.UserEntity)
var UserLock = make(map[int64]*sync.Mutex)

func init() {
	repo.SetLog(log)
}

func GetTraceid(ctx context.Context) string {
	return uuid.NewString()[:8]
}

func GetJsonString(val any) string {
	v, _ := json.Marshal(val)
	return string(v)
}

func NewEntity() model.EntityUseCase {
	var err error
	// db, err = gorm.Open(sqlite.Open("/tmp/test.db"), &gorm.Config{})

	dsn := "admin:admin@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db = db.Debug()

	entityUseCase := &Entity{
		gear: repo.NewGearRepository(),
		// prop: repo.NewPropRepository(),
	}

	rootCtx := context.Background()

	// e.InitGearAllToMemory(rootCtx)
	entityUseCase.InitPropAllToMemory(rootCtx)
	entityUseCase.InitTaskAllToMemory(rootCtx)

	return entityUseCase
}

type Entity struct {
	gear model.GearRepository
	prop model.PropRepository
}

// UserReceiveUserTaskReward implements model.EntityUseCase.
func (e *Entity) UserReceiveUserTaskReward(ctx context.Context, user *model.UserEntity, userTaskId model.GetUserTaskIder) (*model.UserReceiveUserTaskResp, error) {
	var resp = &model.UserReceiveUserTaskResp{}
	now := time.Now().Unix()
	err := db.Transaction(func(tx *gorm.DB) error {
		// 获取用户任务信息。
		userTaskInfo := userTaskRepo.Get(ctx, tx, userTaskId.GetUserTaskId())

		if userTaskInfo.UserId != user.GetId(ctx) {
			return model.ErrUserTaskUserIdNoPermission
		}

		taskInfo := taskRepo.Get(ctx, tx, userTaskInfo.UserId)

		switch userTaskInfo.Status {
		case model.EnumUserTaskStatusReceive:
			// 已领取
			return model.ErrUserTaskHadReceived
		case model.EnumUserTaskStatusComplete:
		default:
			// 当前任务不能领取奖励
			return model.ErrUserTaskNotAllowReceiveAward
		}

		userTaskInfo.ReceiveTs = (now)
		userTaskInfo.Status = (model.EnumUserTaskStatusReceive)
		userTaskRepo.Save(ctx, tx, userTaskInfo)

		// 发放奖励
		propCase.SendReward(ctx, tx, user, now, taskInfo.GetAward())

		return nil
	})

	return resp, err
}

// UserGiveUpUserTask implements model.EntityUseCase.
func (e *Entity) UserGiveUpUserTask(ctx context.Context, user *model.UserEntity, userTaskId model.GetUserTaskIder) error {
	now := time.Now().Unix()
	err := db.Transaction(func(tx *gorm.DB) error {
		// 获取用户任务信息。
		userTaskInfo := userTaskRepo.Get(ctx, tx, userTaskId.GetUserTaskId())

		if userTaskInfo.UserId != user.GetId(ctx) {
			return model.ErrUserTaskUserIdNoPermission
		}

		// taskInfo := taskRepo.Get(ctx, tx, userTaskInfo.GetTaskId())
		// _=taskInfo

		switch userTaskInfo.Status {
		case model.EnumUserTaskStatusGiveUp:
			return nil
		case model.EnumUserTaskStatusDoing, model.EnumUserTaskStatusUnStart, model.EnumUserTaskStatusComplete:
		default:
			// 当前任务不能放弃
			return model.ErrUserTaskStatusNotAllowGiveUp

		}

		userTaskInfo.GiveUpTs = (now)
		userTaskInfo.Status = (model.EnumUserTaskStatusGiveUp)

		userTaskRepo.Save(ctx, tx, userTaskInfo)

		return nil
	})

	return err
}

// UserCreateTask implements model.EntityUseCase.
func (e *Entity) UserCreateTask(ctx context.Context, user *model.UserEntity, taskId model.GetTaskIder) (*model.UserCreateTaskResp, error) {
	var resp = &model.UserCreateTaskResp{}
	now := time.Now().Unix()

	err := db.Transaction(func(tx *gorm.DB) error {
		taskInfo := taskRepo.Get(ctx, tx, taskId.GetTaskId())

		if taskInfo.GetStartTs() > 0 && taskInfo.GetStartTs() < now {
			return model.ErrTaskNotTimeStart
		}
		if taskInfo.GetEndTs() > 0 && taskInfo.GetEndTs() > now {
			return model.ErrTaskTimeEnd
		}

		if taskInfo.GetMaxCount() != 0 {
			maxCount := taskInfo.GetMaxCount()
			var startTs int64
			var endTs int64
			switch taskInfo.GetTaskType() {
			case model.EnumTaskTypeDay:
				startTs = GetTodayStartTs()
				endTs = GetTodayEndTs() + 1
			case model.EnumTaskTypeWeek:
				startTs = GetWeekStartTs()
				endTs = GetWeekEndTs() + 1
			case model.EnumTaskTypeMonth:
				startTs = GetMonthStartTs()
				endTs = GetMonthEndTs() + 1
			default:
				return model.ErrTaskTypeNotSupport
			}

			if userTaskRepo.GetUserTaskCount(ctx, tx, user, taskId.GetTaskId(), startTs, endTs) >= maxCount {
				return model.ErrTaskMaxCount
			}
		}

		// userTaskInfo := model.NewUserTaskEntity(user.GetId(ctx), taskId.GetTaskId())
		// userTaskInfo.SetUserId(user.GetId(ctx))
		// userTaskInfo.SetTaskId(taskId.GetTaskId())
		// userTaskInfo.SetCreateTs(now)
		// userTaskInfo.SetCompleteTs(0)
		// userTaskInfo.SetGiveUpTs(0)
		// userTaskInfo.SetReceiveTs(0)
		// userTaskInfo.SetStatus(model.EnumUserTaskStatusDoing)
		// userTaskInfo.SetStartTs(now)
		userTaskInfo := &model.UserTaskEntity{
			Id:         0,
			UserId:     user.Id,
			TaskId:     taskId.GetTaskId(),
			CreateTs:   now,
			CompleteTs: 0,
			Status:     model.EnumUserTaskStatusDoing,
			GiveUpTs:   0,
			ReceiveTs:  0,
			StartTs:    now,
			EndTs:      0,
		}

		endTs := int64(0)
		if taskInfo.GetDeadline() == 0 {
			endTs = 0
		} else {
			endTs = (now) + taskInfo.GetDeadline()
		}
		userTaskInfo.EndTs = (endTs)

		// 保存更新后的用户任务信息。
		userTaskRepo.Create(ctx, tx, userTaskInfo)

		// 如果没有遇到任何错误，返回nil。
		return nil
	})

	if err != nil {
		return nil, err
	}

	// 返回完成任务的响应结构体和nil错误，表示任务完成成功。
	return resp, nil
}

// CheckUserCanChangeGear implements model.EntityUseCase.
func (e *Entity) CheckUserCanChangeGear(ctx context.Context, user *model.UserEntity, param *model.UserChangeGearReq) bool {

	gear := propAll[param.PropId].Entity.(*model.GearEntity)
	// 等级
	if gear.MinLevel > 0 {
		if user.Level < gear.MinLevel {
			return false
		}
	}

	// 位置
	if gear.Type != param.Position {
		switch param.Position {
		case model.EnumUserGearPositionArm_3_2:
			if gear.Type == model.EnumGearTypeArm_3 {
				return true
			}
		case model.EnumUserGearPositionRing_8_2:
			if gear.Type == model.EnumGearTypeRing_8 {
				return true
			}
		}
		return false
	}

	return true
}

func (e *Entity) UpdateUser(ctx context.Context, user *model.UserEntity) {
	commonCase.UpdateUser(ctx, user)
}

// Debug implements model.EntityUseCase.
func (e *Entity) Debug(ctx context.Context) {
	// db.AutoMigrate(&model.PropEntity{})
	// db.AutoMigrate(&model.UserEntity{})
	// db.AutoMigrate(&model.GearEntity{})
	// db.AutoMigrate(&model.UserGearEntity{})
	db.AutoMigrate(&model.TaskEntity{})
	db.AutoMigrate(&model.UserTaskEntity{})
	// db.AutoMigrate(&model.UserBagEntity{})

	// 登录
	e.Login(ctx, &model.LoginReq{
		Id: 3,
	})

	action := "create_user"
	// action = "prop_test"
	// action = "CompleteAddExperience"
	// action = "SetUserGear"
	// action = "InitGearConfig"
	action = "mock"
	switch action {
	case "create_user":
		user := &model.UserEntity{
			Id:         0,
			Intellect:  0,
			Physique:   0,
			Endurance:  0,
			Experience: 20,
			Avatar:     "11",
			Name:       "user1",
			Age:        10,
		}
		userRepo.Create(ctx, db, user)
		log.Info(GetJsonString(user))

		userRepo.AddExperience(ctx, db, user, 20)
		log.Info(GetJsonString(user))
		user = userRepo.Get(ctx, db, user.Id)
		log.Info(GetJsonString(user))

	case "prop_test":
		// var data = &model.PropEntity{
		// 	Id: 0,
		// 	// PropType: model.PropTypeExpDirect,
		// }
		// repo.NewPropRepository().Create(ctx, db, data)
		// log.Info(GetJsonString(data))

		// data2 := repo.NewPropRepository().Get(ctx, db, 51)
		// log.Info(GetJsonString(data2))
	case "CompleteAddExperience":
		resp, err := e.CompleteUserAction(ctx, UserMap[3], &model.UserActionReq{
			AddExperience:               2000,
			AddLevel_1:                  true,
			ConvertAllExperienceToLevel: false,
			AddAttribute: &model.AddAttributeReq{
				Intellect: 110,
				Physique:  10,
			},
			AddUnallocatedAttribute: &model.AddAttributeReq{
				Intellect: 100,
				Physique:  100,
			},
		})
		fn.PanicErr(err)
		log.Infof("CompleteUserAction resp: %v", GetJsonString(resp))
		debug.DebugUserInfo(ctx, UserMap[3])
	case "SetUserGear":
		e.CompleteUserChangeGear(ctx, UserMap[3], &model.UserChangeGearReq{Position: 1, PropId: 574})
		e.CompleteUserChangeGear(ctx, UserMap[3], &model.UserChangeGearReq{Position: 2, PropId: 575})
	case "InitGearConfig":
		e.InitGearConfig(ctx)
	case "mock":
		debug.DebugAddExpProp(ctx)
	}

}

// InitSignConfig implements model.EntityUseCase.
func (e *Entity) InitSignConfig(ctx context.Context) {
	panic("unimplemented")
}

// CompleteTask implements model.EntityUseCase.
// CompleteTask 完成用户任务。
// 该函数接收一个上下文、一个用户接口实例和一个用户任务ID作为参数，
// 并在数据库中更新任务状态为完成，同时更新完成时间。
// 它返回一个完成任务的响应结构体和一个错误（如果有的话）。
func (e *Entity) CompleteTask(ctx context.Context, user *model.UserEntity, userTaskId model.GetUserTaskIder) (*model.CompleteTaskResp, error) {
	// 初始化完成任务的响应结构体。
	var resp = &model.CompleteTaskResp{}
	// 获取当前时间戳，用于后续的时间比较和更新任务完成时间。
	now := time.Now().Unix()
	// 使用数据库事务处理，确保任务状态更新操作的原子性。
	err := db.Transaction(func(tx *gorm.DB) error {
		// 获取用户任务信息。
		userTaskInfo := userTaskRepo.Get(ctx, tx, userTaskId.GetUserTaskId())

		// 以下两行代码被注释掉，表明它们可能是未使用的代码片段或未来的扩展点。
		// taskInfo := taskRepo.Get(ctx, tx, userTaskInfo.GetTaskId())
		// _=taskInfo

		// 检查任务状态是否为进行中，如果不是，则返回错误。
		if userTaskInfo.Status != model.EnumUserTaskStatusDoing {
			return model.ErrUserTaskStatusNotDoing
		}

		// 检查任务是否已经到了开始时间。
		if userTaskInfo.StartTs > 0 && userTaskInfo.StartTs < now {
			return model.ErrUserTaskTimeNotStart
		}
		// 检查任务是否已经过了结束时间。
		if userTaskInfo.EndTs > 0 && userTaskInfo.EndTs > now {
			return model.ErrUserTaskTimeEnd
		}

		// 更新任务的完成时间和状态为已完成。
		userTaskInfo.CompleteTs = (now)
		userTaskInfo.Status = (model.EnumUserTaskStatusComplete)

		// 保存更新后的用户任务信息。
		userTaskRepo.Save(ctx, tx, userTaskInfo)

		// 如果没有遇到任何错误，返回nil。
		return nil
	})

	if err != nil {
		// 如果遇到错误，返回nil响应和错误。
		return nil, err
	}

	// 返回完成任务的响应结构体和nil错误，表示任务完成成功。
	return resp, nil
}

// CreateUser implements model.EntityUseCase.
func (e *Entity) CreateUser(ctx context.Context, user model.UserRepository) (model.UserRepository, error) {
	// log := log.With("traceid", GetTraceid(ctx))
	db.AutoMigrate(&model.UserEntity{})

	// err := db.Transaction(func(tx *gorm.DB) error {
	// 	user.SetName(ctx, "mock")
	// 	return user.Save(ctx, tx)
	// })

	// if err != nil {
	// 	log.Errorf("create user failed: %v", err)
	// 	return nil, err
	// }

	// log.Infof("create user: %v", GetJsonString(user))

	return user, nil
}

// Login implements model.EntityUseCase.
func (e *Entity) Login(ctx context.Context, param *model.LoginReq) (*model.LoginResp, error) {
	var user *model.UserEntity
	err := db.Transaction(func(tx *gorm.DB) error {
		if fn.Recover(func() {
			user = userRepo.Get(ctx, tx, param.Id)
		}) != nil {
			return model.ErrLoginUserNotExist
		}

		lock.Lock()
		UserMap[user.Id] = user
		UserLock[user.Id] = &sync.Mutex{}
		lock.Unlock()

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &model.LoginResp{
		User: user,
	}, nil
}
func (e *Entity) Reconnect(ctx context.Context, param *model.LoginReq) (*model.LoginResp, error) {
	var user *model.UserEntity
	err := db.Transaction(func(tx *gorm.DB) error {
		if fn.Recover(func() {
			user = userRepo.Get(ctx, tx, param.Id)
		}) != nil {
			return model.ErrLoginUserNotExist
		}

		lock.Lock()
		UserMap[user.Id] = user
		UserLock[user.Id] = &sync.Mutex{}
		lock.Unlock()

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &model.LoginResp{
		User: user,
	}, nil
}

// CompleteUserChangeGear implements model.EntityUseCase.
func (e *Entity) CompleteUserChangeGear(ctx context.Context, user *model.UserEntity, param *model.UserChangeGearReq) (*model.UserChangeGearResp, error) {
	var resp = &model.UserChangeGearResp{
		User: user,
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		// 获取最新用户信息
		user = userRepo.Get(ctx, tx, user.Id)

		// 更换装备
		userRepo.SetUserGear(ctx, tx, user, param.Position, param.PropId, e.CheckUserCanChangeGear)
		// 获取最新用户信息
		user = userRepo.Get(ctx, tx, user.Id)

		return nil
	})

	// 更新用户内存数据
	e.UpdateUser(ctx, user)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CompleteUserAction implements model.EntityUseCase.
func (e *Entity) CompleteUserAction(ctx context.Context, user *model.UserEntity, param *model.UserActionReq) (*model.CompleteAddExperienceResp, error) {
	var resp = &model.CompleteAddExperienceResp{
		LevelList: []int64{},
	}

	var err error
	err = db.Transaction(func(tx *gorm.DB) error {
		resp, err = commonCase.CompleteUserAction(ctx, tx, user, param)
		return err
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *Entity) _actionAddUserBagGear(ctx context.Context, user *model.UserEntity, param *model.UserActionReq) {
	if len(param.AddUserBagGear.Gears) == 0 {
		return
	}

}
