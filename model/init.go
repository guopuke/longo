package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

var DB *Database

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

func (db *Database) Init() {
	DB = &Database{
		Self:   GetSelfDB(),
		Docker: GetDockerDB(),
	}
}

func (db *Database) Close() {
	DB.Self.Close()
	DB.Docker.Close()
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func InitSelfDB() *gorm.DB {
	return openDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
	)
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

func InitDockerDB() *gorm.DB {
	return openDB(
		viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"),
	)
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	// db.DB().SetMaxOpenConns(20000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(0)
}

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s", username, password, addr, name, true, "Local")
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Databases connection failed. Database name: %s", name)
	}
	setupDB(db)

	return db
}
