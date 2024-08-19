package main

import (
	"github.com/gin-gonic/gin"
	handlers "yalo-api/internal/handlers/users_interactions"
	"yalo-api/internal/middlewares"
	"yalo-api/internal/resources/repository"
	"yalo-api/internal/services"
)

func main() {
	db := repository.NewUserInteractionDB()
	userInteractionService := services.NewUserInteractionService(db)

	ginEngine := gin.Default()
	ginEngine.Use(middlewares.ZerologMiddleware())

	ginEngine.POST("/user_interaction", handlers.NewUserInteractionHandler(userInteractionService).SaveUserInteraction)
	ginEngine.GET("/user_interaction/:user_id", handlers.NewUserInteractionHandler(userInteractionService).GetUserInteraction)

	err := ginEngine.Run(":8080")
	if err != nil {
		panic(err)
	}
}
