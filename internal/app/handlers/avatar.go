package handlers

import (
	"ctf01d/pkg/avatar"
	"net/http"
)

func (h *Handlers) UniqueAvatar(w http.ResponseWriter, r *http.Request, username string) {
	xMax := 100
	yMax := 100
	blockSize := 20
	steps := 8

	imageBytes := avatar.GenerateAvatar(username, xMax, yMax, blockSize, steps)

	w.Header().Set("Content-Type", "image/png")
	w.Write(imageBytes)
}
