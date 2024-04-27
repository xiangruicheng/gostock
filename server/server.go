package server

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"gostock/config"
	"log"
)

var RedisClient *redis.Client
var MysqlClient *sql.DB

func InitMysql() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Data.Mysql.Username,
		config.Data.Mysql.Password,
		config.Data.Mysql.Host,
		config.Data.Mysql.Port,
		config.Data.Mysql.Dbname)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Println("datainit MysqlClient fail")
		return
	}

	err = db.Ping()
	if err != nil {
		log.Println("datainit MysqlClient fail")
		return
	}
	// 设置连接池的最大空闲连接数
	db.SetMaxIdleConns(100)
	// 设置连接池的最大打开连接数
	db.SetMaxOpenConns(1000)

	MysqlClient = db
	log.Println("datainit MysqlClient success")
}

// InitRedis 初始化Redis客户端
func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",                   // no password set
		DB:       config.Data.Redis.Db, // use default DB
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Println("redis datainit fail")
		return
	}
	log.Println("datainit RedisClient success")
}
