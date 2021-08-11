package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"./xsorm"
)


type TableName struct {
	Column1 int `gorm:"column:column_1" json:"column_1"`
}

func main() {
	xsorm.AddConnect(xsorm.XConfig{
		Config:       mysql.Config{
			User:                    "root",
			Passwd:                  "1234567",
			Addr:                    "127.0.0.1",
			DBName:                  "cdn",
		},
		IsWrite:      true,
		Debug:        true,
	})

	tb := new(TableName)
	xsorm.NewBuild(tb).Where("column_1", 2).First()

	fmt.Println(tb.Column1)
}
