package usecases

import (
	"github.com/masterraf21/trofonomie-backend/models"
)

type providerUsecase struct {
	Repo models.ProviderRepository
}

// NewProviderUsecase will initiate usecase
func NewProviderUsecase(srr models.ProviderRepository) models.ProviderUsecase {
	return &providerUsecase{Repo: srr}
}

func (u *providerUsecase) CreateProvider(body models.ProviderBody) (id uint32, err error) {
	provider := models.Provider{
		Name:        body.Name,
		Address:     body.Address,
		PhoneNumber: body.PhoneNumber,
		OpenHour:    body.OpenHour,
		ClosedHour:  body.OpenHour,
	}

	oid, err := u.Repo.Store(&provider)
	if err != nil {
		return
	}
	result, err := u.Repo.GetByOID(oid)
	if err != nil {
		return
	}

	id = result.ProviderID

	return
}

func (u *providerUsecase) GetAll() (res []models.Provider, err error) {
	res, err = u.Repo.GetAll()
	if len(res) == 0 {
		res = make([]models.Provider, 0)
	}
	return
}

func (u *providerUsecase) GetByID(id uint32) (res *models.Provider, err error) {
	res, err = u.Repo.GetByID(id)
	return
}
