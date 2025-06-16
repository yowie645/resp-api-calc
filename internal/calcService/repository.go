package calcservice

import (
	"gorm.io/gorm"
)

type CalculationsRepository interface {
	CreateCalculation(calc Calculation) error
	GetAllCalculation() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalculation(calc Calculation) error
	DeleteCalculation(id string) error
}

type calcRepository struct {
	db *gorm.DB
}

func NewCalculationsRepository(db *gorm.DB) CalculationsRepository {
	return &calcRepository{
		db: db,
	}
}

func (r *calcRepository) CreateCalculation(calc Calculation) error {
	return r.db.Create(&calc).Error
}

func (r *calcRepository) GetAllCalculation() ([]Calculation, error) {
	var calculations []Calculation

	err := r.db.Find(&calculations).Error
	return calculations, err
}

func (r *calcRepository) GetCalculationByID(id string) (Calculation, error) {
	var calc Calculation

	err := r.db.First(&calc, "id = ?", id).Error
	return calc, err
}

func (r *calcRepository) UpdateCalculation(calc Calculation) error {
	return r.db.Save(&calc).Error
}

func (r *calcRepository) DeleteCalculation(id string) error {
	return r.db.Delete(&Calculation{ID: id}).Error
}
