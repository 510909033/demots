package usecase

import (
	"api/pkg/model"
	"context"
)

// LoadAllConfigToClient implements model.EntityUseCase.
func (e *Entity) LoadAllConfigToClient(ctx context.Context, user *model.UserEntity) (*model.LoadAllConfigToClientResp, error) {
	var resp = &model.LoadAllConfigToClientResp{}

	return resp, nil
}
