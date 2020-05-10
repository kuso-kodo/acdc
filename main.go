package main

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/api"
	v1 "github.com/name1e5s/acdc/api/v1"
	"github.com/name1e5s/acdc/config"
	"github.com/name1e5s/acdc/db"
	"log"
)

func main() {
	log.Println(config.GetConfig().Title)
	log.Println(config.GetConfig().Postgres.GetConnectionString())
	log.Println(config.GetConfig().RootUser)
	_ = db.GetDataBase()
	defer db.CloseDataBase()

	r := gin.Default()
	r.Group("a").Group("b").GET("/test", v1.GetAllUsers)
	apiRouter := r.Group("api")
	api.BindAPIRouters(apiRouter)
	r.Run()
}
