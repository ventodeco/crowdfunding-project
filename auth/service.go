package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userId int) (string, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("vento_deco_s3cr3t_k3y")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{}

	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil

}
