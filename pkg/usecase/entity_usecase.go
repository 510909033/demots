package usecase

import (
	"api/pkg/model"
	"api/pkg/repo"
	"context"
	"encoding/json"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var _log, _ = zap.NewProduction()
var log = _log.Sugar()

func GetTraceid(ctx context.Context) string {
	return uuid.NewString()[:8]
}

func GetJsonString(val any) string {
	v, _ := json.Marshal(val)
	return string(v)
}

func NewEntity() model.EntityUseCase {
	db, err := gorm.Open(sqlite.Open("/tmp/test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	repo.SetDB(db)

	return &Entity{
		UserList: make([]*model.UserRepository, 0),
		db:       db,
	}
}

type Entity struct {
	UserList []*model.UserRepository
	db       *gorm.DB
}

// Debug implements model.EntityUseCase.
func (e *Entity) Debug(ctx context.Context) {
	e.db.AutoMigrate(&model.PropEntity{})

	var data = &model.PropEntity{
		Id:       0,
		PropType: model.PropTypeExpDirect,
		StartTs:  time.Now().Unix(),
		EndTs:    0,
	}
	repo.NewPropRepository().Create(ctx, e.db, data)
	log.Info(GetJsonString(data))

	// data2 := repo.NewPropRepository().Get(ctx, e.db, 51)
	// log.Info(GetJsonString(data2))

}

// InitSignConfig implements model.EntityUseCase.
func (e *Entity) InitSignConfig(ctx context.Context) {
	panic("unimplemented")
}

// CompleteTask implements model.EntityUseCase.
func (e *Entity) CompleteTask(ctx context.Context, user model.UserRepository, task model.Task) (*model.CompleteTaskResp, error) {
	panic("unimplemented")
}

// CreateUser implements model.EntityUseCase.
func (e *Entity) CreateUser(ctx context.Context, user model.UserRepository) (model.UserRepository, error) {
	log := log.With("traceid", GetTraceid(ctx))
	e.db.AutoMigrate(&repo.User{})

	err := e.db.Transaction(func(tx *gorm.DB) error {
		user.SetName(ctx, "mock")
		return user.Save(ctx, tx)
	})

	if err != nil {
		log.Errorf("create user failed: %v", err)
		return nil, err
	}

	log.Infof("create user: %v", GetJsonString(user))

	return user, nil
}

// Login implements model.EntityUseCase.
func (e *Entity) Login(ctx context.Context, param *model.LoginReq) (model.LoginResp, error) {
	panic("unimplemented")
}
