package main

import (
	"fmt"
	"log"
	"quiz1/model"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

const ERROR_CODE = 99999

var Val *Config

type Config struct {
	ServerPort          string `mapstructure:"SERVER_PORT" json:"SERVER_PORT"`
	MysqlHost           string `mapstructure:"MYSQL_HOST" json:"MYSQL_HOST"`
	MysqlPort           int    `mapstructure:"MYSQL_PORT" json:"MYSQL_PORT"`
	MysqlUser           string `mapstructure:"MYSQL_USER" json:"MYSQL_USER"`
	MysqlPassword       string `mapstructure:"MYSQL_PASSWORD" json:"MYSQL_PASSWORD"`
	MysqlMaxidle        int    `mapstructure:"MYSQL_MAXIDLE" json:"MYSQL_MAXIDLE"`
	MysqlMaxconn        int    `mapstructure:"MYSQL_MAXCONN" json:"MYSQL_MAXCONN"`
	MysqlConnMaxLifeTim int    `mapstructure:"MYSQL_CONNMAXLIFETTIME" json:"MYSQL_CONNMAXLIFETTIME"`
	MysqlSingularTable  bool   `mapstructure:"MYSQL_SINGULARTABLE" json:"MYSQL_SINGULARTABLE"`
	MysqlLogMode        bool   `mapstructure:"MYSQL_LOGMODE" json:"MYSQL_LOGMODE"`
}

func init() {
	//讀取還近變數
	viper.AutomaticEnv()
	//取設定黨
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := viper.Unmarshal(&Val)
	if err != nil {
		panic(err)
	}

	log.Println("ENV:", Val)
	log.Println("Cofing 設定成功")
	//建立mysql pool
	model.UserM = model.MysqlConn()
}

func main() {
	defer model.Close()
	httpServer()
}

func httpServer() {
	r := gin.Default()
	r.Use(gin.Recovery())

	r.GET("/ping", Pong)

	prefix := r.Group("/quiz/v1")
	prefix.POST("/insert", CreateUser)
	prefix.GET("/", GetUser)
	prefix.DELETE("/:id", DeleteUser)
	prefix.PUT("/update", UpdateUser)
	r.Run(string(fmt.Sprintf("%v", viper.Get("SERVER_PORT"))))
}
