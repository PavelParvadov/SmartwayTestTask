package http

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	svc "github.com/pavelParvadov/SmartwayTask/internal/services"
)

func writeError(c fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, svc.ErrSvcInvalidInput):
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": svc.ErrSvcInvalidInput.Error()})
	case errors.Is(err, svc.ErrSvcEmployeeNotFound),
		errors.Is(err, svc.ErrSvcCompanyNotFound),
		errors.Is(err, svc.ErrSvcDepartmentNotFound):
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	case errors.Is(err, svc.ErrSvcPassportExist):
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": svc.ErrSvcPassportExist.Error()})
	case errors.Is(err, svc.ErrSvcPhoneExist):
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": svc.ErrSvcPhoneExist.Error()})
	case errors.Is(err, svc.ErrSvcConflict):
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": svc.ErrSvcConflict.Error()})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal error"})
	}
}
