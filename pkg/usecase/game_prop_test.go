package usecase

import (
	"api/pkg/model"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var enti = NewEntity()

func TestActionUser(t *testing.T) {

	ctx := context.Background()
	err := db.Transaction(func(tx *gorm.DB) error {
		loginResp, err := enti.Login(ctx, &model.LoginReq{
			Id: 3,
		})
		require.Nil(t, err)

		require.Contains(t, UserMap, loginResp.User.Id)
		{
			beforeUser := UserMap[loginResp.User.Id].CopyEntity()

			// param *model.UserActionReq
			param := &model.UserActionReq{
				AddExperience:               110,
				AddLevel_1:                  false,
				ConvertAllExperienceToLevel: false,
				AddAttribute:                &model.AddAttributeReq{},
				AddUnallocatedAttribute:     &model.AddAttributeReq{},
				AddUserBagGear:              model.AddUserBagGear{},
			}
			_, err = commonCase.CompleteUserAction(ctx, tx, loginResp.User, param)
			require.Nil(t, err)

			afterUser := userRepo.Get(ctx, tx, loginResp.User.Id)
			require.Equal(t, beforeUser.Experience+param.AddExperience, afterUser.Experience)
		}

		{
			user := userRepo.Get(ctx, tx, loginResp.User.Id)
			beforeUser := user.CopyEntity()
			param := &model.UserActionReq{
				AddExperience:               -100,
				AddLevel_1:                  false,
				ConvertAllExperienceToLevel: false,
				AddAttribute:                &model.AddAttributeReq{},
				AddUnallocatedAttribute:     &model.AddAttributeReq{},
				AddUserBagGear:              model.AddUserBagGear{},
			}
			_, err = commonCase.CompleteUserAction(ctx, tx, user, param)
			require.Nil(t, err)

			afterUser := userRepo.Get(ctx, tx, user.Id)
			require.Equal(t, beforeUser.Experience+param.AddExperience, afterUser.Experience)
		}

		{
			user := userRepo.Get(ctx, tx, loginResp.User.Id)
			beforeUser := user.CopyEntity()
			param := &model.UserActionReq{
				AddExperience:               model.LevelConfigMap[loginResp.User.Level].NeedExperience + 10,
				AddLevel_1:                  true,
				ConvertAllExperienceToLevel: false,
				AddAttribute:                &model.AddAttributeReq{},
				AddUnallocatedAttribute:     &model.AddAttributeReq{},
				AddUserBagGear:              model.AddUserBagGear{},
			}
			_, err = commonCase.CompleteUserAction(ctx, tx, user, param)
			require.Nil(t, err)

			afterUser := userRepo.Get(ctx, tx, beforeUser.Id)

			require.Equal(t,
				beforeUser.Experience+param.AddExperience-model.LevelConfigMap[beforeUser.Level].NeedExperience,
				afterUser.Experience)
			require.Equal(t, beforeUser.Level+1, afterUser.Level)
			require.Equal(t, beforeUser.UnallocatedIntellect+model.LevelConfigMap[beforeUser.Level].Intellect, afterUser.UnallocatedIntellect)
			require.Equal(t, beforeUser.UnallocatedPhysique+model.LevelConfigMap[beforeUser.Level].Physique, afterUser.UnallocatedPhysique)
			require.Equal(t, beforeUser.UnallocatedEndurance+model.LevelConfigMap[beforeUser.Level].Endurance, afterUser.UnallocatedEndurance)
		}

		{
			user := userRepo.Get(ctx, tx, loginResp.User.Id)
			beforeUser := user.CopyEntity()
			param := &model.UserActionReq{
				AddExperience:               0,
				AddLevel_1:                  false,
				ConvertAllExperienceToLevel: false,
				AddAttribute: &model.AddAttributeReq{
					Intellect: 10,
					Physique:  20,
				},
				AddUnallocatedAttribute: &model.AddAttributeReq{},
				AddUserBagGear:          model.AddUserBagGear{},
			}
			_, err = commonCase.CompleteUserAction(ctx, tx, user, param)
			require.Nil(t, err)

			afterUser := userRepo.Get(ctx, tx, beforeUser.Id)

			require.Equal(t, beforeUser.UnallocatedIntellect-param.AddAttribute.Intellect, afterUser.UnallocatedIntellect)
			require.Equal(t, beforeUser.UnallocatedPhysique-param.AddAttribute.Physique, afterUser.UnallocatedPhysique)
			require.Equal(t, beforeUser.UnallocatedEndurance-param.AddAttribute.Endurance, afterUser.UnallocatedEndurance)
			require.Equal(t, beforeUser.Intellect+param.AddAttribute.Intellect, afterUser.Intellect)
			require.Equal(t, beforeUser.Physique+param.AddAttribute.Physique, afterUser.Physique)
			require.Equal(t, beforeUser.Endurance-param.AddAttribute.Endurance, afterUser.Endurance)
		}

		{
			user := userRepo.Get(ctx, tx, loginResp.User.Id)
			beforeUser := user.CopyEntity()
			param := &model.UserActionReq{
				AddExperience:               0,
				AddLevel_1:                  false,
				ConvertAllExperienceToLevel: false,
				AddAttribute: &model.AddAttributeReq{
					Intellect: 0,
					Physique:  0,
				},
				AddUnallocatedAttribute: &model.AddAttributeReq{
					Intellect: 10,
					Physique:  -10,
					Endurance: 0,
				},
				AddUserBagGear: model.AddUserBagGear{},
			}
			_, err = commonCase.CompleteUserAction(ctx, tx, user, param)
			require.Nil(t, err)

			afterUser := userRepo.Get(ctx, tx, beforeUser.Id)

			require.Equal(t, beforeUser.UnallocatedIntellect+param.AddUnallocatedAttribute.Intellect, afterUser.UnallocatedIntellect)
			require.Equal(t, beforeUser.UnallocatedPhysique+param.AddUnallocatedAttribute.Physique, afterUser.UnallocatedPhysique)
			require.Equal(t, beforeUser.UnallocatedEndurance+param.AddUnallocatedAttribute.Endurance, afterUser.UnallocatedEndurance)
		}

		{
			user := userRepo.Get(ctx, tx, loginResp.User.Id)
			beforeUser := user.CopyEntity()
			param := &model.UserActionReq{
				AddExperience:               0,
				AddLevel_1:                  false,
				ConvertAllExperienceToLevel: false,
				AddAttribute: &model.AddAttributeReq{
					Intellect: 0,
					Physique:  0,
				},
				AddUnallocatedAttribute: &model.AddAttributeReq{
					Endurance: 0,
				},
				AddUserBagGear: model.AddUserBagGear{
					Gears: []*model.GearEntity{},
				},
			}
			_, err = commonCase.CompleteUserAction(ctx, tx, user, param)
			require.Nil(t, err)

			afterUser := userRepo.Get(ctx, tx, beforeUser.Id)

			require.Equal(t, beforeUser.UnallocatedIntellect+param.AddUnallocatedAttribute.Intellect, afterUser.UnallocatedIntellect)
			require.Equal(t, beforeUser.UnallocatedPhysique+param.AddUnallocatedAttribute.Physique, afterUser.UnallocatedPhysique)
			require.Equal(t, beforeUser.UnallocatedEndurance+param.AddUnallocatedAttribute.Endurance, afterUser.UnallocatedEndurance)
		}

		return nil
	})

	require.Nil(t, err)
}

func TestGameProp(t *testing.T) {

	ctx := context.Background()

	err := db.Transaction(func(tx *gorm.DB) error {
		log.Debug("给道具表 增加一个经验值道具")
		var before *model.PropEntity
		var after *model.PropEntity

		var insertId int64
		{
			before = model.NewProper(model.PropTypeExperience, model.NewExperiencer(1000))
			before.SetStartTs(time.Now().Unix())
			before.SetEndTs(time.Now().Add(time.Hour * 24 * 365 * 10).Unix())

			propRepo.Create(ctx, tx, before)
			insertId = before.GetPropId()
		}
		{
			after = propRepo.Get(ctx, tx, insertId)
			log.Infof("添加成功, PropId: %d, Content: %s", after.GetPropId(), after.GetContent())
			log.Infof("data: %v", GetJsonString(after.ConvertExperiencer()))

			require.True(t, before.Equal(after))

			require.Equal(t, before.ConvertExperiencer().GetExperience(), after.ConvertExperiencer().GetExperience())
			require.Equal(t, before.ConvertExperiencer().GetExperience(), 1000)
		}

		return nil
	})

	require.Nil(t, err)

}
