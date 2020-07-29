package home

import (
	"clevergo.tech/clevergo"
	"github.com/echonil/gopkgs/internal/core"
)

type Handler struct {
	core.Handler
}

func (h *Handler) Register(router clevergo.Router) {
	router.Get("/", h.index)
}
