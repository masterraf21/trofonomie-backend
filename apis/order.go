package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/masterraf21/trofonomie-backend/models"
	httpUtil "github.com/masterraf21/trofonomie-backend/utils/http"
)

type orderAPI struct {
	OrderUsecase models.OrderUsecase
}

// NewOrderAPI will create api for order
func NewOrderAPI(r *mux.Router, oru models.OrderUsecase) {
	orderAPI := &orderAPI{
		OrderUsecase: oru,
	}

	r.HandleFunc("/order", orderAPI.Create).Methods("POST")
	r.HandleFunc("/order", orderAPI.GetAll).Methods("GET")
	r.HandleFunc("/order/{id_order}", orderAPI.GetByID).Methods("GET")
	r.HandleFunc("/order/provider/{id_provider}", orderAPI.GetByProviderID).Methods("GET")
	r.HandleFunc("/order/customer/{id_customer}", orderAPI.GetByCustomerID).Methods("GET")
	r.HandleFunc("/order/provider/{id_provider}/accepted", orderAPI.GetAcceptedOrderByProviderID).Methods("GET")
	r.HandleFunc("/order/customer/{id_customer}/accepted", orderAPI.GetAcceptedOrderByCustomerID).Methods("GET")
	r.HandleFunc("/order/provider/{id_provider}/pending", orderAPI.GetPendingOrderByProviderID).Methods("GET")
	r.HandleFunc("/order/customer/{id_customer}/pending", orderAPI.GetPendingOrderByCustomerID).Methods("GET")
	r.HandleFunc("/order/{id_order}/accept", orderAPI.AcceptOrder).Methods("POST")
}

func (o *orderAPI) Create(w http.ResponseWriter, r *http.Request) {
	var body models.OrderBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtil.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	id, err := o.OrderUsecase.CreateOrder(body)
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to create order", http.StatusInternalServerError)
		return
	}

	var response struct {
		ID uint32 `json:"id_order"`
	}
	response.ID = id

	httpUtil.HandleJSONResponse(w, r, response)
}

func (o *orderAPI) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := o.OrderUsecase.GetAll()
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get order data", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Order `json:"data"`
	}
	data.Data = result
	httpUtil.HandleJSONResponse(w, r, data)
}

func (o *orderAPI) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	orderID, err := strconv.ParseInt(params["id_order"], 10, 64)
	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_order"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := o.OrderUsecase.GetByID(uint32(orderID))
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get order data by id", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data *models.Order `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}

func (o *orderAPI) GetByProviderID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ProviderID, err := strconv.ParseInt(params["id_provider"], 10, 64)
	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_provider"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := o.OrderUsecase.GetByProviderID(uint32(ProviderID))
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get order data by id Provider", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Order `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}

func (o *orderAPI) GetByCustomerID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	CustomerID, err := strconv.ParseInt(params["id_customer"], 10, 64)
	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_customer"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := o.OrderUsecase.GetByCustomerID(uint32(CustomerID))
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get order data by id Customer", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Order `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}

func (o *orderAPI) GetAcceptedOrderByProviderID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ProviderID, err := strconv.ParseInt(params["id_provider"], 10, 64)
	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_provider"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := o.OrderUsecase.GetByProviderIDAndStatus(uint32(ProviderID), "Accepted")
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get order data by id Provider", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Order `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}

func (o *orderAPI) GetPendingOrderByProviderID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ProviderID, err := strconv.ParseInt(params["id_provider"], 10, 64)
	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_provider"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := o.OrderUsecase.GetByProviderIDAndStatus(uint32(ProviderID), "Pending")
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get order data by id Provider", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Order `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}

func (o *orderAPI) GetAcceptedOrderByCustomerID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	CustomerID, err := strconv.ParseInt(params["id_customer"], 10, 64)
	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_customer"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := o.OrderUsecase.GetByCustomerIDAndStatus(uint32(CustomerID), "Accepted")
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get order data by id Customer", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Order `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}

func (o *orderAPI) GetPendingOrderByCustomerID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	CustomerID, err := strconv.ParseInt(params["id_customer"], 10, 64)
	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_customer"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := o.OrderUsecase.GetByCustomerIDAndStatus(uint32(CustomerID), "Pending")
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get order data by id Customer", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Order `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}

func (o *orderAPI) AcceptOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	orderID, err := strconv.ParseInt(params["id_order"], 10, 64)
	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_order"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	err = o.OrderUsecase.AcceptOrder(uint32(orderID))
	if err != nil {
		httpUtil.HandleError(w, r, err, "error accepting order", http.StatusInternalServerError)
		return
	}

	httpUtil.HandleNoJSONResponse(w)
}
