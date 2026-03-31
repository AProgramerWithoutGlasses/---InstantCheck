package llm

import (
	"encoding/json"
	"testing"
)

func TestParseAnalyzeResponse(t *testing.T) {
	raw := `{
		"key_points": [
			{"title": "Point 1", "description": "Description 1"},
			{"title": "Point 2", "description": "Description 2"}
		],
		"questions": [
			{
				"id": 1,
				"question": "What is point 1?",
				"options": ["A", "B", "C", "D"],
				"correct_answer": 0,
				"explanation": "Because..."
			}
		]
	}`

	var result AnalyzeResult
	err := json.Unmarshal([]byte(raw), &result)
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	if len(result.KeyPoints) != 2 {
		t.Errorf("expected 2 key points, got %d", len(result.KeyPoints))
	}
	if len(result.Questions) != 1 {
		t.Errorf("expected 1 question, got %d", len(result.Questions))
	}
	if result.Questions[0].CorrectAnswer != 0 {
		t.Errorf("expected correct_answer 0, got %d", result.Questions[0].CorrectAnswer)
	}
}
