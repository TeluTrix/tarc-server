package service

import (
	"fmt"
	"log/slog"

	"github.com/TeluTrix/tarc-server/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(conInfo model.Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		conInfo.Database.User,
		conInfo.Database.Pwd,
		conInfo.Database.Host,
		conInfo.Database.Port,
		conInfo.Database.Name,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}
