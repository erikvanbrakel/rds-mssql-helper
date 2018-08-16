# rds-mssql-helper

This tool assists with restoring backups from S3 into Microsoft SQL Server in
AWS RDS. To import an SQL backup (the typical .bak file), you need to connect
to the database server itself and execute a stored procedure. This CLI tool will
do this for you, as long as you provide the correct parameters. As a finishing
touch it returns JSON objects, so you can parse the contents and use it in your
automation.

## How to use

### Restoring a backup

```
rds-mssql-helper restore --hostname=your-sql-server --port=1433 --username=user --password=pwd --bucket=s3-bucket-name --key=object-key --dbname=database-to-restore-into
```

This returns:

```json
{
    "id" : "1234"
}
```


### Retrieve task status
```
rds-mssql-helper restore --hostname=your-sql-server --port=1433 --username=user --password=pwd --id=1234
```

This returns:

```json
{
    "id" : "1234",
    "info" : "output messages go here",
    "status": "CREATED | IN_PROGRESS | SUCCESS | ERROR | CANCEL_REQUESTED | CANCELLED",
    "percentage" : "0-100"
}
```
