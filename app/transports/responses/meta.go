package responses

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type SucessResp struct {
	Meta base        `json:"meta"`
	Data interface{} `json:"data"`
}

type MetaResp struct {
	Meta base `json:"meta"`
}

type base struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func Json(ctx *fiber.Ctx, status int, data interface{}) error {
	var ress interface{}

	meta := base{
		Message: http.StatusText(status),
		Status:  status,
	}

	if data != nil {
		if str, ok := data.(string); ok && str != "" {
			meta.Message = str
			ress = MetaResp{Meta: meta}
		} else {
			ress = SucessResp{Meta: meta, Data: data}
		}
	} else {
		ress = MetaResp{Meta: meta}
	}

	if err := ctx.Status(status).JSON(ress); err != nil {
		fmt.Println("Error sending JSON response:", err.Error())
		return Json(ctx, fiber.StatusInternalServerError, nil)
	}

	return nil
}
