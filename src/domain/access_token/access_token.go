package access_token

import (
	"bookstore_oauth/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (token *AccessToken) Validate() *errors.RestErr {
	token.AccessToken = strings.TrimSpace(token.AccessToken)
	if token.AccessToken == "" {
		return errors.NewBadRequestError("Invalid Access Token ID")
	}
	if token.UserId <= 0 {
		return errors.NewBadRequestError("Invalid User ID")
	}
	if token.ClientId <= 0 {
		return errors.NewBadRequestError("Invalid Client ID")
	}
	if token.Expires <= 0 {
		return errors.NewBadRequestError("Invalid Expiration Time")
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
