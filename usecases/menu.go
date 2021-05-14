package usecases

import "github.com/masterraf21/trofonomie-backend/models"

type menuUsecase struct {
	Repo       models.MenuRepository
	SellerRepo models.ProviderRepository
}

// NewMenuUsecase will initiate usecase
func NewMenuUsecase(prr models.MenuRepository, ssr models.ProviderRepository) models.MenuUsecase {
	return &menuUsecase{Repo: prr, SellerRepo: ssr}
}

func (u *menuUsecase) CreateMenu(body models.MenuBody) (id uint32, err error) {
	providerPtr, err := u.SellerRepo.GetByID(body.ProviderID)
	if err != nil {
		return
	}

	product := models.Menu{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		ProviderID:  body.ProviderID,
		Provider:    providerPtr,
	}

	oid, err := u.Repo.Store(&product)
	if err != nil {
		return
	}
	result, err := u.Repo.GetByOID(oid)
	if err != nil {
		return
	}

	id = result.MenuID

	return
}

func (u *menuUsecase) GetAll() (res []models.Menu, err error) {
	res, err = u.Repo.GetAll()
	if len(res) == 0 {
		res = make([]models.Menu, 0)
	}
	return
}

func (u *menuUsecase) GetByProviderID(sellerID uint32) (res []models.Menu, err error) {
	res, err = u.Repo.GetByProviderID(sellerID)
	if len(res) == 0 {
		res = make([]models.Menu, 0)
	}
	return
}

func (u *menuUsecase) GetByID(id uint32) (res *models.Menu, err error) {
	res, err = u.Repo.GetByID(id)
	return
}
