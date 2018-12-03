package token

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	UserToken    = "user"
	SessionToken = "session"
)

// MyClaims customer claims
type MyClaims struct {
	*jwt.StandardClaims
	Kind     string
	UserName string `json:"user_name"`
}

// New create MyClaims instance
func New(kind, userName string) *MyClaims {
	return &MyClaims{
		StandardClaims: &jwt.StandardClaims{},
		Kind:           kind,
		UserName:       userName,
	}
}

// Sign signs the token using the given secret hash
// and returns the string value.
func (t *MyClaims) Sign(secret string) (string, error) {
	return t.SignExpires(secret, 0)
}

// SignExpires signs the token using the given secret hash
// with an expiration date.
func (t *MyClaims) SignExpires(secret string, exp int64) (string, error) {
	t.ExpiresAt = exp
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
	raw, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("sign failed")
	}
	return raw, nil
}

// SecretFunc short type
type SecretFunc func(string) (string, error)

// Parse parse raw to Token
func Parse(raw string, fn SecretFunc) (*MyClaims, error) {

	token, err := jwt.ParseWithClaims(raw, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// validate signature
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		// get secret
		if c, ok := token.Claims.(*MyClaims); ok && c.Valid() == nil {
			secret, err := fn(c.UserName)
			return []byte(secret), err
		}
		return nil, jwt.ValidationError{}
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// validate claims
	if c, ok := token.Claims.(*MyClaims); ok && token.Valid && c.Valid() == nil {
		return c, nil
	}

	return nil, fmt.Errorf("parse failed")
}

// ParseRequest parse request token
func ParseRequest(r *http.Request, fn SecretFunc) (*MyClaims, error) {
	var token = r.Header.Get("Authorization")

	// first we attempt to get the token from the
	// authorization header.
	if len(token) != 0 {
		token = r.Header.Get("Authorization")
		fmt.Sscanf(token, "Bearer %s", &token)
		return Parse(token, fn)
	}

	// then we attempt to get the token from the
	// access_token url query parameter
	token = r.FormValue("access_token")
	if len(token) != 0 {
		return Parse(token, fn)
	}

	// and finally we attempt to get the token from
	// the user session cookie
	cookie, err := r.Cookie("user_sess")
	if err != nil {
		return nil, err
	}
	return Parse(cookie.Value, fn)
}

// CheckCSRF check csrf
func CheckCSRF(r *http.Request, fn SecretFunc) error {
	// get and options requests are always
	// enabled, without CSRF checks.
	switch r.Method {
	case "GET", "OPTIONS":
		return nil
	}

	// parse the raw CSRF token value and validate
	raw := r.Header.Get("X-CSRF-TOKEN")
	_, err := Parse(raw, fn)
	return err
}
