package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func main() {
	dsn := "host=localhost user=hibiscus password=S3cret dbname=todo port=5432 sslmode=disable TimeZone=Europe/Istanbul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DBConn = db

	app := fiber.New()

	HandleRoutes(app)

	app.Listen(":3000")

}
