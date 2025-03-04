package usecase

import (
	"api/pkg/model"
	"context"
	"testing"
	"time"

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

	// 创建一个任务
	taskInfo := &model.TaskEntity{
		Id:       0,
		TaskName: time.Now().Format("task__2006-01-02 15:04:05"),
		TaskType: model.EnumTaskTypeDay,
		MaxCount: 2,
		StartTs:  0,
		EndTs:    0,
		Deadline: 5,
		Award:    &model.TaskReward{},
	}
	taskRepo.Create(ctx, db, taskInfo)
	require.Greater(t, taskInfo.Id, int64(0))

	{
		// userTaskRepo.GetUserTaskCount(ctx, db, beforeUser, taskInfo.Id, 0, 0)

		resp, err := enti.UserCreateTask(ctx, beforeUser, taskInfo.Id)
		require.Nil(t, err)
		require.NotNil(t, resp)

		cnt := userTaskRepo.GetUserTaskCount(ctx, db, beforeUser, taskInfo.Id, GetTodayStartTs(), GetTodayEndTs())
		assert.Equal(t, int64(1), cnt)

		resp, err = enti.UserCreateTask(ctx, beforeUser, taskInfo.Id)
		require.Nil(t, err)
		require.NotNil(t, resp)

		cnt = userTaskRepo.GetUserTaskCount(ctx, db, beforeUser, taskInfo.Id, GetTodayStartTs(), GetTodayEndTs())
		assert.Equal(t, int64(2), cnt)

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
