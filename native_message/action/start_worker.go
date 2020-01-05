package action

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

func StartWorker(data interface{}) (interface{}, error) {
	w := &worker{}

	err := mapstructure.Decode(data, w)

	if err != nil {
		return nil, err
	}
	if w.SqlitePath == "" || w.PidfilePath == "" || w.YoutubeDlPath == "" || w.LogDirectory == "" {
		return nil, errors.New("SqlitePath/PidfilePath/YoutubeDlPath/LogDirectory is empty!")
	}

	return w.Start()
}
