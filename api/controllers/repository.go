package controllers

import (
	"ffcs/api/cache"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	Rdb cache.Redis
	DB  *gorm.DB
}

func (r Repository) ConnectAuthRoutes(app *fiber.App) {
	auth := app.Group("auth")
	auth.Use(r.AuthMiddleware())
	auth.Get("/ping", func(c *fiber.Ctx) error { return c.SendString("auth") })
	auth.Post("/signup", r.SignupUser)
	auth.Post("/login", r.LoginUser)
}

func (r Repository) ConnectAdminApiRoutes(app *fiber.App) {
	// Function will group all routes dedicated to admin user only ,
	// including any custom Middlewares applicable to the admin user

	admin := app.Group("/admin")

	admin.Use(r.AdminMiddleware())
	admin.Get("/ping", func(c *fiber.Ctx) error { return c.SendString("admin") })
	admin.Post("/create", r.AdminCreateSlot)
	admin.Get("/read", r.AdminReadallSloat)
	admin.Get("/read/:id", r.AdminReadoneSlot)
}

func (r Repository) ConnectApiRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Use(r.ApiMiddleware())
	api.Get("/ping", func(c *fiber.Ctx) error { return c.SendString("api") })
	api.Get("/courses", r.ApiReadbyCourse)
	api.Post("/register/:id", r.RegisterCourse)
}
