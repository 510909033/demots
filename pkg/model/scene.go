package model

import "context"

type SceneInterface interface {
	GetScene(ctx context.Context) string
}
