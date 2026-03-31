// backend/internal/model/model.go
package model

import "time"

type AnalyzeLog struct {
	ID             int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	InputType      string    `gorm:"type:varchar(10);not null" json:"input_type"`
	InputContent   string    `gorm:"type:text;not null" json:"input_content"`
	InputURL       string    `gorm:"type:varchar(2048)" json:"input_url"`
	KeyPointsJSON  string    `gorm:"type:text" json:"key_points_json"`
	QuestionsJSON  string    `gorm:"type:text" json:"questions_json"`
	KeyPointsCount int       `json:"key_points_count"`
	QuestionsCount int       `json:"questions_count"`
	TokenUsage     int       `json:"token_usage"`
	DurationMs     int       `json:"duration_ms"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type QuizResult struct {
	ID           int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	AnalyzeID    int64     `gorm:"not null" json:"analyze_id"`
	AnswersJSON  string    `gorm:"type:text;not null" json:"answers_json"`
	CorrectCount int       `gorm:"not null" json:"correct_count"`
	TotalCount   int       `gorm:"not null" json:"total_count"`
	Accuracy     float64   `gorm:"type:decimal(5,2);not null" json:"accuracy"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}
