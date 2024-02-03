package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BodyParserAndValidate(v *validator.Validate, ctx *fiber.Ctx, body interface{}) error {
	if err := ctx.BodyParser(body); err != nil {
		return err
	}
	if err := ValidateStruct(v, body); err != nil {
		return err
	}
	return nil
}
