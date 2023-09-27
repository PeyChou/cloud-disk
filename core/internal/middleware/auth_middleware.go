package middleware

import (
	"cloud-disk/core/etc/helper"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte("Unable to authenticate"))
			if err != nil {
				return
			}
			return
		}
		uc, err := helper.AnalyseToken(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte(err.Error()))
			if err != nil {
				return
			}
			return
		}
		r.Header.Set("UserId", string(rune(uc.Id)))
		r.Header.Set("UserName", uc.Name)
		r.Header.Set("UserIdentity", uc.Identity)
		// Passthrough to next handler if you need
		next(w, r)
	}
}
