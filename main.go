package main

import (
	"Authorization/handler"
	"Authorization/model"
	"Authorization/storage"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"strconv"
)

func main() {
	router := gin.Default()

	var conf model.Config

	byte, err := os.ReadFile("./configuration.json")
	if err != nil {
		fmt.Println("Error Read File:", err)
		return
	}

	err = json.Unmarshal(byte, &conf)
	if err != nil {
		fmt.Println("Error Unmarshal:", err)
		return
	}

	dataBase, err := sqlx.Open("mysql", conf.DataSourceName)
	if err != nil {
		panic(err)
		return
	}

	if dataBase == nil {
		fmt.Println("dB nil")
		panic(err)
		return
	}

	server := handler.Server{
		Storage: &storage.UserStorage{
			DataBase: dataBase,
		},
	}

	router.GET("", server.Hand)

	port := ":" + strconv.Itoa(conf.Port)

	err = router.Run(port)
	if err != nil {
		panic(err)
		return
	}
}
