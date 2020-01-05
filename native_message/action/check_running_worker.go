package action

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

func CheckRunningWorker(data interface{}) (interface{}, error) {
	w := &worker{}

	err := mapstructure.Decode(data, w)

	if err != nil {
		return nil, err
	}
	if w.PidfilePath == "" {
		return nil, errors.New("PidfilePath is empty!")
	}

	return w.CheckRunning()
}
