package cat

import (
	"catsocial/http/middleware"
	"catsocial/internal/domain/cat/service"

	"github.com/go-chi/chi"
)

type CatHandler struct {
	CatService    service.CatService
	JWTMiddleware *middleware.JWT
}

func ProvideCatHandler(catService service.CatService, jwt *middleware.JWT) CatHandler {
	return CatHandler{
		CatService:    catService,
		JWTMiddleware: jwt,
	}
}

func (h *CatHandler) Router(r chi.Router) {
	r.Route("/cat", func(r chi.Router) {
		r.Post("/add", h.AddCat)
	})
}
