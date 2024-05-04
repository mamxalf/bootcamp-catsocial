package cat

import (
	"catsocial/internal/domain/cat/request"
	"catsocial/shared/failure"
	"catsocial/shared/response"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
)

// Insert a new match.
// @Summary Insert a new match
// @Description This endpoint is used to insert a new match.
// @Tags Match
// @Accept json
// @Produce json
// @Param request body request.MatchRequest true "Request Body"
// @Success 201 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/cat/match [post]
func (h *CatHandler) InsertNewMatch(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var matchRequest request.MatchRequest
	if err := decoder.Decode(&matchRequest); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err := matchRequest.Validate(); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	res, err := h.CatService.InsertNewMatch(r.Context(), matchRequest)
	if err != nil {
		log.Warn().Err(err).Msg("[Add New Cat]")
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, res)
}

// Get All Match Data.
// @Summary GetAllMatch.
// @Description This endpoint for get all Match data.
// @Tags Match
// @Accept  json
// @Produce json
// @Success 200 {object} response.Base
// @Failure 503 {object} response.Base
// @Router /v1/cat/match [get]
func (h *CatHandler) FindAllMatchData(w http.ResponseWriter, r *http.Request) {
	res, err := h.CatService.GetAllMatchesData(r.Context())
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}

// Approve Cat Match.
// @Summary Approve Cat Match
// @Description This endpoint is used to Approve Cat Match.
// @Tags Match
// @Accept json
// @Produce json
// @Param request body request.MatchApproval true "Request Body"
// @Success 201 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/cat/match/approve [post]
func (h *CatHandler) ApproveCatMatch(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	res, err := h.CatService.ApproveCatMatch(r.Context(), idStr)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}

// Reject Cat Match.
// @Summary Reject Cat Match
// @Description This endpoint is used to Reject Cat Match.
// @Tags Match
// @Accept json
// @Produce json
// @Param request body request.MatchRequest true "Request Body"
// @Success 201 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/cat/match/reject [post]
func (h *CatHandler) RejectCatMatch(w http.ResponseWriter, r *http.Request) {
	res := "success"
	response.WithJSON(w, http.StatusOK, res)
}

// Delete match data.
// @Summary Delete match data
// @Description This endpoint is used to delete match.
// @Tags Match
// @Accept json
// @Produce json
// @Param id path string true "ID by match"
// @Success 200 {object} response.Base
// @Failure 401 {object} response.Base
// @Failure 404 {object} response.Base
// @Router /v1/cat/match/{id} [delete]
func (h *CatHandler) DeleteCatMatch(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	res, err := h.CatService.DeleteCatMatch(r.Context(), idStr)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}
