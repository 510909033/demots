package model

import (
	"context"

	"gorm.io/gorm"
)

type CommonUseCase interface {
	CompleteUserAction(ctx context.Context, tx *gorm.DB, user *UserEntity, param *UserActionReq) (*CompleteAddExperienceResp, error)
	UpdateUser(ctx context.Context, user *UserEntity)
}
