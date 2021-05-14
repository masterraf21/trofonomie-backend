package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/masterraf21/trofonomie-backend/models"

	httpUtil "github.com/masterraf21/trofonomie-backend/utils/http"
)

type customerAPI struct {
	customerUsecase models.CustomerUsecase
}

// NewCustomerAPI will create api for customer
func NewCustomerAPI(r *mux.Router, cus models.CustomerUsecase) {
	customerAPI := &customerAPI{
		customerUsecase: cus,
	}

	r.HandleFunc("/customer", customerAPI.Create).Methods("POST")
	r.HandleFunc("/customer", customerAPI.GetAll).Methods("GET")
	r.HandleFunc("/customer/{id_customer}", customerAPI.GetByID).Methods("GET")
}

func (b *customerAPI) Create(w http.ResponseWriter, r *http.Request) {
	var body models.CustomerBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtil.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	id, err := b.customerUsecase.CreateCustomer(body)
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to creata customer", http.StatusInternalServerError)
		return
	}

	var response struct {
		ID uint32 `json:"id_customer"`
	}
	response.ID = id

	httpUtil.HandleJSONResponse(w, r, response)
}

func (b *customerAPI) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := b.customerUsecase.GetAll()
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get buyer data", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Customer `json:"data"`
	}
	data.Data = result
	httpUtil.HandleJSONResponse(w, r, data)
}

func (b *customerAPI) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	customerID, err := strconv.ParseInt(params["id_customer"], 10, 64)
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

	result, err := b.customerUsecase.GetByID(uint32(customerID))
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get customer data by id", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data *models.Customer `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}
