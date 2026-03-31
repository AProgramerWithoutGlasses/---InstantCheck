package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/AProgramerWithoutGlasses/instant-check/backend/internal/config"
	"github.com/AProgramerWithoutGlasses/instant-check/backend/internal/database"
	"github.com/AProgramerWithoutGlasses/instant-check/backend/internal/handler"
	"github.com/AProgramerWithoutGlasses/instant-check/backend/internal/llm"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	llmClient := llm.NewClient(cfg.LLM)

	analyzeHandler := &handler.AnalyzeHandler{DB: db, LLMClient: llmClient}
	quizResultHandler := &handler.QuizResultHandler{DB: db}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
	}))

	api := r.Group("/api")
	{
		api.POST("/analyze", analyzeHandler.Handle)
		api.POST("/quiz-result", quizResultHandler.Handle)
	}

	log.Printf("Server starting on :%s", cfg.Server.Port)
	log.Fatal(r.Run(":" + cfg.Server.Port))
}
