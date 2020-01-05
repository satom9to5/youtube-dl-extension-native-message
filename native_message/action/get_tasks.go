package action

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mitchellh/mapstructure"

	ydqueue "github.com/satom9to5/youtube-dl-queue"
)

type getTaskParameter struct {
	SqlitePath string `mapstructure:"sqlite_path"`
}

func GetTasks(data interface{}) (interface{}, error) {
	gtp := &getTaskParameter{}

	err := mapstructure.Decode(data, gtp)

	if err != nil {
		return nil, err
	}
	if gtp.SqlitePath == "" {
		return nil, errors.New("SqlitePath is empty!")
	}

	return gtp.getAllTasks()
}

func (gtp *getTaskParameter) getAllTasks() (tasks []ydqueue.Task, err error) {
	db, err := sql.Open("sqlite3", gtp.SqlitePath)
	if err != nil {
		return nil, err
	}

	if err = ydqueue.InitializeSchema(db); err != nil {
		return nil, err
	}

	return ydqueue.GetAllTasks()
}
