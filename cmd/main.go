package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/magnojunior07/golang-crud/pkg/common/db"
	"github.com/spf13/viper"
)
	

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	router := gin.Default()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	db.Init(dbUrl)
	router.Run(port)

	fmt.Println("Hello, world")
}