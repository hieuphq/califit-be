package token

import (
	"encoding/base64"
	"time"
)

// CreateForgetPasswordToken : create new token from name, email, time when token is create
func CreateForgetPasswordToken(name, email string) (string, time.Time) {
	t := time.Now().AddDate(0, 0, 1)
	parsedTime := t.String()
	key := []byte(parsedTime + ";" + name + email)

	return base64.StdEncoding.EncodeToString(key), t
}
