package calcservice

import (
	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
)

type CalculationsService interface {
	CreateCalculation(expression string) (Calculation, error)
	GetAllCalculations() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalculation(id, expression string) (Calculation, error)
	DeleteCalculation(id string) error
}

type calcService struct {
	repo CalculationsRepository
}

func NewCalculationService(r CalculationsRepository) CalculationsService {
	return &calcService{repo: r}
}

func (s *calcService) calculateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result), nil
}

// CreateCalculation implements CaltulationsService.
func (s *calcService) CreateCalculation(expression string) (Calculation, error) {
	result, err := s.calculateExpression(expression)

	if err != nil {
		return Calculation{}, err
	}

	calc := Calculation{
		ID:         uuid.NewString(),
		Expression: expression,
		Result:     result,
	}

	if err := s.repo.CreateCalculation(calc); err != nil {
		return Calculation{}, err
	}
	return calc, nil
}

// DeleteCalculation implements CaltulationsService.
func (s *calcService) DeleteCalculation(id string) error {
	return s.repo.DeleteCalculation(id)
}

// GetAllCalculations implements CaltulationsService.
func (s *calcService) GetAllCalculations() ([]Calculation, error) {
	return s.repo.GetAllCalculation()
}

// GetCalculationByID implements CaltulationsService.
func (s *calcService) GetCalculationByID(id string) (Calculation, error) {
	return s.repo.GetCalculationByID(id)
}

// UpdateCalculation implements CaltulationsService.
func (s *calcService) UpdateCalculation(id string, expression string) (Calculation, error) {
	calc, err := s.repo.GetCalculationByID(id)
	if err != nil {
		return Calculation{}, err
	}
	result, err := s.calculateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}
	calc.Expression = expression
	calc.Result = result

	if err := s.repo.UpdateCalculation(calc); err != nil {
		return Calculation{}, err
	}
	return calc, nil
}
