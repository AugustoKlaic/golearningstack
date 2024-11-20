package api

import "time"

/*
 - Defining the object response in its struct
 - I do not have a persistence layer to save and get data, so I created a variable to mock the response
*/

type Message struct {
	Content  string    `json:"content"`
	DateTime time.Time `json:"dateTime"`
}

var Messages = []Message{
	{Content: "Hello World!", DateTime: time.Now()},
}
