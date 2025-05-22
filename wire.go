// wire.go
//go:build wireinject
// +build wireinject

package main

import (
	"api_pegadaian/controllers"
	"api_pegadaian/infrastucture"
	"api_pegadaian/repository"
	"api_pegadaian/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	db, err := infrastucture.ConnectDB("postgres://postgres:password@localhost/db_pegadaian?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}

var RepositorySet = wire.NewSet(
	repository.NewBranchLabaRepository,
)

var UsecaseSet = wire.NewSet(
	usecase.NewBranchLabaUsecase,
)

var ControllerSet = wire.NewSet(
	controllers.NewBranchLabaController,
)

func InitializeApplication() (*fiber.App, error) {
	wire.Build(
		InitializeDB,
		RepositorySet,
		UsecaseSet,
		ControllerSet,
		NewApp,
	)
	return &fiber.App{}, nil
}

func NewApp(branchLabaRepository repository.BranchLabaRepository, branchLabaController *controllers.BranchLabaController) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	SetupRoutes(app, branchLabaRepository, branchLabaController)
	return app
}
