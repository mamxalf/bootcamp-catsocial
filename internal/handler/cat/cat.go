package cat

import (
	"catsocial/internal/domain/cat/request"
	"catsocial/shared/failure"
	"catsocial/shared/response"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

// AddCat adds a new cat.
// @Summary Add a new cat
// @Description This endpoint is used to add a new cat.
// @Tags Cat
// @Accept json
// @Produce json
// @Param request body request.InsertCatRequest true "Request Body"
// @Success 201 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/cat/add [post]
func (h *CatHandler) InsertNewCat(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var catRequest request.InsertCatRequest
	if err := decoder.Decode(&catRequest); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err := catRequest.Validate(); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	res, err := h.CatService.InsertNewCat(r.Context(), catRequest)
	if err != nil {
		log.Warn().Err(err).Msg("[Add New Cat]")
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, res)
}

// Get cat by id.
// @Summary GetCat.
// @Description This endpoint for get cat data by ID.
// @Tags Cat
// @Accept  json
// @Produce json
// @Param id path string true "ID by cat"
// @Success 200 {object} response.Base
// @Failure 503 {object} response.Base
// @Router /v1/cat/:id [get]
func (h *CatHandler) Find(w http.ResponseWriter, r *http.Request) {
	res := "success"
	response.WithJSON(w, http.StatusOK, res)
}

// Get All Cat Data.
// @Summary GetAllCat.
// @Description This endpoint for get all cat data.
// @Tags Cat
// @Accept  json
// @Produce json
// @Success 200 {object} response.Base
// @Failure 503 {object} response.Base
// @Router /v1/cat/ [get]
func (h *CatHandler) FindAllCatData(w http.ResponseWriter, r *http.Request) {
	res := "success"
	response.WithJSON(w, http.StatusOK, res)
}

// Update cat data.
// @Summary Update cat data
// @Description This endpoint is used to update cat.
// @Tags Cat
// @Accept json
// @Produce json
// @Param id path string true "ID by cat"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 401 {object} response.Base
// @Failure 404 {object} response.Base
// @Router /v1/cat/:id [put]
func (h *CatHandler) UpdateCatData(w http.ResponseWriter, r *http.Request) {
	res := "success"
	response.WithJSON(w, http.StatusOK, res)
}

// Delete cat data.
// @Summary Delete cat data
// @Description This endpoint is used to delete cat.
// @Tags Cat
// @Accept json
// @Produce json
// @Param id path string true "ID by cat"
// @Success 200 {object} response.Base
// @Failure 401 {object} response.Base
// @Failure 404 {object} response.Base
// @Router /v1/cat/:id [delete]
func (h *CatHandler) DeleteCatData(w http.ResponseWriter, r *http.Request) {
	res := "success"
	response.WithJSON(w, http.StatusOK, res)
}
