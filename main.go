package main

import (
	"ffcs/api/cache"
	"ffcs/api/controllers"
	"ffcs/api/db"
	"ffcs/api/migrations"
	"ffcs/api/utils"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	utils.ImportEnvs() // importing all envionment varibles

	if viper.GetBool("MIGRATE") {
		migrations.Migrate()
	}

	// creating and connecting all database instance's

	r := controllers.Repository{
		Rdb: cache.Redis{
			Auth:  cache.CreateRedisInstance(0),
			Cache: cache.CreateRedisInstance(1),
		},
		DB: db.Connect(),
	}

	// creating fiber instance

	app := fiber.New()

	// connecting api routes
	r.ConnectAuthRoutes(app)
	r.ConnectAdminApiRoutes(app)
	r.ConnectApiRoutes(app)

	// // // creating a temporary admin role
	// r.DB.Create(&models.Users{
	// 	Rollno:    "OOOO",
	// 	Firstname: "Admin",
	// 	Lastname:  "Admin",
	// 	Email:     "admin@admin.com",
	// 	Branch:    "NIL",
	// 	IsAdmin:   true,
	// 	Password:  "adminpassword",
	// })
	// listerning at port

	log.Fatal(app.Listen(fmt.Sprintf(":%s", viper.GetString("PORT"))))
}
