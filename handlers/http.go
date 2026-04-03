package handlers

import (
	"encoding/json"
	"net"
	"net/http"
	"time"
	"vantagrid/core"
	"vantagrid/storage"
	"vantagrid/types"
)

type HTTPHandler struct {
	Store *storage.Store
}

func getIP(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

func (h *HTTPHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ip := getIP(r)

	raw := ip + r.UserAgent() + r.URL.Path

	event := types.Event{
		IP:        ip,
		Path:      r.URL.Path,
		Method:    r.Method,
		UserAgent: r.UserAgent(),
		Score:     core.Score(r.UserAgent(), r.Method, r.URL.Path),
		Fingerprint: core.Fingerprint(raw),
		Time:      time.Now().Format(time.RFC3339),
	}

	h.Store.Add(event)

	if r.URL.Path == "/login" && r.Method == "POST" {
		r.ParseForm()
		_ = r.Form.Get("username")
		_ = r.Form.Get("password")
	}

	w.Write([]byte("OK"))
}
