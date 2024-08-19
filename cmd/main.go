package main

import (
	handlers "yalo-api/pkg/handlers/users_interactions"
	"yalo-api/pkg/resources/repository"
	"yalo-api/pkg/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := repository.NewUserInteractionDB()
	userInteractionService := services.NewUserInteractionService(db)

	r := gin.Default()
	r.POST("/user_interaction", handlers.NewUserInteractionHandler(userInteractionService).SaveUserInteraction)
	r.GET("/user_interaction/:user_id", handlers.NewUserInteractionHandler(userInteractionService).GetUserInteraction)

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(err)
	}
}
