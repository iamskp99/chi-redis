package cache

import (
	"gitlab.com/pragmaticreviews/golang-mux-api/entity"
)

type PostCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}
