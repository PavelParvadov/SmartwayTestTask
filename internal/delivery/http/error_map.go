package http

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	svc "github.com/pavelParvadov/SmartwayTask/internal/services"
)

func writeError(c fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, svc.ErrSvcEmployeeNotFound),
		errors.Is(err, svc.ErrSvcCompanyNotFound),
		errors.Is(err, svc.ErrSvcDepartmentNotFound):
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	case errors.Is(err, svc.ErrSvcPassportExist),
		errors.Is(err, svc.ErrSvcPhoneExist):
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal error"})
	}
}
