package repository

import "github.com/PGNahuel/FinHelp/internal/domain"

type Deuda interface {
	AgregarDeuda(d domain.Deduda) (id int, err error)
	ObtenerDeuda(deudaId int) (deuda domain.Deduda, err error)
	ObtenerDeudas() (deudas []domain.Deduda, err error)
}
