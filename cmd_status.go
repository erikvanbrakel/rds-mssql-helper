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

	status := TaskStatus{
		Id:         c.TaskId,
		Info:       safeGetString(values[6].(*interface{})),
		Status:     safeGetString(values[5].(*interface{})),
		Percentage: safeGetInt(values[3].(*interface{})),
	}

	r, _ := json.Marshal(status)
	fmt.Fprintf(os.Stdout, "%v\n", string(r))

	return nil
}

func safeGetString(i *interface{}) string {
	if *i == nil {
		return ""
	} else {
		return (*i).(string)
	}
}

func safeGetInt(i *interface{}) int64 {
	if *i == nil {
		return 0
	} else {
		return (*i).(int64)
	}
}

type TaskStatus struct {
	Id         int64  `json:"id"`
	Info       string `json:"info"`
	Status     string `json:"status"`
	Percentage int64  `json:"percentage"`
}
