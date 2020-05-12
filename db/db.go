package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/name1e5s/acdc/config"
	"github.com/name1e5s/acdc/model"
	"log"
	"sync"
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

	// Add our root user
	db.Create(&model.Admin{
		UserName: config.GetConfig().RootUser.UserName,
		Password: config.GetConfig().RootUser.Password,
		Role:     model.SuperUserMask,
	})
}

func CloseDataBase() {
	_ = database.Close()
}
