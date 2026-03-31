package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/AProgramerWithoutGlasses/instant-check/backend/internal/model"
)

type QuizResultHandler struct {
	DB *gorm.DB
}

type QuizAnswer struct {
	QuestionID     int `json:"question_id"`
	SelectedAnswer int `json:"selected_answer"`
	CorrectAnswer  int `json:"correct_answer"`
}

type QuizResultRequest struct {
	AnalyzeID int64        `json:"analyze_id" binding:"required"`
	Answers   []QuizAnswer `json:"answers"`
}

func (h *QuizResultHandler) Handle(c *gin.Context) {
	var req QuizResultRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请提供作答结果"})
		return
	}

	if len(req.Answers) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "答案不能为空"})
		return
	}

	correctCount := 0
	for _, a := range req.Answers {
		if a.SelectedAnswer == a.CorrectAnswer {
			correctCount++
		}
	}

	total := len(req.Answers)
	accuracy := float64(correctCount) / float64(total) * 100

	answersJSON, _ := json.Marshal(req.Answers)

	result := model.QuizResult{
		AnalyzeID:    req.AnalyzeID,
		AnswersJSON:  string(answersJSON),
		CorrectCount: correctCount,
		TotalCount:   total,
		Accuracy:     accuracy,
	}

	if h.DB != nil {
		if dbResult := h.DB.Create(&result); dbResult.Error != nil {
			log.Printf("failed to save quiz result: %v", dbResult.Error)
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
