package db

import (
	"encoding/json"
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
	db.Create(&model.Admin{
		UserName: config.GetConfig().RootUser.UserName,
		Password: config.GetConfig().RootUser.Password,
		Role:     model.SuperUserMask,
	})
	// Add placeholders
	db.Create(&model.User{
		UserName: "root",
		Password: "----",
		Phone:    "0000",
	})
	db.Create(&model.Room{
		RoomName:           "root",
		IsPowerOn:          false,
		IsServicing:        false,
		CurrentTemperature: 0,
		TargetTemperature:  0,
		FanSpeed:           0,
		LastOnTime:         time.Time{},
	})
	db.Create(&model.Ticket{
		StartAt:      time.Now(),
		EndAt:        time.Now(),
		ServiceCount: 0,
		FanSpeed:     0,
		TotalFee:     0,
		RoomRefer:    1,
		UserRefer:    1,
	})
	var ticket []model.Ticket
	var cnt int
	db.Debug().Model(&model.Room{
		RoomID: 1,
	}).Association("Tickets").Find(&ticket)
	data, _ := json.Marshal(ticket)
	log.Println(string(data))
	log.Println(cnt)
}

func CloseDataBase() {
	_ = database.Close()
}
