package usecase

import (
	"api/pkg/model"
	"context"

	"gorm.io/gorm"
)

func NewPropUseCase() model.PropUseCase {
	return &propUseCase{}
}

type propUseCase struct {
}

// SendReward implements model.PropUseCase.
func (p *propUseCase) SendReward(ctx context.Context, tx *gorm.DB, user *model.UserEntity, now int64, taskRewarder *model.TaskReward) {
	for _, reward := range taskRewarder.GetProps() {
		p._SendRewardOne(ctx, tx, user, now, reward)
	}
}
func (p *propUseCase) _SendRewardOne(ctx context.Context, tx *gorm.DB, user *model.UserEntity, now int64, rewarder model.Reward) {
	propInfo := propRepo.Get(ctx, tx, rewarder.GetPropId())

	switch propInfo.GetType() {
	case model.PropTypeGear:
		useBag := model.NewUserBager(user, propInfo)
		useBag.SetAddTs(now)
		useBag.SetEndTs(0)
		useBag.SetCount(rewarder.GetCount())

		useBagRepo.Create(ctx, tx, useBag)
	case model.PropTypeExperience:
		commonCase.CompleteUserAction(ctx, tx, user, &model.UserActionReq{
			AddExperience:               rewarder.GetCount() * propInfo.ConvertExperiencer().GetExperience(),
			AddLevel_1:                  false,
			ConvertAllExperienceToLevel: false,
			AddAttribute:                &model.AddAttributeReq{},
			AddUserBagGear:              model.AddUserBagGear{},
		})
	}
}
