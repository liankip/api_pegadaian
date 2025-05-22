package usecase

import (
	"api_pegadaian/entities"
	"api_pegadaian/repository"
	"api_pegadaian/utils"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"strconv"
	"time"
)

type BranchLabaUsecase struct {
	BranchLabaRepository repository.BranchLabaRepository
}

func NewBranchLabaUsecase(branchLabaRepository repository.BranchLabaRepository) *BranchLabaUsecase {
	return &BranchLabaUsecase{
		BranchLabaRepository: branchLabaRepository,
	}
}

func (branchLabaUsecase *BranchLabaUsecase) Collections(pagination utils.Pagination) (*utils.Pagination, error) {
	profile, err := branchLabaUsecase.BranchLabaRepository.CollectionsBranchLabaSebelumPajakPenghasilanTax(pagination)

	return profile, err
}

func (branchLabaUsecase *BranchLabaUsecase) Export() (*excelize.File, error) {
	var data, err = branchLabaUsecase.BranchLabaRepository.Export()
	if err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetSheetName("Sheet1", sheet)

	headers := []string{"ID", "Label Rekonsiliasi Fiskal", "Periode", "Nilai"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	for i, d := range data {
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), d.ID)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), d.LabelRekonsiliasiFiskal)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), d.Periode)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), float64(d.Nilai))
	}

	return f, nil
}

func (branchLabaUsecase *BranchLabaUsecase) Import(file io.Reader) error {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return err
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return err
	}

	var result []entities.BranchLabaSebelumPajakPenghasilanTax

	for i, row := range rows {
		if i == 0 {
			continue
		}

		if len(row) < 4 {
			continue
		}

		periode, err := time.Parse("2006-01-02", row[2])
		if err != nil {
			return fmt.Errorf("invalid periode at row %d: %w", i+1, err)
		}

		nilaiFloat, err := strconv.ParseFloat(row[3], 32)
		if err != nil {
			return fmt.Errorf("invalid nilai at row %d: %w", i+1, err)
		}

		data := entities.BranchLabaSebelumPajakPenghasilanTax{
			LabelRekonsiliasiFiskal: row[1],
			Periode:                 periode,
			Nilai:                   float32(nilaiFloat),
		}
		result = append(result, data)
	}

	return branchLabaUsecase.BranchLabaRepository.Import(result)
}
