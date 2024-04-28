package repository

import (
	"catsocial/infras"
	"catsocial/shared/logger"
)

type HealthRepository interface {
	Ping() (message string)
	PingDB() (message string, err error)
}

type HealthRepositoryInfra struct {
	DB *infras.PostgresConn
}

func ProvideHealthRepositoryInfra(db *infras.PostgresConn) *HealthRepositoryInfra {
	r := new(HealthRepositoryInfra)
	r.DB = db
	return r
}

func (h *HealthRepositoryInfra) Ping() (message string) {
	return "PONG!!!"
}

func (h *HealthRepositoryInfra) PingDB() (message string, err error) {
	if err = h.DB.PG.Ping(); err != nil {
		logger.ErrorWithStack(err)
		return
	}
	return "Postgres is healthy", nil
}
