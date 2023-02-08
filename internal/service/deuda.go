package service

import (
	"github.com/PGNahuel/FinHelp/internal/domain"
	"github.com/PGNahuel/FinHelp/internal/repository"
)

type Deuda interface {
	AgregarDeuda(d domain.Deduda) (id int, err error)
	ObtenerDeuda(deudaId int) (deuda domain.Deduda, err error)
	ObtenerDeudas() (deudas []domain.Deduda, err error)
}

type deudaService struct {
	Deuda
	repo repository.Deuda
}

func NewDeudaService(r repository.Deuda) *deudaService {
	return &deudaService{
		repo: r,
	}
}
