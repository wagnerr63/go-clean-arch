package configs

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetReaderGorm() *gorm.DB {
	var host string = os.Getenv("HOST")
	var port string = os.Getenv("PORT")
	var dbname string = os.Getenv("DBNAME")
	var user string = os.Getenv("DBUSER")
	var password string = os.Getenv("DBPASSWORD")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"
	reader, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return reader
}

func GetWriterGorm() *gorm.DB {
	var host string = os.Getenv("HOST")
	var port string = os.Getenv("PORT")
	var dbname string = os.Getenv("DBNAME")
	var user string = os.Getenv("DBUSER")
	var password string = os.Getenv("DBPASSWORD")

	fmt.Println("user " + user)

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"
	writer, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return writer
}
