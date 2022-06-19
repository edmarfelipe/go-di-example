package main

import (
	"log"

	"github.com/edmarfelipe/go-di-example/controllers"
	"github.com/edmarfelipe/go-di-example/services"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

func main() {
	injector := do.New()

	do.Provide(injector, services.NewUserService)
	do.Provide(injector, controllers.NewUserController)

	userCtrl := do.MustInvoke[controllers.UserController](injector)

	app := fiber.New()

	app.Get("/users/:id", userCtrl.Get)
	app.Post("/users/", userCtrl.Create)
	app.Delete("/users/:id", userCtrl.Delete)

	log.Fatal(app.Listen(":3000"))
}
