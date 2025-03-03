package usecase

import (
	"api/pkg/model"
	"api/pkg/utils/fn"
	"context"

	"gorm.io/gorm"
)

// 初始化装备数据到 gear_entity 表中, 只会执行一次
func (e *Entity) InitGearConfig(ctx context.Context) {
	db.Transaction(func(tx *gorm.DB) error {
		for level, gearList := range model.GearConfigMap {
			for _, gearInfo := range gearList {

				e.prop.Create(ctx, tx, &model.PropEntity{
					Id:   0,
					Type: model.PropTypeGear,
					Entity: &model.GearEntity{
						Name:       gearInfo.Name,
						Type:       gearInfo.Type,
						Color:      gearInfo.Color,
						MinLevel:   level,
						MinAttack:  gearInfo.MinAttack,
						MaxAttack:  gearInfo.MaxAttack,
						MinDefense: gearInfo.MinDefense,
						MaxDefense: gearInfo.MaxDefense,
					},
				})
				// e.gear.Create(ctx, tx, &model.GearEntity{
				// 	Id:         gearInfo.Id,
				// 	Name:       gearInfo.Name,
				// 	Type:       gearInfo.Type,
				// 	Color:      gearInfo.Color,
				// 	MinLevel:   level,
				// 	MinAttack:  gearInfo.MinAttack,
				// 	MaxAttack:  gearInfo.MaxAttack,
				// 	MinDefense: gearInfo.MinDefense,
				// 	MaxDefense: gearInfo.MaxDefense,
				// })
			}
		}

		return nil
	})

}

// // 将装备数据从表加载到内容（gearAll）变量中
// func (e *Entity) InitGearAllToMemory(ctx context.Context) {
// 	var list []*model.GearEntity
// 	fn.PanicErr(db.Model(&model.GearEntity{}).Where("1=1").Find(&list).Error)
// 	for _, gearInfo := range list {
// 		gearAll[gearInfo.Id] = gearInfo
// 	}
// }

// 将所有道具填充到内容中
func (e *Entity) InitPropAllToMemory(ctx context.Context) {
	var list []*model.PropEntity
	fn.PanicErr(db.Model(&model.PropEntity{}).Where("1=1").Find(&list).Error)
	for _, propInfo := range list {
		propInfo.Convert()
		propAll[propInfo.Id] = propInfo
	}
}
