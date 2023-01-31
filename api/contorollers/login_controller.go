package contorollers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tahahmmcgl/kullanici_api/api/auth"
	"github.com/tahahmmcgl/kullanici_api/api/models"
	"github.com/tahahmmcgl/kullanici_api/api/responses"
	"github.com/tahahmmcgl/kullanici_api/api/utils"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := utils.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnauthorized, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}
func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
func (server *Server) IsAdminUser() (bool, error) {
	user := models.User{}
	err := server.DB.Debug().Model(models.User{}).Where("id = ?", user.ID).Take(&user).Error
	if err != nil {
		return false, err
	}
	if user.Role == 2 {
		return true, nil
	}
	return false, nil
}
