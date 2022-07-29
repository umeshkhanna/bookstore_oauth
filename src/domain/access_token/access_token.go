package access_token

import (
	"bookstore_oauth/src/utils/crypto_utils"
	"bookstore_oauth/src/utils/errors"
	"fmt"
	"strings"
	"time"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id,omitempty"`
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

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	//used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	//used for client credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (tokenRequest *AccessTokenRequest) Validate() *errors.RestErr {
	switch tokenRequest.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		errors.NewBadRequestError("Invalid grant_type parameter")
	}
	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
