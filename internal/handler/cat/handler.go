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
		r.Post("/add", h.InsertNewCat)
		r.Get("/{id}", h.Find)
		r.Get("/", h.FindAllCatData)
		r.Put("/{id}", h.UpdateCatData)
		r.Delete("/{id}", h.DeleteCatData)
		r.Post("/match", h.InsertNewMatch)
		r.Get("/match", h.FindAllMatchData)
		r.Post("/match/approve", h.ApproveCatMatch)
		r.Post("/match/reject", h.RejectCatMatch)
		r.Delete("/match/{id}", h.DeleteCatMatch)
	})
}
