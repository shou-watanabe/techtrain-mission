package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	myError "go-gacha-server/src/core/error"
	"go-gacha-server/src/presen/request"
	"go-gacha-server/src/presen/response"
	"go-gacha-server/src/usecase"

	"go.uber.org/zap"
)

type GachaHandler interface {
	Draw(http.ResponseWriter, *http.Request)
}

type gachaHandler struct {
	gachaUsecase usecase.GachaUsecase
}

func NewGachaHandler(gachaUsecase usecase.GachaUsecase) GachaHandler {
	return &gachaHandler{gachaUsecase: gachaUsecase}
}

func (gh *gachaHandler) Draw(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		zap.Error(myError.ErrMethodNotFound)
		return
	}

	token := r.Header.Get("X-Token")
	if token == "" {
		zap.Error(myError.ErrTokenNotFound)
	}

	var req request.GachaDrawRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		zap.Error(err)
	}

	characters, err := gh.gachaUsecase.Draw(r.Context(), req.Times, token)
	if err != nil {
		zap.Error(err)
	}

	var results []response.GachaResult
	for _, ce := range characters {
		result := response.GachaResult{
			CharacterID: strconv.Itoa(ce.Id),
			Name:        ce.Name,
			IconURL:     ce.IconURL,
			Rarity:      ce.Rarity,
		}
		results = append(results, result)
	}

	res := response.GachaDrawResponse{
		Results: results,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	je := json.NewEncoder(w)
	if err := je.Encode(res); err != nil {
		zap.Error(err)
	}
}
