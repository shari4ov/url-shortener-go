package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/shari4ov/url-shortener-go/pkg/payload"
)

type Validation struct {
	Validator *validator.Validate
}

func (v *Validation) ValidateURLPayload(c *fiber.Ctx) error {
	body := new(payload.UrlPayload)
	c.BodyParser(&body)
	err := v.Validator.Struct(body)
	if err != nil {
		payloadError := payload.ErrorPayload{
			Message: "Invalid payload",
			Status:  fiber.StatusBadRequest,
		}
		return c.Status(fiber.StatusBadRequest).JSON(payloadError)
	}
	return c.Next()
}
