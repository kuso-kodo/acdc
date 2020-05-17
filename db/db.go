package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/name1e5s/acdc/config"
	"github.com/name1e5s/acdc/model"
	"log"
	"sync"
	"time"
)

var (
	database *gorm.DB
	once     sync.Once
)

func GetDataBase() *gorm.DB {
	once.Do(func() {
		var err error
		database, err = gorm.Open("postgres", config.GetConfig().Postgres.GetConnectionString())
		if err != nil {
			log.Println("ERR: Postgres DB open failed.")
			log.Panicln(err)
		}
		autoMigrate(database)
	})
	return database
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Room{})
	db.AutoMigrate(&model.Ticket{})

	// Add our root user
	db.FirstOrCreate(&model.Admin{
		UserName: config.GetConfig().RootUser.UserName,
		Password: config.GetConfig().RootUser.Password,
		Role:     model.SuperUserMask,
	})
	// Add placeholders
	db.FirstOrCreate(&model.User{
		UserName: "root",
		Password: "----",
		Phone:    "0000",
	})
	db.FirstOrCreate(&model.Room{
		RoomName:           "root",
		IsPowerOn:          false,
		IsServicing:        false,
		CurrentTemperature: 0,
		TargetTemperature:  0,
		FanSpeed:           0,
		LastOnTime:         time.Time{},
	})
	db.FirstOrCreate(&model.Ticket{
		StartAt:      time.Now(),
		EndAt:        time.Now(),
		ServiceCount: 0,
		FanSpeed:     0,
		TotalFee:     0,
		RoomRefer:    1,
		UserRefer:    1,
	})
}

func CloseDataBase() {
	_ = database.Close()
}
