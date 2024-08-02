package handlers

import (
	"net/http"

	"ctf01d/internal/app/server"
	"ctf01d/pkg/avatar"
)

func (h *Handlers) UniqueAvatar(w http.ResponseWriter, r *http.Request, username string, params server.UniqueAvatarParams) {
	xMax := 100
	yMax := 100
	if params.Max != nil {
		xMax = *params.Max
		yMax = *params.Max
	}
	blockSize := 20
	if params.BlockSize != nil {
		blockSize = *params.BlockSize
	}
	steps := 8
	if params.Steps != nil {
		steps = *params.Steps
	}

	imageBytes := avatar.GenerateAvatar(username, xMax, yMax, blockSize, steps)

	w.Header().Set("Content-Type", "image/png")
	w.Write(imageBytes)
}
