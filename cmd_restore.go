package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
)

type RestoreCommand struct {
	Bucket       string `long:"bucket" description:"Bucket which holds the backup file." required:"true"`
	Key          string `long:"key" description:"Key of the backup file." required:"true"`
	DatabaseName string `long:"dbname" description:"Name of the target database." required:"true"`
}

var restoreCommand RestoreCommand

func init() {
	parser.AddCommand(
		"restore",
		"Restore from S3",
		"Restore from S3.",
		&restoreCommand)
}

func (c *RestoreCommand) Execute(args []string) error {
	db, err := Connect(&connectionInfo)

	if err != nil {
		return err
	}

	row := db.QueryRow("msdb.dbo.rds_restore_database",
		sql.Named("restore_db_name", c.DatabaseName),
		sql.Named("s3_arn_to_restore_from", fmt.Sprintf("arn:aws:s3:::%s/%s", c.Bucket, c.Key)),
	)

	values := make([]interface{}, 11)

	for i := 0; i < len(values); i++ {
		values[i] = new(interface{})
	}

	err = row.Scan(values...)

	if err != nil {
		return err
	}

	status := TaskStatus{
		Id: (*values[0].(*interface{})).(int64),
	}

	r, _ := json.Marshal(status)
	fmt.Fprintf(os.Stdout, "%v\n", string(r))

	return nil
}
