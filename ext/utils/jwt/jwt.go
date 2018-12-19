package jwt

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	uuid "github.com/satori/go.uuid"

	"github.com/hieuphq/califit-be/goa/app"
)

type JWTSecurityFunc func() *goa.JWTSecurity

func GetUserIDFromCtx(ctx context.Context) (int, error) {
	token := jwt.ContextJWT(ctx)
	claims, ok := token.Claims.(jwtgo.MapClaims)
	if !ok {
		return 0, jwt.ErrJWTError("unsupported claims shape")

	}

	idIF, ok := claims["user.id"]
	if !ok {
		return 0, errors.New("missing user.id in jwt claims")
	}

	idStr, ok := idIF.(string)
	if !ok {
		return 0, errors.New("id in ctx not is a string")
	}

	return strconv.Atoi(idStr)
}

func NewJWTMiddleware(jwtSecurity JWTSecurityFunc, keyPath string) (goa.Middleware, error) {
	validationHandler, _ := goa.NewMiddleware(func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		token := jwt.ContextJWT(ctx)
		_, ok := token.Claims.(jwtgo.MapClaims)
		if !ok {
			return jwt.ErrJWTError("unsupported claims shape")
		}
		return nil
	})
	keys, err := LoadJWTPublicKeys(keyPath)
	if err != nil {
		return nil, err
	}
	return jwt.New(keys, validationHandler, app.NewJWTSecurity()), nil
}

// LoadJWTPublicKeys loads PEM encoded RSA public keys used to validata and decrypt the JWT.
func LoadJWTPublicKeys(keyPath string) (*rsa.PublicKey, error) {
	keyFile, err := filepath.Glob(keyPath)
	if err != nil {
		return nil, err
	}
	pem, err := ioutil.ReadFile(keyFile[0])
	if err != nil {
		return nil, err
	}
	key, err := jwtgo.ParseRSAPublicKeyFromPEM([]byte(pem))
	if err != nil {
		return nil, fmt.Errorf("failed to load key %s: %s", keyFile, err)
	}

	return key, nil
}

// LoadJWTPrivateKeys loads PEM encoded RSA private key.
func LoadJWTPrivateKey(keyPath string) (*rsa.PrivateKey, error) {
	keyFile, err := filepath.Glob(keyPath)
	if err != nil {
		return nil, err
	}
	if keyFile == nil {
		return nil, fmt.Errorf("no file match from filepath.Glod() for key path: %s", keyPath)
	}
	pem, err := ioutil.ReadFile(keyFile[0])
	if err != nil {
		return nil, err
	}

	key, err := jwtgo.ParseRSAPrivateKeyFromPEM([]byte(pem))
	if err != nil {
		return nil, fmt.Errorf("failed to load key %s: %s", keyFile, err)
	}
	return key, nil
}

func CreateDriverJWTToken(driverID, userID int) (string, error) {
	token := GenerateDriverJWT(driverID, userID)
	privKey, err := LoadJWTPrivateKey(os.Getenv("JWT_PRIVATE_KEY"))
	if err != nil {
		return "", err
	}

	return token.SignedString(privKey)
}

func GenerateDriverJWT(driverID, userID int) *jwtgo.Token {
	token := jwtgo.New(jwtgo.SigningMethodRS512)
	oneMonth := time.Now().Add(time.Duration(24*30) * time.Hour).Unix()
	uuid, _ := uuid.NewV4()
	token.Claims = jwtgo.MapClaims{
		"iss":       "Issuer",               // who creates the token and signs it
		"aud":       "Audience",             // to whom the token is intended to be sent
		"exp":       oneMonth,               // time when the token will expire (10 minutes from now)
		"jti":       uuid.String(),          // a unique identifier for the token
		"iat":       time.Now().Unix(),      // when the token was issued/created (now)
		"nbf":       2,                      // time before which the token is not yet valid (2 minutes ago)
		"sub":       "llra api token",       // the subject/principal is whom the token is about
		"user.id":   strconv.Itoa(userID),   // user id - not a standard claim
		"driver.id": strconv.Itoa(driverID), // driver id - not a standard claim
		"scopes":    "api:access",           // token scope - not a standard claim
	}
	return token
}

func CreateJWTToken(userID int) (string, error) {
	token := GenerateJWT(userID)
	privKey, err := LoadJWTPrivateKey(os.Getenv("JWT_PRIVATE_KEY"))
	if err != nil {
		return "", err
	}

	return token.SignedString(privKey)
}

func GenerateJWT(userID int) *jwtgo.Token {
	token := jwtgo.New(jwtgo.SigningMethodRS512)
	oneMonth := time.Now().Add(time.Duration(24*30) * time.Hour).Unix()
	uuid, _ := uuid.NewV4()
	token.Claims = jwtgo.MapClaims{
		"iss":     "Issuer",                          // who creates the token and signs it
		"aud":     "Audience",                        // to whom the token is intended to be sent
		"exp":     oneMonth,                          // time when the token will expire (10 minutes from now)
		"jti":     uuid.String(),                     // a unique identifier for the token
		"iat":     time.Now().Unix(),                 // when the token was issued/created (now)
		"nbf":     2,                                 // time before which the token is not yet valid (2 minutes ago)
		"sub":     "llra api token",                  // the subject/principal is whom the token is about
		"user.id": strconv.Itoa(userID),              // user email - not a standard claim
		"scopes":  "api:access company_account:true", // token scope - not a standard claim
	}
	return token
}
