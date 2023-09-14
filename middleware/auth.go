package middleware

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/shoelfikar/golang_backend_api/helper"
	"github.com/shoelfikar/golang_backend_api/model"
)

// GenerateJWT is func to generate the token
func GenerateJWT(user *model.User) (string, error) {
	var expMins time.Duration = 2400
	token := jwt.New(jwt.SigningMethodHS256)
	jwtKey := []byte(helper.GetViperEnvVariable("JWT_KEY"))

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * expMins).Unix()
	claims["user_id"] = user.Id
	claims["username"] = user.Username

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		helper.DefaultLoggingError(err.Error())
		return "", err
	}
	return tokenString, nil
}

//IsAuthorized is the func for validating the JWT token
func IsAuthorized(endpoint http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] != nil {
			if r.Header["Authorization"][0] == helper.GetViperEnvVariable("API_KEY") {
				endpoint.ServeHTTP(w,r)
			} else {
				resp := model.ApiResponse{
					Code: http.StatusUnauthorized,
					Message: "UNAUTHORIZED",
					Status: "failed",
				}
				helper.WriteToResponseBody(w,resp)
				return
			}

		} else {
			resp := model.ApiResponse{
				Code: http.StatusUnauthorized,
				Message: "AUTH NOT FOUND",
				Status: "failed",
			}
			helper.WriteToResponseBody(w,resp)
			return
		}
	})
}