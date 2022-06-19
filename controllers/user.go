package controllers

import (
	"github.com/edmarfelipe/go-di-example/models"
	"github.com/edmarfelipe/go-di-example/services"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

type UserController struct {
	UserService services.UserServiceInterface
}

func NewUserController(i *do.Injector) (UserController, error) {
	ctrl := UserController{
		UserService: do.MustInvoke[services.UserServiceInterface](i),
	}
	return ctrl, nil
}

func (ctrl UserController) Get(c *fiber.Ctx) error {
	idUser := c.Params("id")
	if idUser == "" {
		return c.SendStatus(400)
	}

	user, err := ctrl.UserService.Get(idUser)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(user)
}

func (ctrl UserController) Delete(c *fiber.Ctx) error {
	idUser := c.Params("id")
	if idUser == "" {
		return c.SendStatus(400)
	}

	err := ctrl.UserService.Delete(idUser)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(204)
}

func (ctrl UserController) Create(c *fiber.Ctx) error {
	model := models.User{}

	err := c.BodyParser(&model)
	if err != nil {
		return c.SendStatus(400)
	}

	err = ctrl.UserService.Create(model)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(204)
}
