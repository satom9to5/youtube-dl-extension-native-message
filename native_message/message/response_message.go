package message

import "fmt"

type ResponseMessage struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func NewResponseMessage() *ResponseMessage {
	return &ResponseMessage{}
}

func (rm ResponseMessage) String() string {
	return fmt.Sprintf("Data:\"%s\", Error: \"%s\"", rm.Data, rm.Error)
}
