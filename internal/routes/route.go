package routes

import (
	"log"
	"net/http"
	"note/internal/auth"
	"note/internal/handlers"
	"note/internal/repository"
	"note/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, authConfig *auth.Auth) *gin.Engine {

	r := gin.Default()

	r.Use(auth.EnableCORS())

	trustedProxies := []string{"192.168.0.1", "10.0.0.1"}
	if err := r.SetTrustedProxies(trustedProxies); err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/test/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "Hello %s", name)
	})

	//user logic
	userRepo := repository.NewGormUserRepository(db)
	authService := service.NewAuthService(userRepo, authConfig)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)
	//todo logic
	todoRepo := repository.NewGormTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	r.POST("/users", userHandler.RegisterUser)
	r.POST("/login", authHandler.Login)
	r.POST("/logout", authHandler.Logout)

	authGroup := r.Group("/auth")
	authGroup.Use(auth.AuthMiddleware(userRepo))
	{
		authGroup.GET(`/todos`, todoHandler.GetAllTodos)
		authGroup.POST(`/todos`, todoHandler.CreateTodo)
		authGroup.GET(`/todos/:id`, todoHandler.FindTodo)
		authGroup.PATCH(`/todos/:id/complete`, todoHandler.MarkTodoComplete)

		authGroup.GET(`/users/:id`, userHandler.GetUser)
		authGroup.PUT(`/users/:id`, userHandler.UpdateUser)
		authGroup.DELETE(`/users/:id`, userHandler.DeleteUser)
	}

	return r
}
