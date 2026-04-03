package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *HTTPHandler) API(w http.ResponseWriter, r *http.Request) {
	data := h.Store.All()
	json.NewEncoder(w).Encode(data)
}
