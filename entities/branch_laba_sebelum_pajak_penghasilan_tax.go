package entities

import "time"

type BranchLabaSebelumPajakPenghasilanTax struct {
	ID                      int       `json:"id"`
	LabelRekonsiliasiFiskal string    `json:"label_rekonsiliasi_fiskal"`
	Periode                 time.Time `json:"periode"`
	Nilai                   float32   `json:"nilai" gorm:"type:numeric(20,2);not null"`
}
