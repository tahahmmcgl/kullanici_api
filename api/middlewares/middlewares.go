package middlewares

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/tahahmmcgl/kullanici_api/api/auth"
	"github.com/tahahmmcgl/kullanici_api/api/responses"
)

//setmiddlewarejson json formatında veri döndürmek için kullanılır
func SetMilddlewareJson(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

//setmiddlewareauthentication token doğrulama için kullanılır
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}

//setmiddlewarelogger loglama için kullanılır
func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("middleware çalıştı")
		log.Println("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}
