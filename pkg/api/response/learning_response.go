package response

import "time"

type Message struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	DateTime time.Time `json:"dateTime"`
}
