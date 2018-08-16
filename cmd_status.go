package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
)

type StatusCommand struct {
	TaskId int64 `long:"id" description:"Task ID to check." required:"true"`
}

var statusCommand StatusCommand

func init() {
	parser.AddCommand(
		"status",
		"Retrieve task status",
		"Retrieve task status.",
		&statusCommand)
}

func (c *StatusCommand) Execute(args []string) error {

	db, err := Connect(&connectionInfo)

	if err != nil {
		return err
	}

	result := db.QueryRow("msdb.dbo.rds_task_status",
		sql.Named("task_id", c.TaskId),
	)

	values := make([]interface{}, 12)

	for i := 0; i < len(values); i++ {
		values[i] = new(interface{})
	}

	err = result.Scan(values...)

	if err != nil {
		return err
	}

	info := ""
	if values[6].(*interface{}) != nil {
		info = (*values[6].(*interface{})).(string)
	}

	status := TaskStatus{
		Id:         c.TaskId,
		Info:       info,
		Status:     (*values[5].(*interface{})).(string),
		Percentage: (*values[3].(*interface{})).(int64),
	}

	r, _ := json.Marshal(status)
	fmt.Fprintf(os.Stdout, "%v\n", string(r))

	return nil
}

type TaskStatus struct {
	Id         int64  `json:"id"`
	Info       string `json:"info"`
	Status     string `json:"status"`
	Percentage int64  `json:"percentage"`
}
