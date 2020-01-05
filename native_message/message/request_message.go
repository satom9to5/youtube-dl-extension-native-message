package message

import (
	"fmt"
	"native_message/action"
)

const (
	StartWorker          = "startWorker"
	StopWorker           = "stopWorker"
	CheckRunningWorker   = "checkRunningWorker"
	AddQueue             = "addQueue"
	GetTasks             = "getTasks"
	GetFailedTasks       = "getFailedTasks"
	GetTasksByIds        = "getTasksByIds"
	CheckYoutubeDLUpdate = "checkYoutubeDLUpdate"
	CheckFFMpegUpdate    = "checkFFMpegUpdate"
)

type RequestMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func NewRequestMessage() *RequestMessage {
	return &RequestMessage{}
}

func (rm RequestMessage) String() string {
	return fmt.Sprintf("Type: %s\tData: [%s]", rm.Type, rm.Data)
}

func (rm *RequestMessage) Run() (interface{}, error) {
	//fmt.Fprintln(os.Stderr, rm.String())

	switch rm.Type {
	case StartWorker:
		return action.StartWorker(rm.Data)
	case StopWorker:
		return action.StopWorker(rm.Data)
	case CheckRunningWorker:
		return action.CheckRunningWorker(rm.Data)
	case AddQueue:
		return action.AddQueue(rm.Data)
	case GetTasks:
		return action.GetTasks(rm.Data)
	case GetFailedTasks:
		return action.GetFailedTasks(rm.Data)
	case GetTasksByIds:
		return action.GetTasksByIds(rm.Data)
	case CheckYoutubeDLUpdate:
		return nil, nil
	case CheckFFMpegUpdate:
		return nil, nil
	default:
		return nil, nil
	}
}
