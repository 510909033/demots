package repo

import (
	"api/pkg/model"
	"api/pkg/utils/fn"
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var log *zap.SugaredLogger

func SetLog(l *zap.SugaredLogger) {
	log = l
}

func GetLog(user *model.UserEntity) *zap.SugaredLogger {
	return log.With("userId", user.Id)
}

func NewUserRepository() model.UserRepository {
	return &userRepo{}
}

type userRepo struct {
	Id         int64  `json:"id" gorm:"primaryKey"`
	Name       string `json:"name" gorm:"column:name;type:varchar(255);not null;default:''"`
	Experience int64  `json:"experience" gorm:"column:experience;type:bigint;not null;default:0"`
}

// SetUserGear implements model.UserRepository.
func (u *userRepo) SetUserGear(ctx context.Context, tx *gorm.DB, user *model.UserEntity, position int64, propId int64, check func(ctx context.Context, user *model.UserEntity, param *model.UserChangeGearReq) bool) {
	if user.UserGear == nil {
		user.UserGear = make(model.TUserGear)
	}

	if !check(ctx, user, &model.UserChangeGearReq{
		PropId:   propId,
		Position: position,
	}) {
		fn.PanicErr(model.ErrUserNotSatisfyGear)
	}

	user.UserGear[position] = &model.UserGearEntity{
		PropId:   propId,
		Position: position,
	}

	GetLog(user).Infof("userRepo.SetUserGear, position: %v, propId: %v", position, propId)
	err := tx.WithContext(ctx).Model(user).Where("id = ?", user.Id).Update("user_gear", user.UserGear).Error

	fn.PanicErr(err)

}

// AddAttribute implements model.UserRepository.
func (u *userRepo) AddAttribute(ctx context.Context, tx *gorm.DB, user *model.UserEntity, param *model.AddAttributeReq) {
	if param.Intellect != 0 {
		GetLog(user).Infof("userRepo.AddAttribute, added intellect: %v", param.Intellect)
		err := tx.WithContext(ctx).Model(user).Where("id = ?", user.Id).Update("intellect", gorm.Expr("intellect + ?", param.Intellect)).Error
		fn.PanicErr(err)
	}
	if param.Physique != 0 {
		GetLog(user).Infof("userRepo.AddAttribute, added physique: %v", param.Physique)
		err := tx.WithContext(ctx).Model(user).Where("id = ?", user.Id).Update("physique", gorm.Expr("physique + ?", param.Physique)).Error
		fn.PanicErr(err)
	}
}

// AddUnallocatedAttribute implements model.UserRepository.
func (u *userRepo) AddUnallocatedAttribute(ctx context.Context, tx *gorm.DB, user *model.UserEntity, cfg *model.LevelConfig) {
	if cfg.Intellect != 0 {
		GetLog(user).Infof("userRepo.AddUnlocatedAttribute, added unallocated_intellect: %v", cfg.Intellect)
		err := tx.WithContext(ctx).Model(user).Where("id = ?", user.Id).Update("unallocated_intellect", gorm.Expr("unallocated_intellect + ?", cfg.Intellect)).Error
		fn.PanicErr(err)
	}
	if cfg.Physique != 0 {
		GetLog(user).Infof("userRepo.AddUnlocatedAttribute, added unallocated_physique: %v", cfg.Physique)
		err := tx.WithContext(ctx).Model(user).Where("id = ?", user.Id).Update("unallocated_physique", gorm.Expr("unallocated_physique + ?", cfg.Physique)).Error
		fn.PanicErr(err)
	}
}

// AddLevel implements model.UserRepository.
func (u *userRepo) AddLevel(ctx context.Context, tx *gorm.DB, user *model.UserEntity, level int64) {
	err := tx.WithContext(ctx).Model(user).Where("id = ?", user.Id).Update("level", gorm.Expr("level + ?", level)).Error
	GetLog(user).Infof("userRepo.AddLevel, old level: %v, add level: %v", user.Level, level)

	fn.PanicErr(err)
}

// Create implements model.UserRepository.
func (u *userRepo) Create(ctx context.Context, tx *gorm.DB, user *model.UserEntity) {
	err := tx.WithContext(ctx).Create(&user).Error
	if err != nil {
		panic(err)
	}
}

// Get implements model.UserRepository.
func (u *userRepo) Get(ctx context.Context, tx *gorm.DB, userId int64) *model.UserEntity {
	var user model.UserEntity
	err := tx.WithContext(ctx).Where("id = ?", userId).First(&user).Error

	fn.PanicErr(err)

	return &user
}

// AddExperience implements model.User.
func (u *userRepo) AddExperience(ctx context.Context, tx *gorm.DB, user *model.UserEntity, experience int64) {
	GetLog(user).Infof("userRepo.AddExperience, added experience: %v", experience)
	err := tx.WithContext(ctx).Model(user).Where("id = ?", user.Id).Update("experience", gorm.Expr("experience + ?", experience)).Error

	fn.PanicErr(err)
}
