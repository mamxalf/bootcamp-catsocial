package health

import (
	"catsocial/shared/response"
	"net/http"

	"github.com/rs/zerolog/log"
)

// Ping get health message.
// @Summary Ping.
// @Description This endpoint for get health message.
// @Tags Health
// @Accept  json
// @Produce json
// @Success 200 {object} response.Base
// @Failure 503 {object} response.Base
// @Router /v1/health/ping [get]
func (h *HealthHandler) Ping(w http.ResponseWriter, _ *http.Request) {
	message := h.HealthService.Ping()
	response.WithMessage(w, http.StatusOK, message)
}

// PingDB get health message.
// @Summary PingDB.
// @Description This endpoint for get health message from postgres ping.
// @Tags Health
// @Accept  json
// @Produce json
// @Success 200 {object} response.Base
// @Failure 503 {object} response.Base
// @Router /v1/health/ping-db [get]
func (h *HealthHandler) PingDB(w http.ResponseWriter, _ *http.Request) {
	message, err := h.HealthService.PingDB()
	if err != nil {
		log.Warn().Err(err).Msg("[Handler Ping DB] - error pinging db")
		response.WithError(w, err)
		return
	}
	response.WithMessage(w, http.StatusOK, message)
}
