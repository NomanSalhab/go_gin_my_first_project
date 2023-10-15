package entity

import "time"

type Complaint struct {
	ID                    int       `json:"id"`
	UserID                int       `json:"user_id" binding:"required"`
	Text                  string    `json:"text" binding:"required"`
	Date                  time.Time `json:"date"`
	AboutDelivery         bool      `json:"about_delivery"`
	AboutTheApp           bool      `json:"about_the_app"`
	ImprovementSuggestion bool      `json:"improvement_suggestion"`
	OtherReason           bool      `json:"other_reason"`
}
