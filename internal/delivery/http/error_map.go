package http

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	repo "github.com/pavelParvadov/SmartwayTask/internal/repository/postgres"
	svc "github.com/pavelParvadov/SmartwayTask/internal/services"
)

func writeError(c fiber.Ctx, err error) error {
	if errors.Is(err, repo.ErrPassportExist) {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "passport already exists"})
	}
	if errors.Is(err, repo.ErrPhoneExist) {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "phone already exists"})
	}
	switch {
	case errors.Is(err, svc.ErrSvcInvalidInput):
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	case errors.Is(err, svc.ErrSvcEmployeeNotFound),
		errors.Is(err, svc.ErrSvcCompanyNotFound),
		errors.Is(err, svc.ErrSvcDepartmentNotFound):
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	case errors.Is(err, svc.ErrSvcConflict):
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "conflict"})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal error"})
	}
}
