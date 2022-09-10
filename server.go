package main

import (
	"os"

	"gitlab.com/pragmaticreviews/golang-mux-api/cache"
	"gitlab.com/pragmaticreviews/golang-mux-api/controller"
	router "gitlab.com/pragmaticreviews/golang-mux-api/http"
	"gitlab.com/pragmaticreviews/golang-mux-api/repository"
	"gitlab.com/pragmaticreviews/golang-mux-api/service"
)

var (
	postRepository repository.PostRepository = repository.NewSQLiteRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postCache      cache.PostCache           = cache.NewRedisCache("localhost:6379", 1, 10)
	postController controller.PostController = controller.NewPostController(postService, postCache)
	httpRouter     router.Router             = router.NewChiRouter(os.Getenv("PORT"))
)

func main() {
	httpRouter.GET("/posts/{id}", postController.GetPostByID)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(os.Getenv("PORT"))
}
