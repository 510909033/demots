package usecase

import (
	"api/pkg/model"
	"context"

	"gorm.io/gorm"
)

func NewCommonUseCase() model.CommonUseCase {
	return &commonUseCase{}
}

type commonUseCase struct {
}

// GetUserDayTaskList implements model.CommonUseCase.
func (c *commonUseCase) GetUserDayTaskList(ctx context.Context, user *model.UserEntity, now int64) (*model.UserDayTaskListResp, error) {
	panic("unimplemented")
}

// CompleteUserAction implements model.CommonUseCase.
func (c *commonUseCase) CompleteUserAction(ctx context.Context, tx *gorm.DB, user *model.UserEntity, param *model.UserActionReq) (*model.CompleteAddExperienceResp, error) {
	var resp = &model.CompleteAddExperienceResp{
		LevelList: []int64{},
	}

	// 获取最新用户信息
	user = userRepo.Get(ctx, tx, user.Id)

	// 直接操作经验值
	if param.AddExperience != 0 {
		if param.AddExperience < 0 && param.AddExperience+user.Experience < 0 {
			// 经验值不足
			return nil, model.ErrExperienceNotEnough
		}
		userRepo.AddExperience(ctx, tx, user, param.AddExperience)
		// 获取最新用户信息
		user = userRepo.Get(ctx, tx, user.Id)
	}

	// 等级+1
	if param.AddLevel_1 {
		// 判断经验值是否足够
		if user.Experience < model.LevelConfigMap[user.Level].NeedExperience {
			// 经验值不足
			return nil, model.ErrExperienceNotEnough
		}

		// 等级+1
		userRepo.AddLevel(ctx, tx, user, 1)
		// 经验值减少
		userRepo.AddExperience(ctx, tx, user, -model.LevelConfigMap[user.Level].NeedExperience)
		// 增加未分配属性
		userRepo.AddUnallocatedAttribute(ctx, tx, user, model.LevelConfigMap[user.Level])

		// 获取最新用户信息
		user = userRepo.Get(ctx, tx, user.Id)

		resp.LevelList = append(resp.LevelList, user.Level)
	}

	// 将全部经验值转换为提升等级
	if param.ConvertAllExperienceToLevel {
		// 提升等级
		for user.Experience >= model.LevelConfigMap[user.Level].NeedExperience {
			log.Infof("now level: %v, need experience: %v, total experience: %v", user.Level, model.LevelConfigMap[user.Level].NeedExperience, user.Experience)

			// 等级+1
			userRepo.AddLevel(ctx, tx, user, 1)
			// 经验值减少
			userRepo.AddExperience(ctx, tx, user, -model.LevelConfigMap[user.Level].NeedExperience)
			// 增加属性
			userRepo.AddUnallocatedAttribute(ctx, tx, user, model.LevelConfigMap[user.Level])

			// 获取最新用户信息
			user = userRepo.Get(ctx, tx, user.Id)

			resp.LevelList = append(resp.LevelList, user.Level)
		}
	}

	// 分配 【未分配属性】
	if param.AddUnallocatedAttribute != nil {
		userRepo.AddUnallocatedAttribute(ctx, tx, user, &model.LevelConfig{
			NeedExperience: 0,
			Intellect:      param.AddUnallocatedAttribute.Intellect,
			Physique:       param.AddUnallocatedAttribute.Physique,
			// TODO W 继续补充
		})

		// 获取最新用户信息
		user = userRepo.Get(ctx, tx, user.Id)
	}

	// 分配属性
	if param.AddAttribute != nil {
		// 增加已分配属性
		userRepo.AddAttribute(ctx, tx, user, param.AddAttribute)
		// 减少未分配属性
		userRepo.AddUnallocatedAttribute(ctx, tx, user, &model.LevelConfig{
			NeedExperience: 0,
			Intellect:      -param.AddAttribute.Intellect,
			Physique:       -param.AddAttribute.Physique,
			// TODO W 继续补充
		})

		// 获取最新用户信息
		user = userRepo.Get(ctx, tx, user.Id)
	}

	// 更新用户内存数据
	c.UpdateUser(ctx, user)

	resp.LeaveExperience = user.Experience
	resp.Cfg = model.LevelConfigMap[user.Level]
	resp.User = user
	return resp, nil
}
func (e *commonUseCase) UpdateUser(ctx context.Context, user *model.UserEntity) {
	lock.Lock()
	UserMap[user.Id] = user
	lock.Unlock()
}

func (c *commonUseCase) _useUserBagList(ctx context.Context, tx *gorm.DB, user *model.UserEntity, bags []model.UseUserBag) {
	if len(bags) == 0 {
		return
	}
}
