package main

import (
	"api/pkg/repo"
	"api/pkg/usecase"
	"context"
)

func main() {
	entity := usecase.NewEntity(
		repo.NewUserRepository(),
	)
	entity.Debug(context.Background())
	// entity.CreateUser(context.Background(), &repo.User{
	// 	Id:         0,
	// 	Name:       "new1",
	// 	Experience: 0,
	// })

}
