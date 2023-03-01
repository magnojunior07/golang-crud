package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/magnojunior07/golang-crud/pkg/common/config"
	"github.com/magnojunior07/golang-crud/pkg/common/db"
	"github.com/spf13/viper"
)
	

func main() {
	config.LoadConfig()
	router := gin.Default()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	db.Init(dbUrl)
	router.Run(port)

	fmt.Println("Hello, world")
}