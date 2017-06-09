package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLProxy struct {
	MySQL   *sql.DB
	address string
	open    int
	idle    int
}

func (this *MySQLProxy) MySQLProxy(address string, open int, idle int) error {
	var err error
	this.address = address
	this.open = open
	this.idle = idle
	this.MySQL, err = sql.Open("mysql", address)
	if err != nil {
		fmt.Println("MySQLProxy Connect Error: ", err)
		return err
	}
	this.MySQL.SetMaxOpenConns(open)
	this.MySQL.SetMaxIdleConns(idle)
	fmt.Println("MySQL Proxy Ready...")
	return err
}

func (this *MySQLProxy) Ping() error {
	err := this.MySQL.Ping()
	if err != nil {
		fmt.Println("MySQLProxy Ping Error: ", err)
		return err
	} else {
		fmt.Println("MySQLProxy Ping Success.")
	}
	return err
}

func (this *MySQLProxy) Close() error {
	var err error
	err = this.MySQL.Close()
	if err != nil {
		fmt.Println("MySQLProxy Close Error: ", err)
		return err
	}
	return err
}
