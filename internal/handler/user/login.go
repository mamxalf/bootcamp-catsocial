package user

import (
	"catsocial/internal/domain/user/request"
	"catsocial/shared/failure"
	"catsocial/shared/response"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

// Login sign in user.
// @Summary Login User
// @Description This endpoint for Login User.
// @Tags User
// @Accept  json
// @Produce json
// @Param request body request.LoginRequest true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/user/login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var loginRequest request.LoginRequest
	if err := decoder.Decode(&loginRequest); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err := loginRequest.Validate(); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	res, err := h.UserService.LoginUser(r.Context(), loginRequest)
	if err != nil {
		log.Warn().Err(err).Msg("[Login - User Handler]")
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, res)
}
