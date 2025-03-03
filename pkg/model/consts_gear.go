package model

import "fmt"

// 装备配置结构
type GearConfig struct {
}

// 装备配置， key是等级
var GearConfigMap = map[int64][]*GearEntity{}

func initGearConfig() {
	var gearId int64

	for gearType := 1; gearType <= 9; gearType++ {

		base := GearEntity{
			Name:       "",
			Type:       int64(gearType),
			Color:      EnumGearColorWhite_0,
			MinLevel:   1,
			MinAttack:  10,
			MaxAttack:  15,
			MinDefense: 0,
			MaxDefense: 5,
		}

		for level := 1; level < 1000; level = level + 100 {
			for color := 0; color <= 6; color++ {
				gearId++

				baseCopy := base
				baseCopy.Name = fmt.Sprintf("level(%d)_%s_%s", level, TransGearType(int64(gearType)), TransGearColor(int64(color)))
				baseCopy.Color = int64(color)
				baseCopy.MinAttack += int64(color) * int64(level)
				baseCopy.MaxAttack += int64(color) * int64(level)
				baseCopy.MinDefense += int64(color) * int64(level)
				baseCopy.MaxDefense += int64(color) * int64(level)

				GearConfigMap[int64(level)] = append(GearConfigMap[int64(level)], &baseCopy)
			}
		}
	}
}
