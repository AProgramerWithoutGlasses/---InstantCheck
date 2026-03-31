package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestQuizResultHandler_EmptyAnswers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	h := &QuizResultHandler{}
	r.POST("/api/quiz-result", h.Handle)

	body := `{"analyze_id":1,"answers":[]}`
	req := httptest.NewRequest("POST", "/api/quiz-result", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}
