package router

import (
	"pfms/api/saving"
	"pfms/middleware"

	"github.com/gofiber/fiber/v2"
)

func setRouteSaving(v1 fiber.Router) {
	savingV1 := v1.Group("/saving")

	savingV1.Post("/add-saving", middleware.AuthJwt(), saving.AddSaving)
	savingV1.Delete("/delete-saving", middleware.AuthJwt(), saving.DeleteSavingFromName)

}
