package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yamirghofran/todolist-go/internal/db"
	"github.com/yamirghofran/todolist-go/internal/handlers"
)

func main() {
	databaseURL := os.Getenv("DATANASE_URL")
	databaseURL = "postgres://yamirghofran0:@localhost:5432/todogo?sslmode=disable"

	dbPool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	defer dbPool.Close()

	queries := db.New(dbPool)

	// initialize database service
	todoService := db.NewTodoService(queries)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// initialize handlers
	todoHandler := handlers.NewTodoHandler(todoService)

	// set up gin router
	r := gin.Default()

	// TODO: CORS Middleware

	todoHandler.RegisterRoutes(r)

	// health check endpoint
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	// start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
