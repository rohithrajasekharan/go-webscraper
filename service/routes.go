package service

import (
	"net/http"

	"github.com/rohithrajasekharan/web-scraper/types"
	"github.com/rohithrajasekharan/web-scraper/utils"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /scrape", h.ScrapeWeb)
}

func (h *Handler) ScrapeWeb(w http.ResponseWriter, r *http.Request) {
	var payload types.ScrapeRequest
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	ps, err := h.createOrder(payload.Url)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, ps)
}
