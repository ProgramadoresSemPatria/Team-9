var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
