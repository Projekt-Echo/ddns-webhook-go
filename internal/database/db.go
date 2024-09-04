package database

import (
	"github.com/yxlimo/xormzap"
	"gofiber-template/internal/logger"
	"xorm.io/xorm"
)
import _ "github.com/mattn/go-sqlite3"

func InitDatabase() {
	db, err := xorm.NewEngine("sqlite3", "./data.db")
	if err != nil {
		logger.Logger.Sugar().Fatalf("Error opening database: %v", err)
	}
	db.ShowSQL(true)
	db.SetLogger(xormzap.Logger(logger.Logger))
}
