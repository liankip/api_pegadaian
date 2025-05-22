package main

import (
	"api_pegadaian/controllers"
	"api_pegadaian/repository"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, branchLabaRepository repository.BranchLabaRepository, branchLabaController *controllers.BranchLabaController) {

	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.Get("/", func(ctx *fiber.Ctx) error {
				return ctx.JSON(fiber.Map{
					"message": "api/v1",
				})
			})

			branchLaba := v1.Group("/branch-laba")
			{
				branchLaba.Get("/", branchLabaController.Collections)
				branchLaba.Get("/export", branchLabaController.Export)
				branchLaba.Get("/import", branchLabaController.Import)
			}
		}
	}
}
