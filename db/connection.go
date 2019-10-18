package db

import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Connection interface {
    Close() error
    Create(model interface{}) Connection
}

type GormConnection struct {
    db *gorm.DB
}

func Open(username string, password string, host string, port string, database string) (*GormConnection, error) {
    db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s", username, password, host, port, database))
    if err != nil {
	return nil, err
    }

    return &GormConnection{db: db}, nil
}

func (connection *GormConnection) Close() error {
    return connection.db.Close()
}

func (connection *GormConnection) Create(model interface{}) Connection {
    connection.db.Create(model)

    return connection
}
