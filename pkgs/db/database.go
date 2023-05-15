package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nann-e-backend/config"
)

type Connection struct {
	MYSQL config.Config
}

func NewDatabaseConnection(M config.Config) *Connection {
	return &Connection{
		MYSQL: M,
	}
}

func (db *Connection) DBConnect() *sql.DB {
	dbConn, errConn := sql.Open(
		"mysql", db.MYSQL.Database.Username+":"+db.MYSQL.Database.Password+"@tcp("+db.MYSQL.Database.Host+")/"+db.MYSQL.Database.Database,
	)
	fmt.Println(errConn)
	if errConn != nil {
		return nil
	}
	for dbConn.Ping() != nil {
		fmt.Println("Retrying...")
		time.Sleep(5 * time.Second)
	}
	dbConn.SetMaxIdleConns(db.MYSQL.Database.MaxIdleConn)
	dbConn.SetMaxOpenConns(db.MYSQL.Database.MaxOpenConn)
	return dbConn
}
