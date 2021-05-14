package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/masterraf21/trofonomie-backend/models"
	httpUtil "github.com/masterraf21/trofonomie-backend/utils/http"
)

type menuAPI struct {
	MenuUsecase models.MenuUsecase
}

// NewMenuAPI will create api for Menu
func NewMenuAPI(r *mux.Router, mus models.MenuUsecase) {
	menuAPI := &menuAPI{
		MenuUsecase: mus,
	}

	r.HandleFunc("/menu", menuAPI.Create).Methods("POST")
	r.HandleFunc("/menu", menuAPI.GetAll).Methods("GET")
	r.HandleFunc("/menu/{id_menu}", menuAPI.GetByID).Methods("GET")
	r.HandleFunc("/menu/provider/{id_provider}", menuAPI.GetBySellerID).Methods("GET")
}

func (p *menuAPI) Create(w http.ResponseWriter, r *http.Request) {
	var body models.MenuBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtil.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	id, err := p.MenuUsecase.CreateMenu(body)
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to creata Menu", http.StatusInternalServerError)
		return
	}

	var response struct {
		ID uint32 `json:"id_menu"`
	}
	response.ID = id

	httpUtil.HandleJSONResponse(w, r, response)
}

func (p *menuAPI) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := p.MenuUsecase.GetAll()
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get Menu data", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Menu `json:"data"`
	}
	data.Data = result
	httpUtil.HandleJSONResponse(w, r, data)
}

func (p *menuAPI) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	MenuID, err := strconv.ParseInt(params["id_menu"], 10, 64)
	if err != nil {
		httpUtil.HandleError(
			w,
			r,
			err,
			params["id_menu"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := p.MenuUsecase.GetByID(uint32(MenuID))
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get Menu data by id", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data *models.Menu `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}

func (p *menuAPI) GetBySellerID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	sellerID, err := strconv.ParseInt(params["id_provider"], 10, 64)
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

	result, err := p.MenuUsecase.GetByProviderID(uint32(sellerID))
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get Menu data by seller id", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Menu `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}
