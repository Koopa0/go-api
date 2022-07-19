package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/koopa0/go-api/internal/helpers"
	"github.com/koopa0/go-api/internal/models"
	"github.com/pascaldekloe/jwt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var validUser = models.User{
	ID:       10,
	Email:    "koopa@go.com",
	Password: "$2a$12$yiKMqrNvbxzMtL6HHcAyb.k7aeircJMP2hr9Xpvyw3to6OfhnMujq",
}

func (app *application) Broker(w http.ResponseWriter, r *http.Request) {

	payload := helpers.JsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = helpers.WriteJSON(w, http.StatusOK, payload)
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {

	var cred models.Credentials

	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		helpers.ErrorJSON(w, errors.New("unauthorized"))
		return
	}

	hashedPassword := validUser.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(cred.Password))
	if err != nil {
		helpers.ErrorJSON(w, errors.New("unauthorized"))
		return
	}

	var claims jwt.Claims
	claims.Subject = fmt.Sprint(validUser.ID)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "mydomain.com"
	claims.Audiences = []string{"mydomain.com"}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(app.config.jwt.secret))
	if err != nil {
		helpers.ErrorJSON(w, errors.New("error signing"))
		return
	}

	helpers.WriteJSON(w, http.StatusOK, string(jwtBytes))
}
