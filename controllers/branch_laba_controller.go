package controllers

import (
	"api_pegadaian/entities"
	"api_pegadaian/usecase"
	"api_pegadaian/utils"
	"bytes"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type BranchLabaController struct {
	BranchLabaUsecase *usecase.BranchLabaUsecase
}

func NewBranchLabaController(branchLabaUsecase *usecase.BranchLabaUsecase) *BranchLabaController {
	return &BranchLabaController{BranchLabaUsecase: branchLabaUsecase}
}

func (v *BranchLabaController) Collections(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	result, _ := v.BranchLabaUsecase.Collections(utils.Pagination{
		Page:  page,
		Limit: limit,
		Sort:  c.Query("sort"),
	})

	return c.JSON(entities.Response{
		Message: "Collection branch laba sebelum pajak penghasilan tax successful",
		Data:    result,
	})
}

func (v *BranchLabaController) Export(c *fiber.Ctx) error {
	f, err := v.BranchLabaUsecase.Export()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ResponseError{
			Message: "Failed to export",
			Error:   err.Error(),
		})
	}

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ResponseError{
			Message: "Failed to write Excel buffer",
			Error:   err.Error(),
		})
	}

	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set("Content-Disposition", "attachment; filename=branch_laba.xlsx")
	return c.SendStream(&buf)
}

func (v *BranchLabaController) Import(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ResponseError{Message: "File required"})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ResponseError{Message: "Failed to open file"})
	}
	defer file.Close()

	if err := v.BranchLabaUsecase.Import(file); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ResponseError{Message: "Import failed", Error: err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Import successful"})
}
