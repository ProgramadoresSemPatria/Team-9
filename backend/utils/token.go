const (
	SubjectClaim   = "sub"
	ExpiryClaim    = "exp"
	IssuedAtClaim  = "iat"
	NotBeforeClaim = "nbf"
)
func CreateJWTToken(ttl time.Duration, payload interface{}, privateKey string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("could not decode private key: %w", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	now := time.Now().UTC()

	claims := jwt.MapClaims{
		SubjectClaim:   payload,
		ExpiryClaim:    now.Add(ttl).Unix(),
		IssuedAtClaim:  now.Unix(),
		NotBeforeClaim: now.Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return token, nil
}
