package action

import (
	"database/sql"
	"errors"

	"github.com/mitchellh/mapstructure"

	ydqueue "github.com/satom9to5/youtube-dl-queue"
)

type getFailedTaskParameter struct {
	SqlitePath string `mapstructure:"sqlite_path"`
}

func GetFailedTasks(data interface{}) (interface{}, error) {
	gftp := &getFailedTaskParameter{}

	err := mapstructure.Decode(data, gftp)

	if err != nil {
		return nil, err
	}
	if gftp.SqlitePath == "" {
		return nil, errors.New("SqlitePath is empty!")
	}

	return gftp.getFailedTasks()
}

func (gftp *getFailedTaskParameter) getFailedTasks() (failedTasks []ydqueue.FailedTask, err error) {
	db, err := sql.Open("sqlite3", gftp.SqlitePath)
	if err != nil {
		return nil, err
	}

	if err = ydqueue.InitializeSchema(db); err != nil {
		return nil, err
	}

	return ydqueue.GetAllFailedTasks()
}
