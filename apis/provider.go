package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/masterraf21/trofonomie-backend/models"

	httpUtil "github.com/masterraf21/trofonomie-backend/utils/http"
)

type providerAPI struct {
	ProviderUsecase models.ProviderUsecase
}

// NewProviderAPI will create api for provider
func NewProviderAPI(r *mux.Router, puc models.ProviderUsecase) {
	providerAPI := &providerAPI{
		ProviderUsecase: puc,
	}

	r.HandleFunc("/provider", providerAPI.Create).Methods("POST")
	r.HandleFunc("/provider", providerAPI.GetAll).Methods("GET")
	r.HandleFunc("/provider/{id_provider}", providerAPI.GetByID).Methods("GET")
}

func (s *providerAPI) Create(w http.ResponseWriter, r *http.Request) {
	var body models.ProviderBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtil.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id, err := s.ProviderUsecase.CreateProvider(body)
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to create provider", http.StatusInternalServerError)
		return
	}

	var response struct {
		ID uint32 `json:"id_provider"`
	}
	response.ID = id

	httpUtil.HandleJSONResponse(w, r, response)
}

func (s *providerAPI) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := s.ProviderUsecase.GetAll()
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get provider data", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []models.Provider `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}

func (s *providerAPI) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	providerID, err := strconv.ParseInt(params["id_provider"], 10, 64)
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

	result, err := s.ProviderUsecase.GetByID(uint32(providerID))
	if err != nil {
		httpUtil.HandleError(w, r, err, "failed to get provider data by id", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data *models.Provider `json:"data"`
	}
	data.Data = result

	httpUtil.HandleJSONResponse(w, r, data)
}
