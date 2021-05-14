package usecases

import (
	"github.com/masterraf21/trofonomie-backend/models"
)

type customerUsecase struct {
	Repo models.CustomerRepository
}

// NewCustomerUsecase will initiate usecase
func NewCustomerUsecase(br models.CustomerRepository) models.CustomerUsecase {
	return &customerUsecase{Repo: br}
}

func (u *customerUsecase) CreateCustomer(body models.CustomerBody) (id uint32, err error) {
	customer := models.Customer{
		Name:    body.Name,
		Address: body.Address,
	}

	oid, err := u.Repo.Store(&customer)
	if err != nil {
		return
	}
	result, err := u.Repo.GetByOID(oid)
	if err != nil {
		return
	}

	id = result.CustomerID

	return
}

func (u *customerUsecase) GetAll() (res []models.Customer, err error) {
	res, err = u.Repo.GetAll()
	return
}

func (u *customerUsecase) GetByID(id uint32) (res *models.Customer, err error) {
	res, err = u.Repo.GetByID(id)
	return
}
