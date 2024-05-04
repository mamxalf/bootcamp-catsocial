package cat

import (
	"catsocial/http/middleware"
	"catsocial/internal/domain/cat/request"
	"catsocial/shared/failure"
	"catsocial/shared/response"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

// InsertNewCat adds a new cat.
// @Summary Add a new cat
// @Description This endpoint is used to add a new cat.
// @Tags Cat
// @Accept json
// @Produce json
// @Security BearerToken
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
	claimUser, ok := middleware.GetClaimsUser(r.Context()).(jwt.MapClaims)
	if !ok {
		log.Warn().Msg("invalid claim jwt")
		err := failure.Unauthorized("invalid claim jwt")
		response.WithError(w, err)
		return
	}
	userID, err := uuid.Parse(claimUser["ownerID"].(string))
	if err != nil {
		log.Warn().Msg(err.Error())
		err = failure.Unauthorized("invalid format user_id")
		response.WithError(w, err)
		return
	}
	catRequest.UserID = userID
	res, err := h.CatService.InsertNewCat(r.Context(), catRequest)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusCreated, res)
}

// Find cat by id.
// @Summary GetCat.
// @Description This endpoint for get cat data by ID.
// @Tags Cat
// @Accept  json
// @Produce json
// @Security BearerToken
// @Param id path string true "ID by cat"
// @Success 200 {object} response.Base
// @Failure 503 {object} response.Base
// @Router /v1/cat/{id} [get]
func (h *CatHandler) Find(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	claimUser, ok := middleware.GetClaimsUser(r.Context()).(jwt.MapClaims)
	if !ok {
		log.Warn().Msg("invalid claim jwt")
		err := failure.Unauthorized("invalid claim jwt")
		response.WithError(w, err)
		return
	}
	userID, err := uuid.Parse(claimUser["ownerID"].(string))
	if err != nil {
		log.Warn().Msg(err.Error())
		err = failure.Unauthorized("invalid format user_id")
		response.WithError(w, err)
		return
	}

	res, err := h.CatService.GetCatData(r.Context(), userID, idStr)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}

// FindAllCatData All Cat Data.
// @Summary GetAllCat
// @Description This endpoint retrieves all cat data based on optional filtering parameters.
// @Tags Cat
// @Accept  json
// @Produce  json
// @Security BearerToken
// @Param id query string false "Limit the output based on the cat’s ID"
// @Param limit query int false "Limit the number of items in the response (default 5)" default(5)
// @Param offset query int false "Offset the start point of data return (default 0)" default(0)
// @Param race query string false "Filter based on race" Enums(Persian, Maine Coon, Siamese, Ragdoll, Bengal, Sphynx, British Shorthair, Abyssinian, Scottish Fold, Birman)
// @Param sex query string false "Filter based on sex" Enums(male, female)
// @Param hasMatched query bool false "Filter based on match status" Enums(true, false)
// @Param ageInMonth query string false "Filter based on age in months with operators (e.g., '=>4', '=<4', '=4')"
// @Param owned query bool false "Filter if the user owns the cat or not" Enums(true, false)
// @Param search query string false "Display information containing the search term in the name of the cat"
// @Success 200 {object} response.Base
// @Failure 503 {object} response.Base
// @Router /v1/cat/ [get]
func (h *CatHandler) FindAllCatData(w http.ResponseWriter, r *http.Request) {
	var request request.CatQueryParams
	query := r.URL.Query()
	// Directly retrieving the values
	request.Search = query.Get("search")
	request.ID = query.Get("id")
	request.Race = query.Get("race")
	request.Sex = query.Get("sex")
	request.AgeInMonth = query.Get("ageInMonth")

	// For numerical and boolean values, we need to safely parse them
	// because they require conversion from string
	if limit, err := strconv.Atoi(query.Get("limit")); err == nil {
		request.Limit = limit
	} else {
		request.Limit = 5 // default value if not specified or error in conversion
	}

	if offset, err := strconv.Atoi(query.Get("offset")); err == nil {
		request.Offset = offset
	} else {
		request.Offset = 0 // default value if not specified or error in conversion
	}

	if hasMatched, err := strconv.ParseBool(query.Get("hasMatched")); err == nil {
		request.HasMatched = hasMatched
	} // defaults to false if not specified or error

	if owned, err := strconv.ParseBool(query.Get("owned")); err == nil {
		request.Owned = owned
	}

	//if err != nil {
	//	err := failure.BadRequestFromString("request doesn’t pass validation")
	//	response.WithError(w, err)
	//	return
	//}

	claimUser, ok := middleware.GetClaimsUser(r.Context()).(jwt.MapClaims)
	if !ok {
		log.Warn().Msg("invalid claim jwt")
		err := failure.Unauthorized("invalid claim jwt")
		response.WithError(w, err)
		return
	}
	userID, err := uuid.Parse(claimUser["ownerID"].(string))
	if err != nil {
		log.Warn().Msg(err.Error())
		err = failure.Unauthorized("invalid format user_id")
		response.WithError(w, err)
		return
	}

	res, err := h.CatService.GetAllCatData(r.Context(), userID, request)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}

// UpdateCatData cat data.
// @Summary Update cat data
// @Description This endpoint is used to update cat.
// @Tags Cat
// @Accept json
// @Produce json
// @Security BearerToken
// @Param id path string true "ID by cat"
// @Param request body request.UpdateCatRequest true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 401 {object} response.Base
// @Failure 404 {object} response.Base
// @Router /v1/cat/{id} [put]
func (h *CatHandler) UpdateCatData(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	catID, err := uuid.Parse(idStr)
	if err != nil {
		log.Warn().Msg(err.Error())
		err = failure.Unauthorized("invalid format cat id")
		response.WithError(w, err)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var catRequest request.UpdateCatRequest
	if err := decoder.Decode(&catRequest); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err := catRequest.Validate(); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	claimUser, ok := middleware.GetClaimsUser(r.Context()).(jwt.MapClaims)
	if !ok {
		log.Warn().Msg("invalid claim jwt")
		err := failure.Unauthorized("invalid claim jwt")
		response.WithError(w, err)
		return
	}
	_, err = uuid.Parse(claimUser["ownerID"].(string))
	if err != nil {
		log.Warn().Msg(err.Error())
		err = failure.Unauthorized("invalid format user_id")
		response.WithError(w, err)
		return
	}
	res, err := h.CatService.UpdateCatData(r.Context(), catID, catRequest)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}

// DeleteCatData cat data.
// @Summary Delete cat data
// @Description This endpoint is used to delete cat.
// @Tags Cat
// @Accept json
// @Produce json
// @Security BearerToken
// @Param id path string true "ID by cat"
// @Success 200 {object} response.Base
// @Failure 401 {object} response.Base
// @Failure 404 {object} response.Base
// @Router /v1/cat/{id} [delete]
func (h *CatHandler) DeleteCatData(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	claimUser, ok := middleware.GetClaimsUser(r.Context()).(jwt.MapClaims)
	if !ok {
		log.Warn().Msg("invalid claim jwt")
		err := failure.Unauthorized("invalid claim jwt")
		response.WithError(w, err)
		return
	}
	_, err := uuid.Parse(claimUser["ownerID"].(string))
	if err != nil {
		log.Warn().Msg(err.Error())
		err = failure.Unauthorized("invalid format user_id")
		response.WithError(w, err)
		return
	}
	res, err := h.CatService.DeleteCatData(r.Context(), idStr)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}
