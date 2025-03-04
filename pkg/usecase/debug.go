package usecase

import (
	"api/pkg/model"
	"context"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

func NewDebug() Debuger {
	return &_debug{}
}

type Debuger interface {
	// 用户信息概览
	DebugUserInfo(ctx context.Context, user *model.UserEntity)
	// 给道具表 增加一个经验值道具
	DebugAddExpProp(ctx context.Context)
	// 打印用户任务信息
	DebugUserTaskInfo(ctx context.Context, user *model.UserEntity)
}

type _debug struct {
}

// DebugAddExpProp implements Debuger.
func (d *_debug) DebugAddExpProp(ctx context.Context) {
	db.Transaction(func(tx *gorm.DB) error {
		log.Debug("给道具表 增加一个经验值道具")
		prop := model.NewProper(model.PropTypeExperience, model.NewExperiencer(1000))
		prop.SetStartTs(time.Now().Unix())
		prop.SetEndTs(time.Now().Add(time.Hour * 24 * 365 * 10).Unix())

		propRepo.Create(ctx, tx, prop)

		prop = propRepo.Get(ctx, tx, prop.GetPropId())
		log.Infof("添加成功, PropId: %d, Content: %s", prop.GetPropId(), prop.GetContent())
		log.Infof("data: %v", GetJsonString(prop.ConvertExperiencer()))

		return nil
	})
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
	msg = append(msg, "用户装备:")

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

func (e *_debug) DebugUserTaskInfo(ctx context.Context, user *model.UserEntity) {

	var msg = make([]string, 0, 100)

	list := userTaskRepo.DebugGetUserTaskList(ctx, db, user)
	msg = append(msg, fmt.Sprintf("全部任务数据: %d", len(list)))
	for _, userTask := range list {
		msg = append(msg, fmt.Sprintf("\tuserTask.Id: %d, 任务Id: %d, 状态: %s", userTask.Id, userTask.TaskId, userTask.Status.Summary()))
	}

	fmt.Printf("%s\n", strings.Join(msg, "\n"))
}
