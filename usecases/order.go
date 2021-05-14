package usecases

import (
	"github.com/masterraf21/trofonomie-backend/models"
)

type orderUsecase struct {
	Repo         models.OrderRepository
	CustomerRepo models.CustomerRepository
	ProviderRepo models.ProviderRepository
	MenuRepo     models.MenuRepository
}

// NewOrderUsecase will nititate usecase for order
func NewOrderUsecase(
	orr models.OrderRepository,
	brr models.CustomerRepository,
	slr models.ProviderRepository,
	prr models.MenuRepository,
) models.OrderUsecase {
	return &orderUsecase{
		Repo:         orr,
		CustomerRepo: brr,
		ProviderRepo: slr,
		MenuRepo:     prr,
	}
}

func (u *orderUsecase) CreateOrder(body models.OrderBody) (id uint32, err error) {
	ProviderPtr, err := u.ProviderRepo.GetByID(body.ProviderID)
	if err != nil {
		return
	}

	CustomerPtr, err := u.CustomerRepo.GetByID(body.CustomerID)
	if err != nil {
		return
	}

	orderDetails := make([]models.OrderDetail, 0)

	var menuPtr *models.Menu
	for _, menu := range body.Menus {
		menuPtr, err = u.MenuRepo.GetByID(menu.MenuID)
		if err != nil {
			return
		}

		var orderDetail models.OrderDetail

		if menuPtr != nil {
			orderDetail = models.OrderDetail{
				MenuID:     menu.MenuID,
				Quantity:   menu.Quantity,
				Menu:       menuPtr,
				TotalPrice: float32(menu.Quantity) * menuPtr.Price,
			}
		} else {
			orderDetail = models.OrderDetail{
				MenuID:   menu.MenuID,
				Quantity: menu.Quantity,
			}
		}

		orderDetails = append(orderDetails, orderDetail)
	}

	totalPrice := float32(0)
	for _, orderDetail := range orderDetails {
		totalPrice += orderDetail.TotalPrice
	}

	order := models.Order{
		CustomerID:   body.CustomerID,
		Customer:     CustomerPtr,
		ProviderID:   body.ProviderID,
		Provider:     ProviderPtr,
		OrderDetails: orderDetails,
		TotalPrice:   totalPrice,
		Status:       "Pending",
	}
	if ProviderPtr != nil {
		order.SourceAddress = ProviderPtr.Address
	}
	if CustomerPtr != nil {
		order.DeliveryAddress = CustomerPtr.Address
	}

	oid, err := u.Repo.Store(&order)
	if err != nil {
		return
	}

	result, err := u.Repo.GetByOID(oid)
	if err != nil {
		return
	}
	id = result.ID

	return
}

func (u *orderUsecase) AcceptOrder(id uint32) (err error) {
	err = u.Repo.UpdateArbitrary(id, "status", "Accepted")
	return
}

func (u *orderUsecase) GetAll() (res []models.Order, err error) {
	res, err = u.Repo.GetAll()
	return
}

func (u *orderUsecase) GetByID(id uint32) (res *models.Order, err error) {
	res, err = u.Repo.GetByID(id)
	return
}

func (u *orderUsecase) GetByProviderID(ProviderID uint32) (res []models.Order, err error) {
	res, err = u.Repo.GetByProviderID(ProviderID)
	return
}

func (u *orderUsecase) GetByCustomerID(CustomerID uint32) (res []models.Order, err error) {
	res, err = u.Repo.GetByCustomerID(CustomerID)
	return
}

func (u *orderUsecase) GetByCustomerIDAndStatus(CustomerID uint32, status string) (res []models.Order, err error) {
	res, err = u.Repo.GetByCustomerIDAndStatus(CustomerID, status)
	return
}

func (u *orderUsecase) GetByProviderIDAndStatus(ProviderID uint32, status string) (res []models.Order, err error) {
	res, err = u.Repo.GetByCustomerIDAndStatus(ProviderID, status)
	return
}

func (u *orderUsecase) GetByStatus(status string) (res []models.Order, err error) {
	res, err = u.Repo.GetByStatus(status)
	return
}
