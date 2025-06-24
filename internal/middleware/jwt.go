package middleware

type IJWTService interface {
	GenerateToken(userID int64, email string) (string, error)
}
