package request

import "time"

type MessageRequest struct {
	Content  string    `json:"content"`
	DateTime time.Time `json:"dateTime"`
}
