package repository

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/mysql/mysql"
	"github.com/getsentry/sentry-go"
)

var (
	clientDB               *sql.DB
	onceConnectToDB        sync.Once
	instanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME")
	dbPort                 = os.Getenv("DB_PORT")
	dbTCPHost              = os.Getenv("INSTANCE_HOST")
	dbName                 = os.Getenv("DB_NAME")
	dbUser                 = os.Getenv("DB_USER")
	dbPwd                  = os.Getenv("DB_PASS")
)

func init() {
	onceConnectToDB.Do(
		func() {
			mySqlClientDB, err := connectToDB()
			if err != nil {
				sentry.CaptureException(err)
			}
			clientDB = mySqlClientDB
		},
	)
}

func connectToDB() (*sql.DB, error) {
	_, err := mysql.RegisterDriver("cloudsql-mysql", cloudsqlconn.WithCredentialsFile("./key.json"))
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	connSting := fmt.Sprintf("%s:%s@cloudsql-mysql(%s)/%s", dbUser, dbPwd, instanceConnectionName, dbName)
	mySqlClientDB, err := sql.Open(
		"cloudsql-mysql",
		connSting,
	)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	if err = mySqlClientDB.Ping(); err != nil {
		sentry.CaptureException(err)
	}

	return mySqlClientDB, nil
}
