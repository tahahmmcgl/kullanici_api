package contorollers

import (
	"net/http"

	"github.com/tahahmmcgl/kullanici_api/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to this awesome USER API")
}
