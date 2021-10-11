package middleware

type iAuthService interface {
	GenerateToken(email string) (string, error)
	ValidateToken(token string) (string, error)
}
