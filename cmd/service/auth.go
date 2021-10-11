package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jhonnycpp/lds-0221-common/cmd/errors"
)

type AuthService struct {
	secretKey string
	issuer    string
}

type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewAuthService() *AuthService {
	return &AuthService{
		// Chave de criptografia do token
		secretKey: "8aJjqTSGwZ8dWES2Ir66O5k4bkrwc6SbWcBDxTJisHCFMU7WD3SRz4LQe3C60UYMyZPktaJdiMiuigHlZa4OUGfzAiw04JCrjSSTscxQpWCN2J0Rkuh51Cu2lgsfqwNQEgkx9YEouG3nQs62ddyUgiNHF6nxr1DxXokXDoBzw5cCnY5H2sdYQqgXnJJWBNIY2sw6YyuJsMMXlHfe182HpZfSjook2y9uguq5ibERoRoPaJuQImZyC4HljwLbjLLG",
		// Identificação de geração de token
		issuer: "orgen-login-api",
	}
}

func (s *AuthService) GenerateToken(email string) (string, error) {
	claim := &Claim{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // Tempo de expiração do token
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *AuthService) ValidateToken(token string) (email string, err error) {
	var claim Claim
	_, err = jwt.ParseWithClaims(token, &claim, s.parse)

	if err != nil {
		return "", err
	}

	return claim.Email, err
}

func (s *AuthService) parse(token *jwt.Token) (interface{}, error) {
	if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
		return nil, errors.InvalidAuth
	}

	return []byte(s.secretKey), nil
}
