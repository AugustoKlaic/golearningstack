package request

import "time"

type MessageRequest struct {
	Content  string    `json:"content" binding:"required"`
	DateTime time.Time `json:"dateTime" binding:"required"`
}
