package action

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mitchellh/mapstructure"

	ydqueue "github.com/satom9to5/youtube-dl-queue"
)

type getTasksByIdsParameter struct {
	SqlitePath string   `mapstructure:"sqlite_path"`
	Ids        []string `mapstructure:"ids"`
}

func GetTasksByIds(data interface{}) (interface{}, error) {
	gtbip := &getTasksByIdsParameter{}

	err := mapstructure.Decode(data, gtbip)

	if err != nil {
		return nil, err
	}
	if gtbip.SqlitePath == "" || len(gtbip.Ids) == 0 {
		return nil, errors.New("SqlitePath/Ids is empty!")
	}

	return gtbip.getTasksByIds()
}

func (gtbip *getTasksByIdsParameter) getTasksByIds() (tasks []ydqueue.Task, err error) {
	db, err := sql.Open("sqlite3", gtbip.SqlitePath)
	if err != nil {
		return nil, err
	}

	if err = ydqueue.InitializeSchema(db); err != nil {
		return nil, err
	}

	return ydqueue.GetTasksMapByIds(gtbip.Ids)
}
