package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelwp/goblog/api/controller"
)

func NewRouter(r *gin.Engine, config *controller.Config) {
	userController := controller.NewUserController(config)
	articleController := controller.NewArticleController(config)
	categoryController := controller.NewCategoryController(config)
	authorizationController := controller.NewAuthorizationController(config)

	r.GET("/ping", controller.HealthCheck)

	api := r.Group("/api")
	v1 := api.Group("/v1")

	User(v1, userController)
	Article(v1, articleController)
	Category(v1, categoryController)
	Authorization(v1, authorizationController)
}

func User(r *gin.RouterGroup, controller controller.UserController) {
	users := r.Group("/users")

	users.POST("/create", controller.CreateUser)
	users.GET("/", controller.GetUserList)
	users.PUT("/update", controller.UpdateUser)
	users.GET("/:id", controller.GetUser)
}

func Article(r *gin.RouterGroup, controller controller.ArticleController) {
	articles := r.Group("/articles")

	articles.POST("/create", controller.CreateArticle)
	articles.GET("/", controller.GetArticleList)
	articles.PUT("/update", controller.UpdateArticle)
	articles.GET("/:id", controller.GetArticle)
}

func Category(r *gin.RouterGroup, controller controller.CategoryController) {
	categories := r.Group("/categories")

	categories.POST("/create", controller.CreateCategory)
	categories.GET("/", controller.GetCategoryList)
	categories.PUT("/update", controller.UpdateCategory)
	categories.GET("/:id", controller.GetCategory)
}

func Authorization(r *gin.RouterGroup, controller controller.AuthorizationController) {
	auths := r.Group("/auths")

	auths.POST("/login", controller.Login)
	auths.GET("/logout", controller.Logout)
}
