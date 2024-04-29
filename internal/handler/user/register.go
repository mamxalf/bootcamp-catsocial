package user

import (
	"catsocial/internal/domain/user/request"
	"catsocial/shared/failure"
	"catsocial/shared/response"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

// Register sign up user.
// @Summary Register User
// @Description This endpoint for Register User.
// @Tags User
// @Accept  json
// @Produce json
// @Param request body request.RegisterRequest true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/user/register [post]
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var registerRequest request.RegisterRequest
	if err := decoder.Decode(&registerRequest); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err := registerRequest.Validate(); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	res, err := h.UserService.RegisterNewUser(r.Context(), registerRequest)
	if err != nil {
		log.Warn().Err(err).Msg("[Register Handler]")
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, res)
}
