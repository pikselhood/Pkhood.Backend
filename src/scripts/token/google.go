package token

import (
	"context"
	"google.golang.org/api/idtoken"
	"net/http"
)

var httpClient = &http.Client{}

func VerifyGoogleIdToken(idToken string) (*idtoken.Payload, error) {

	ctx := context.Background()

	validate, err := idtoken.Validate(ctx, idToken, "")
	if err != nil {
		return nil, err
	}

	return validate, nil
}
