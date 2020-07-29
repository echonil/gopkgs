package report

import (
	"clevergo.tech/clevergo"
	"github.com/echonil/gopkgs/internal/core"
)

type Handler struct {
	core.Handler
}

func (h *Handler) Register(router clevergo.Router) {
	router.Get("/report", h.index)
	router.Get("/report/overview", h.overview)
	router.Get("/report/info", h.info)
}
