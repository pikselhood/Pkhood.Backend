package accountcontroller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/api/idtoken"
	"log"
	"net/http"
	"pkhood-backend/src/domain/account"
	"pkhood-backend/src/repository"
	"pkhood-backend/src/scripts/token"
)

type accountAuthenticateResponse struct {
	Success bool   `json:"success"`
	PToken  string `json:"pToken"`
}

// Authenticate godoc
//
//		@Summary		Authenticate an account
//		@Description	get string by ID
//		@Accept			json
//		@Produce		json
//	 	@Tags      		account
//		@Param			googleIdToken	path		string	true	"Google Id Token"
//		@Success		200				{object}	accountAuthenticateResponse
//		@Router			/account/authorize/{googleIdToken} [get]
func Authenticate(e *echo.Group, m *repository.MongoClient[account.Account]) {

	var (
		err    error
		result *idtoken.Payload
		jwt    string
	)

	e.GET("/account/authorize/:googleIdToken", func(c echo.Context) error {

		googleIdToken := c.Param("googleIdToken")

		if result, err = token.VerifyGoogleIdToken(googleIdToken); err != nil {
			return err
		}

		filter := bson.D{{"email", result.Claims["email"]}}

		u, err := m.Get(filter)
		if err != nil {
			return err
		}

		if jwt, err = token.GenerateJwt(u.Id, u.Email); err != nil {
			log.Printf(err.Error())
			return err
		}

		return c.JSON(http.StatusOK, accountAuthenticateResponse{
			PToken:  jwt,
			Success: true,
		})
	})
}

type accountCreateRequest struct {
	GoogleIdToken string `json:"googleIdToken"`
}

// Create godoc
//
//		@Summary		Create an account
//		@Description	create account
//		@Accept			json
//		@Produce		json
//	 	@Tags      		account
//		@Param			googleIdToken	body		accountCreateRequest	true	"Google Id Token"
//		@Success		200				{object}	accountAuthenticateResponse
//		@Router			/account/create [post]
func Create(e *echo.Group, m *repository.MongoClient[account.Account]) {

	var (
		result *idtoken.Payload
		jwt    string
	)

	e.POST("/account/create", func(c echo.Context) error {

		req := accountCreateRequest{}

		err := json.NewDecoder(c.Request().Body).Decode(&req)

		if result, err = token.VerifyGoogleIdToken(req.GoogleIdToken); err != nil {
			return err
		}

		u := account.New(result)

		_, err = m.Insert(u)
		if err != nil {
			return err
		}

		if jwt, err = token.GenerateJwt(u.Id, u.Email); err != nil {
			log.Printf(err.Error())
			return err
		}

		return c.JSON(http.StatusOK, accountAuthenticateResponse{
			PToken:  jwt,
			Success: true,
		})
	})
}
