package usecase

import "api/pkg/model"

type useBagUseCase struct {
}

func NewUserBagUseCase() model.UserBagUseCase {
	return &useBagUseCase{}
}
