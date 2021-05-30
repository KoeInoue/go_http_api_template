package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"go_http_api_template/infrastructure"
	"time"
)

var (
	db  *gorm.DB
	err error
)

type User struct {
	Id        int    `gorm:"primary_key"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"size:255"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
}

func init() {
	infrastructure.Init()
	db = infrastructure.Db()
	fmt.Println(db)
}

// Up is executed when this migration is applied
func Up_20210530144301(txn *sql.Tx) {
	// Create table for `User`
	db.CreateTable(&User{})
}

// Down is executed when this migration is rolled back
func Down_20210530144301(txn *sql.Tx) {
	db.DropTable(&User{})
}