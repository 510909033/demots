package usecase

import (
	"api/pkg/model"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// TestCreate
func TestCreate(t *testing.T) {
	ctx := context.Background()

	db.Transaction(func(tx *gorm.DB) error {
		param := &model.TaskEntity{
			Id:       0,
			TaskName: "",
			TaskType: model.EnumTaskTypeDay,
			MaxCount: 2,
			StartTs:  0,
			EndTs:    0,
			Deadline: 2,
			Award:    &model.TaskReward{Props: []model.Reward{{PropId: propRepo.GetRandomProp(ctx, tx).GetPropId(), Count: 2}}},
		}
		taskRepo.Create(ctx, tx, param)

		assert.Greater(t, param.Id, int64(0))

		afterTask := taskRepo.Get(ctx, tx, param.Id)
		assert.True(t, afterTask.Equal(param))

		return nil
	})

}

// TestCreate
func TestUserCreateTask(t *testing.T) {
	ctx := context.Background()

	loginResp, err := enti.Login(ctx, &model.LoginReq{
		Id: 3,
	})
	require.Nil(t, err)
	require.Contains(t, UserMap, loginResp.User.Id)

	beforeUser := UserMap[loginResp.User.Id]

	// 获取一个任务id
	var taskInfo model.TaskEntity
	err = db.Model(&model.TaskEntity{}).Where("1=1").Order("RAND()").First(&taskInfo).Error
	require.Nil(t, err)
	require.Greater(t, taskInfo.Id, int64(0))

	{
		resp, err := enti.UserCreateTask(ctx, beforeUser, &taskInfo)
		require.Nil(t, err)
		require.NotNil(t, resp)
	}

}
func TestUserTaskSummary(t *testing.T) {
	ctx := context.Background()

	loginResp, err := enti.Login(ctx, &model.LoginReq{
		Id: 3,
	})
	require.Nil(t, err)
	require.Contains(t, UserMap, loginResp.User.Id)

	debug.DebugUserTaskInfo(ctx, loginResp.User)

}
