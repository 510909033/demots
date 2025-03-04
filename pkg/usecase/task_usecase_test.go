package usecase

import (
	"api/pkg/model"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// TestCreate
func TestCreate(t *testing.T) {
	ctx := context.Background()

	db.Transaction(func(tx *gorm.DB) error {
		param := &model.TaskEntity{
			Id:       0,
			TaskType: model.EnumTaskTypeDay,
			MaxCount: 2,
			StartTs:  0,
			EndTs:    0,
			Deadline: 2,
			Award: &model.TaskReward{
				Props: []model.Reward{
					{
						PropId: propRepo.GetRandomProp(ctx, tx).GetPropId(),
						Count:  2,
					},
				},
			},
		}
		taskRepo.Create(ctx, tx, param)

		assert.Greater(t, param.Id, int64(0))

		afterTask := taskRepo.Get(ctx, tx, param.Id)
		assert.True(t, afterTask.Equal(param))

		return nil
	})

}
