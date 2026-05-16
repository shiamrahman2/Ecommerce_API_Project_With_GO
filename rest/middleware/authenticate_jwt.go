package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecomerce/config"
	"encoding/base64"
	"net/http"
	"strings"
)

func AuthenticationJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		headerArr := strings.Split(header, " ")

		if len(headerArr) != 2 && headerArr[0] != "Bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		accessToken := headerArr[1]

		jwtParts := strings.Split(accessToken, ".")

		if len(jwtParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		jwtHeader := jwtParts[0]
		jwtPayload := jwtParts[1]
		signature := jwtParts[2]

		message := jwtHeader + "." + jwtPayload

		cnf := config.GetConfig()

		h := hmac.New(sha256.New, []byte(cnf.JwtSecretKey))
		h.Write([]byte(message))
		hash := h.Sum(nil)

		newSignature := base64.URLEncoding.
			WithPadding(base64.NoPadding).
			EncodeToString(hash)

		if signature != newSignature {
			http.Error(w, "Unauthorized, Unknown Access", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
