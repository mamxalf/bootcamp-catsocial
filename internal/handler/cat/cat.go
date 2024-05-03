package cat

import (
	// "catsocial/internal/domain/cat/request"
	// "catsocial/shared/failure"
	"catsocial/shared/response"
	// "encoding/json"
	"net/http"
	// "github.com/rs/zerolog/log"
)

// AddCat adds a new cat.
// @Summary Add a new cat
// @Description This endpoint is used to add a new cat.
// @Tags Cat
// @Accept json
// @Produce json
// @Param request body request.CatRequest true "Request Body"
// @Success 201 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/cat/add [post]
// func (h *CatHandler) AddCat(w http.ResponseWriter, r *http.Request) {
// 	decoder := json.NewDecoder(r.Body)
// 	var catRequest request.CatRequest
// 	if err := decoder.Decode(&catRequest); err != nil {
// 		response.WithError(w, failure.BadRequest(err))
// 		return
// 	}

// 	if err := catRequest.Validate(); err != nil {
// 		response.WithError(w, failure.BadRequest(err))
// 		return
// 	}
// 	res, err := h.CatService.AddCat(r.Context(), catRequest)
// 	if err != nil {
// 		log.Warn().Err(err).Msg("[Add New Cat]")
// 		response.WithError(w, err)
// 		return
// 	}

//		response.WithJSON(w, http.StatusOK, res)
//	}

func (h *CatHandler) AddCat(w http.ResponseWriter, r *http.Request) {
	successMessage := "Success"
	response.WithJSON(w, http.StatusOK, &response.Base{Message: &successMessage})
}
