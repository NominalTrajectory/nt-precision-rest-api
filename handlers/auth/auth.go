package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/NominalTrajectory/nt-precision-rest-api/database"
	"github.com/NominalTrajectory/nt-precision-rest-api/models"
	"github.com/NominalTrajectory/nt-precision-rest-api/utils"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtSecretKey string = os.Getenv("JWT_SECRET_KEY")

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), utils.ErrToStatusCode(err))
		return
	}

	pwd, _ := bcrypt.GenerateFromPassword([]byte(user.Pwd), 14)
	user.Pwd = string(pwd)

	if err := database.DB.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), utils.ErrToStatusCode(err))
	} else {
		utils.WriteJSONResult(w, user)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var credentials models.Credentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), utils.ErrToStatusCode(err))
		return
	}

	if err := database.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "This combination of email and password is incorrect", utils.ErrToStatusCode(err))
			return
		} else {
			http.Error(w, err.Error(), utils.ErrToStatusCode(err))
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(credentials.Pwd)); err != nil {
		http.Error(w, "This combination of email and password is incorrect", utils.ErrToStatusCode(err))
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // expires after 24 hours, move to a config file
	})

	token, err := claims.SignedString([]byte(jwtSecretKey))
	if err != nil {
		http.Error(w, "Could not login, please try later", utils.ErrToStatusCode(err))
		return
	}

	utils.WriteJSONResult(w, models.Token{Token: token})
}
