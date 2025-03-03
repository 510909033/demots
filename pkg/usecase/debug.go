package usecase

import (
	"api/pkg/model"
	"context"
	"fmt"
	"strings"
)

type _debug struct {
}

func NewDebug() *_debug {
	return &_debug{}
}

// 用户信息概览
func (d *_debug) DebugUserInfo(ctx context.Context, user *model.UserEntity) {
	var msg = make([]string, 0, 100)
	msg = append(msg, "用户信息概览")
	msg = append(msg, fmt.Sprintf("用户: %s(%d)", user.Name, user.Id))
	msg = append(msg, fmt.Sprintf("等级: %d", user.Level))
	msg = append(msg, fmt.Sprintf("经验: %d", user.Experience))
	msg = append(msg, fmt.Sprintf("智力: %d", user.Intellect))
	msg = append(msg, fmt.Sprintf("体质: %d", user.Physique))
	msg = append(msg, fmt.Sprintf("耐力: %d", user.Endurance))
	msg = append(msg, fmt.Sprintf("未分配的智力: %d", user.UnallocatedIntellect))
	msg = append(msg, fmt.Sprintf("未分配的体质: %d", user.UnallocatedPhysique))
	msg = append(msg, fmt.Sprintf("未分配的耐力: %d", user.UnallocatedEndurance))
	msg = append(msg, fmt.Sprintf("用户装备:"))

	for i := 0; i < 10; i++ {
		pos := int64(i)
		if user.UserGear[pos] == nil {
			msg = append(msg, fmt.Sprintf("\t位置: %d, PropId: %d(不存在)", pos, 0))
			continue
		}
		PropId := user.UserGear[pos].PropId
		msg = append(msg, fmt.Sprintf("\t位置: %d, PropId: %d - %s", pos, PropId, propAll[PropId].Type.String()))
	}

	fmt.Printf("%s\n", strings.Join(msg, "\n"))
}
