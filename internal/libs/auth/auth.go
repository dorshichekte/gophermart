package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

func New(accessSecret string) Auth {
	return &auth{accessSecret: accessSecret}
}

func (auth *auth) Generate(userID int) (*TokenPair, error) {
	claims := newClaims(userID)
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(auth.accessSecret))
	if err != nil {
		return nil, err
	}

	authData := &TokenPair{
		AccessToken: accessToken,
	}

	return authData, nil
}

func (auth *auth) ParseAccessToken(accessToken string) (*UserData, error) {
	parseAccessTokenFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errUnexpectedSigningMethod
		}

		return []byte(auth.accessSecret), nil
	}

	claims := &claims{}
	_, err := jwt.ParseWithClaims(accessToken, claims, parseAccessTokenFunc)
	if err != nil {
		return nil, err
	}

	isValidErr := claims.valid()
	if isValidErr != nil {
		return nil, isValidErr
	}

	return &claims.userData, nil
}
