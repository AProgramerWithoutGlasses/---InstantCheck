package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAnalyzeHandler_EmptyContent(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	h := &AnalyzeHandler{}
	r.POST("/api/analyze", h.Handle)

	body := `{"type":"text","content":""}`
	req := httptest.NewRequest("POST", "/api/analyze", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestAnalyzeHandler_InvalidType(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	h := &AnalyzeHandler{}
	r.POST("/api/analyze", h.Handle)

	body := `{"type":"pdf","content":"some text"}`
	req := httptest.NewRequest("POST", "/api/analyze", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp["error"] == "" {
		t.Error("expected error message in response")
	}
}
