package main

import "errors"

type BackupCommand struct{}

var backupCommand BackupCommand

func init() {
	parser.AddCommand(
		"backup",
		"Backup to S3",
		"Backup to S3.",
		&backupCommand)
}

func (c *BackupCommand) Execute(args []string) error {
	return errors.New("not implemented")
}
