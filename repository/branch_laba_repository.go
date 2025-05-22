package repository

import (
	"api_pegadaian/entities"
	"api_pegadaian/utils"

	"gorm.io/gorm"
)

type BranchLabaRepository interface {
	CollectionsBranchLabaSebelumPajakPenghasilanTax(pagination utils.Pagination) (*utils.Pagination, error)

	Export() ([]entities.BranchLabaSebelumPajakPenghasilanTax, error)

	Import([]entities.BranchLabaSebelumPajakPenghasilanTax) error
}

type BranchLabaRepositoryImpl struct {
	db *gorm.DB
}

func NewBranchLabaRepository(db *gorm.DB) BranchLabaRepository {
	return &BranchLabaRepositoryImpl{db: db}
}

func (p *BranchLabaRepositoryImpl) CollectionsBranchLabaSebelumPajakPenghasilanTax(pagination utils.Pagination) (*utils.Pagination, error) {
	var result []*entities.BranchLabaSebelumPajakPenghasilanTax

	tx := p.db.Model(&entities.BranchLabaSebelumPajakPenghasilanTax{}).
		Select("label_rekonsiliasi_fiskal, periode, SUM(nilai) as nilai").
		Group("label_rekonsiliasi_fiskal, periode")

	tx.Scopes(utils.Paginate(result, &pagination, p.db)).Find(&result)
	pagination.Rows = result

	return &pagination, nil
}

func (p *BranchLabaRepositoryImpl) Export() ([]entities.BranchLabaSebelumPajakPenghasilanTax, error) {
	var data []entities.BranchLabaSebelumPajakPenghasilanTax

	if err := p.db.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *BranchLabaRepositoryImpl) Import(data []entities.BranchLabaSebelumPajakPenghasilanTax) error {
	return r.db.Create(&data).Error
}
