package cache

import (
	"context"

	"gitlab.com/pragmaticreviews/golang-mux-api/entity"
)

type PostCache interface {
	Set(key string, value *entity.Post)
	Get(ctx context.Context, key string) *entity.Post
}
