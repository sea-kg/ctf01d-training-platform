package handler

import (
	"net/http"
	"runtime"
	"time"

	"ctf01d/internal/helper"
)

var (
	version   = "dev"
	buildTime = time.Now()
)

func (h *Handler) GetVersion(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"version":    version,
		"golang":     runtime.Version(),
		"build_time": buildTime.Format(time.RFC822Z),
	}
	helper.RespondWithJSON(w, http.StatusOK, res)
}
