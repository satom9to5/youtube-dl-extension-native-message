package action

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mitchellh/mapstructure"

	ydqueue "github.com/satom9to5/youtube-dl-queue"
)

type queueParameter struct {
	SqlitePath  string `mapstructure:"sqlite_path"`
	Url         string `mapstructure:"url"`
	Title       string `mapstructure:"title"`
	VideoFormat string `mapstructure:"video_format"`
	AudioFormat string `mapstructure:"audio_format"`
	OutputPath  string `mapstructure:"output_path"`
	Parameter   string `mapstructure:"parameter"`
}

func AddQueue(data interface{}) (interface{}, error) {
	qp := &queueParameter{}

	err := mapstructure.Decode(data, qp)

	if err != nil {
		return nil, err
	}
	if qp.Url == "" || qp.Title == "" || qp.VideoFormat == "" || qp.OutputPath == "" {
		return nil, errors.New("least one of Url/Title/VideoFormat/OutputPath is empty!")
	}

	return qp.add()
}

func (qp *queueParameter) add() (task *ydqueue.Task, err error) {
	db, err := sql.Open("sqlite3", qp.SqlitePath)
	if err != nil {
		return nil, err
	}

	if err = ydqueue.InitializeSchema(db); err != nil {
		return nil, err
	}

	task = &ydqueue.Task{
		Url:         qp.Url,
		Title:       qp.Title,
		VideoFormat: qp.VideoFormat,
		AudioFormat: qp.AudioFormat,
		OutputPath:  qp.OutputPath,
		Parameter:   qp.Parameter,
	}

	return task, task.QueueTask()
}
