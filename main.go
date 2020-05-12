package main

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/api"
	v1 "github.com/name1e5s/acdc/api/admin"
	"github.com/name1e5s/acdc/config"
	"github.com/name1e5s/acdc/db"
	_ "github.com/name1e5s/acdc/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
