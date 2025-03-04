package usecase

import (
	"api/pkg/model"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompleteUserChangeGear(t *testing.T) {
	ctx := context.Background()

	loginResp, err := enti.Login(ctx, &model.LoginReq{
		Id: 3,
	})
	require.Nil(t, err)

	require.Contains(t, UserMap, loginResp.User.Id)

	{
		user := UserMap[loginResp.User.Id]
		beforeUser := user.CopyEntity()

		param := &model.UserChangeGearReq{
			Position: model.EnumUserGearPositionHelmet_5,
			PropId:   601,
		}

		resp, err := enti.CompleteUserChangeGear(ctx, user, param)

		afterUser := UserMap[beforeUser.Id]

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, user.Id, resp.User.Id)
		assert.Equal(t, param.PropId, afterUser.GetUserGear(param.Position).PropId)
	}
}

// CheckUserCanChangeGear
func TestCheckUserCanChangeGear(t *testing.T) {
	ctx := context.Background()
	loginResp, err := enti.Login(ctx, &model.LoginReq{
		Id: 3,
	})
	require.Nil(t, err)

	require.Contains(t, UserMap, loginResp.User.Id)

	{
		user := UserMap[loginResp.User.Id]

		param := &model.UserChangeGearReq{
			Position: model.EnumUserGearPositionHelmet_5,
			PropId:   601,
		}

		canChangeGear := enti.CheckUserCanChangeGear(ctx, user, param)
		assert.True(t, canChangeGear)
	}

	{
		user := UserMap[loginResp.User.Id]

		var propInfo model.PropEntity
		err = db.Model(&model.PropEntity{}).Where("type=?", model.PropTypeGear).Order("RAND()").Find(&propInfo).Error
		require.Nil(t, err)
		require.Greater(t, propInfo.Id, int64(0))

		t.Logf("propInfo.Id=%d", propInfo.Id)

		param := &model.UserChangeGearReq{
			Position: propInfo.ConvertGearEntity().Type,
			PropId:   propInfo.Id,
		}

		canChangeGear := enti.CheckUserCanChangeGear(ctx, user, param)
		assert.Equal(t, canChangeGear, propInfo.ConvertGearEntity().MinLevel <= user.Level)
	}
}
