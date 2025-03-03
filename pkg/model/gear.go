package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

const (
	// 类型

	EnumGearTypeWeapon_1   = 1 // 武器
	EnumGearTypeShield_2   = 2 // 防御
	EnumGearTypeArm_3      = 3 // 护腕
	EnumGearTypeShoe_4     = 4 // 鞋子
	EnumGearTypeHelmet_5   = 5 // 头盔
	EnumGearTypeNecklace_6 = 6 // 项链
	EnumGearTypeClothes_7  = 7 // 衣袍
	EnumGearTypeRing_8     = 8 // 戒指
	EnumGearTypeJade_9     = 9 // 玉佩

	// 颜色品质

	EnumGearColorWhite_0  = 0 // 白色
	EnumGearColorBlue_1   = 1 // 蓝色
	EnumGearColorYellow_2 = 2 // 黄色
	EnumGearColorGreen_3  = 3 // 绿色
	EnumGearColorPurple_4 = 4 // 紫色
	EnumGearColorOrange_5 = 5 // 橙色
	EnumGearColorRed_6    = 6 // 红色

	// // 质量品质

	// EnumGearQualityRough_0     = 0 // 粗糙
	// EnumGearQualityCommon_1    = 1 // 普通
	// EnumGearQualityGood_2      = 2 // 精良
	// EnumGearQualityExcellent_3 = 3 // 优秀
	// EnumGearQualityEpic_4      = 4 // 史诗
	// EnumGearQualityLegendary_5 = 5 // 传说

)

// 装备
type GearEntity struct {
	// // id
	// Id int64
	// 显示名称
	Name string
	// 类型， 1武器， 2护盾，3护腕 4鞋子， 5头盔， 6项链， 7衣袍， 8戒指， 9玉佩
	Type int64
	// 颜色品质， 0白色， 1蓝色， 2黄色 3绿色， 4紫色， 5橙色， 6红色
	Color int64
	// // // 质量品质 0粗糙， 1普通， 2精良， 3优秀， 4史诗， 5传说
	// Quality int64
	// 使用最低等级， 0表示不限制
	MinLevel int64
	// 最小物理攻击
	MinAttack int64
	// 最大物理攻击
	MaxAttack int64
	// 最小防御
	MinDefense int64
	// 最大防御
	MaxDefense int64
}

type GearRepository interface {
	Create(ctx context.Context, tx *gorm.DB, param *GearEntity)
}

func TransGearType(gearType int64) string {
	switch gearType {
	case EnumGearTypeWeapon_1:
		return "武器"
	case EnumGearTypeShield_2:
		return "防御"
	case EnumGearTypeArm_3:
		return "护腕"
	case EnumGearTypeShoe_4:
		return "鞋子"
	case EnumGearTypeHelmet_5:
		return "头盔"
	case EnumGearTypeNecklace_6:
		return "项链"
	case EnumGearTypeClothes_7:
		return "衣袍"
	case EnumGearTypeRing_8:
		return "戒指"
	case EnumGearTypeJade_9:
		return "玉佩"
	default:
		return fmt.Sprintf("unknown(%d)", gearType)
	}
}

func TransGearColor(gearColor int64) string {
	switch gearColor {
	case EnumGearColorWhite_0:
		return "白色"
	case EnumGearColorBlue_1:
		return "蓝色"
	case EnumGearColorYellow_2:
		return "黄色"
	case EnumGearColorGreen_3:
		return "绿色"
	case EnumGearColorPurple_4:
		return "紫色"
	case EnumGearColorOrange_5:
		return "橙色"
	case EnumGearColorRed_6:
		return "红色"
	default:
		return fmt.Sprintf("unknown(%d)", gearColor)
	}
}
