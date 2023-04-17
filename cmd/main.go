package main

import (
	"github.com/gin-gonic/gin"
	"github.com/magnojunior07/golang-crud/pkg/common/db"
	"github.com/magnojunior07/golang-crud/pkg/songs"
	"github.com/spf13/viper"
)
	
func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()
	
	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)
	
	router := gin.Default()
	dbHandler := db.Init(dbUrl)

	songs.RegisterRoutes(router, dbHandler)

	router.Run(port)

}