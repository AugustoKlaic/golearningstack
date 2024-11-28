package response

import "time"

/*
 - Defining the object response in its struct
*/

type Message struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	DateTime time.Time `json:"dateTime"`
}
