package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type student struct {
	gorm.Model
	Name  string
	Email string `gorm:"unique;not null"`
}

var gDb *gorm.DB

func main() {
	var err error
	var newLogger = logger.New(
		log.New(os.Stdout, "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		},
	)

	dsn := "host=localhost user=diwakar password=root dbname=postgres port=5432 sslmode=disable"
	gDb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}
	conn, err := gDb.DB()

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Println("connected")
	CreateTable()
	InsertData()
	SearchAll()
	SearchWhere()
	//Delete()
}

func CreateTable() {

	//err := gDb.Migrator().DropTable(&student{})
	//if err != nil {
	//	log.Fatalln(err)
	//}

	err := gDb.Migrator().AutoMigrate(&student{})
	if err != nil {
		log.Fatalln(err)
	}

}

func InsertData() {
	s := []student{
		{
			Name:  "diwakar",
			Email: "diwakar@email.com",
		},
		{
			Name:  "Raj",
			Email: "raj@email.com",
		},
		{
			Name:  "dev",
			Email: "dev@email.com",
		},
	}

	err := gDb.Create(&s)
	if err != nil {
		log.Println(err)
		return
	}
}

func SearchAll() {

	var s []student

	err := gDb.Find(&s).Error

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(s)

}

func SearchWhere() {

	var s student
	name := "Raj"
	tx := gDb.Where("name = ?", name)
	err := tx.First(&s).Error
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(s)
}
func Delete() {

	var s student
	gDb.First(&s)
	//err := gDb.Delete(&s).Error //soft delete
	err := gDb.Unscoped().Delete(&s).Error
	if err != nil {
		log.Println(err)
		return
	}
}
