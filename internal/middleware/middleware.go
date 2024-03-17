package middleware

type Middleware interface {
	VerifyJWTToken(tokenString string) (*JWTClaims, error)
	GenerateToken(u User) (string, error)
}
