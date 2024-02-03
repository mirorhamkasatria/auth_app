package controllers

import (
	"github.com/auth_app/app/models"
	"github.com/auth_app/app/services"
	"github.com/auth_app/app/transports/requests"
	"github.com/auth_app/app/transports/responses"
	"github.com/auth_app/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type controllerUser struct {
	sv        services.ServiceUser
	validator *validator.Validate
}

func NewControllerUser(service services.ServiceUser, validator *validator.Validate) *controllerUser {
	return &controllerUser{
		sv:        service,
		validator: validator,
	}
}

func (c *controllerUser) LoginUser(ctx *fiber.Ctx) error {
	body := new(requests.Login)
	if err := utils.BodyParserAndValidate(c.validator, ctx, body); err != nil {
		return responses.Json(ctx, fiber.StatusBadRequest, err.Error())
	}

	data := models.User{}
	if err := utils.Copier(&data, body); err != nil {
		return responses.Json(ctx, fiber.StatusUnprocessableEntity, err.Error())
	}

	token, err := c.sv.Login(&data)
	if err != nil {
		return responses.Json(ctx, fiber.StatusUnprocessableEntity, err.Error())
	}

	return responses.Json(ctx, fiber.StatusOK, responses.UserToken{Token: *token})
}

func (c *controllerUser) RegisterUser(ctx *fiber.Ctx) error {
	body := new(requests.Register)
	if err := utils.BodyParserAndValidate(c.validator, ctx, body); err != nil {
		return responses.Json(ctx, fiber.StatusBadRequest, err.Error())
	}

	data := models.User{}
	if err := utils.Copier(&data, body); err != nil {
		return responses.Json(ctx, fiber.StatusUnprocessableEntity, err.Error())
	}

	token, err := c.sv.RegisterUser(&data)
	if err != nil {
		return responses.Json(ctx, fiber.StatusUnprocessableEntity, err.Error())
	}

	if token == nil {
		return responses.Json(ctx, fiber.StatusInternalServerError, "Registration failed")
	}

	return responses.Json(ctx, fiber.StatusOK, responses.UserToken{Token: *token})
}
