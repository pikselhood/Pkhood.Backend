package cmd

import (
	"github.com/labstack/echo/v4"
	"pkhood-backend/src/api/usercontroller"
	"pkhood-backend/src/domain/account"
	"pkhood-backend/src/repository"
	"pkhood-backend/src/settings"
)

func ExecuteUser(g *echo.Group) {

	mongoClient := repository.NewMongoClient[account.Account](settings.DatabaseName, settings.CollectionName)

	accountcontroller.Authenticate(g, mongoClient)
	accountcontroller.Create(g, mongoClient)

}
