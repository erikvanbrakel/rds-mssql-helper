package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jessevdk/go-flags"
	"net/url"
	"os"
)

func main() {
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}

type ConnectionInfo struct {
	Hostname string `long:"hostname" description:"The hostname of the MS SQL server." required:"true"`
	Port     int    `long:"port" description:"The port the MS SQL server is listening on." required:"true"`
	Username string `long:"username" description:"The username to authenticate to the SQL server." required:"true"`
	Password string `long:"password" description:"The password to authenticate to the SQL server." required:"true"`
}

var connectionInfo ConnectionInfo
var parser = flags.NewParser(&connectionInfo, flags.Default)

func Connect(c *ConnectionInfo) (*sql.DB, error) {
	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(c.Username, c.Password),
		Host:   fmt.Sprintf("%s:%d", c.Hostname, c.Port),
	}

	return sql.Open("sqlserver", u.String())
}
