package main

import (
	"fiber-gorm-rest-api/address"
	"fiber-gorm-rest-api/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func initDb() {

	database.DB, err = gorm.Open(mysql.Open(database.Dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

}

// Auto Migrate Table for Address
// func migrate() {
// 	err = database.DB.AutoMigrate(&address.Address{})

// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

func main() {
	app := fiber.New()
	var host string = "localhost:4000"

	// Connect to Database
	initDb()

	// Create table for Address
	// migrate()

	// Enable CORS
	app.Use(cors.New())

	// Routes
	app.Get("/addresses", address.GetAll)
	app.Get("/address/:id", address.Get)
	app.Post("/address", address.Create)
	app.Put("/address", address.Update)
	app.Delete("/address/:id", address.Delete)

	app.Listen(host)
}
