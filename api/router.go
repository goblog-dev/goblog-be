package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelwp/goblog/api/controller"
	"github.com/michaelwp/goblog/middleware"
)

func NewRouter(r *gin.Engine, config *controller.Config) {
	userController := controller.NewUserController(config)
	articleController := controller.NewArticleController(config)
	categoryController := controller.NewCategoryController(config)
	authorizationController := controller.NewAuthorizationController(config)

	r.GET("/ping", controller.HealthCheck)

	api := r.Group("/api")
	v1 := api.Group("/v1")

	User(v1, userController, config)
	Article(v1, articleController, config)
	Category(v1, categoryController, config)
	Authorization(v1, authorizationController, config)
}

func User(r *gin.RouterGroup, controller controller.UserController, config *controller.Config) {
	users := r.Group("/users").Use(middleware.AuthMiddleware(config))
	{
		users.POST("/create", controller.CreateUser)
		users.GET("/", controller.GetUserList)
		users.PUT("/update", controller.UpdateUser)
		users.GET("/:id", controller.GetUser)
	}
}

func Article(r *gin.RouterGroup, controller controller.ArticleController, config *controller.Config) {
	articles := r.Group("/articles").Use(middleware.AuthMiddleware(config))
	{
		articles.POST("/create", controller.CreateArticle)
		articles.GET("/", controller.GetArticleList)
		articles.PUT("/update", controller.UpdateArticle)
		articles.GET("/:id", controller.GetArticle)
		articles.DELETE("/delete", controller.DeleteArticle)
	}
}

func Category(r *gin.RouterGroup, controller controller.CategoryController, config *controller.Config) {
	categories := r.Group("/categories").Use(middleware.AuthMiddleware(config))
	{
		categories.POST("/create", controller.CreateCategory)
		categories.GET("/", controller.GetCategoryList)
		categories.PUT("/update", controller.UpdateCategory)
		categories.GET("/:id", controller.GetCategory)
	}
}

func Authorization(r *gin.RouterGroup, controller controller.AuthorizationController, config *controller.Config) {
	auths := r.Group("/auths")
	{
		auths.POST("/login", controller.Login)
		auths.Use(middleware.AuthMiddleware(config))
		{
			auths.GET("/logout", controller.Logout)
		}
	}

}
