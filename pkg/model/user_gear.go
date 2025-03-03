package model

const (
	EnumUserGearPositionWeapon_1   = 1  // 武器
	EnumUserGearPositionShield_2   = 2  // 防御
	EnumUserGearPositionArm_3      = 3  // 护腕
	EnumUserGearPositionShoe_4     = 4  // 鞋子
	EnumUserGearPositionHelmet_5   = 5  // 头盔
	EnumUserGearPositionNecklace_6 = 6  // 项链
	EnumUserGearPositionClothes_7  = 7  // 衣袍
	EnumUserGearPositionRing_8     = 8  // 戒指
	EnumUserGearPositionJade_9     = 9  // 玉佩
	EnumUserGearPositionArm_3_2    = 10 // 护腕2
	EnumUserGearPositionRing_8_2   = 11 // 戒指2
)

type UserGearEntity struct {
	PropId int64
	// 位置， 1武器， 2护盾，3护腕1 4鞋子， 5头盔， 6项链， 7衣袍， 8戒指1, 9玉佩,  额外： 10护腕2， 11戒指2
	Position int64
}
