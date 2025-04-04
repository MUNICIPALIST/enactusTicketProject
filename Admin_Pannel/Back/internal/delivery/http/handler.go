// internal/delivery/http/handler.go
package http

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
	"ticket-service/internal/usecase"
)

type Handler struct {
	postUC    *usecase.PostUseCase
	prePostUC *usecase.PrePostUseCase
	login     string
	password  string
}

func NewHandler(postUC *usecase.PostUseCase, prePostUC *usecase.PrePostUseCase, login, password string) *Handler {
	return &Handler{
		postUC:    postUC,
		prePostUC: prePostUC,
		login:     login,
		password:  password,
	}
}

func (h *Handler) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Basic" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		payload, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		pair := strings.Split(string(payload), ":")
		if len(pair) != 2 || pair[0] != h.login || pair[1] != h.password {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.postUC.GetPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func (h *Handler) GetPrePosts(w http.ResponseWriter, r *http.Request) {
	prePosts, err := h.prePostUC.GetPrePosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prePosts)
}
