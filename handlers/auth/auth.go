package auth

import (
	"encoding/json"
	"fmt"
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

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		http.Error(w, "Unable to register, please try again later", utils.ErrToStatusCode(err))
		return
	}

	user.Password = hashedPassword

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
			http.Error(w, "This combination of email and password is incorrect", http.StatusUnauthorized)
			return
		} else {
			http.Error(w, err.Error(), utils.ErrToStatusCode(err))
			return
		}
	}

	hashedPassword := user.Password
	passwordFromRequest := credentials.Password

	if err := checkPassword(passwordFromRequest, hashedPassword); err != nil {
		http.Error(w, "This combination of email and password is incorrect", http.StatusUnauthorized)
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

	authCookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24), // expires after 24 hours, move to a config file,
		HttpOnly: true,
	}

	http.SetCookie(w, &authCookie)
	utils.WriteJSONResult(w, "Successfully logged in")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	expiredCookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, &expiredCookie)
	utils.WriteJSONResult(w, "Logged out")
}

func User(w http.ResponseWriter, r *http.Request) {
	authCookie, err := r.Cookie("jwt")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token, err := jwt.ParseWithClaims(authCookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	if err := database.DB.Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	utils.WriteJSONResult(w, user)
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Failed to hash password: %w", err)
	}

	return string(hashedPassword), nil
}

func checkPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
