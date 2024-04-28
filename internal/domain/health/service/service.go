package service

import (
	"catsocial/configs"
	"catsocial/internal/domain/health/repository"

	"github.com/rs/zerolog/log"
)

type HealthService interface {
	Ping() (message string)
	PingDB() (message string, err error)
}

type HealthServiceImpl struct {
	HealthRepository repository.HealthRepository
	Config           *configs.Config
}

// ProvideHealthServiceImpl is the provider for this service.
func ProvideHealthServiceImpl(
	healthRepository repository.HealthRepository,
	config *configs.Config,
) *HealthServiceImpl {
	s := new(HealthServiceImpl)
	s.HealthRepository = healthRepository
	s.Config = config

	return s
}

func (h *HealthServiceImpl) Ping() (message string) {
	return h.HealthRepository.Ping()
}

func (h *HealthServiceImpl) PingDB() (message string, err error) {
	message, err = h.HealthRepository.PingDB()
	if err != nil {
		log.Err(err).Msg("[PingDB] - error pinging database")
		return
	}
	return
}
