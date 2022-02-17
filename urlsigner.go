package slimhttp

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/url"

	uuid "github.com/satori/go.uuid"
)

var (
	TokenKey        = "token"
	ErrNoTokenFound = errors.New("No token on request")
	ErrInvalidToken = errors.New("Invalid token recieved")
)

// URLSigner contains the secret used to sign URLs
type URLSigner struct{ secret string }

// NewURLSigner creates a new URLSigner instance
func NewURLSigner() *URLSigner {
	return &URLSigner{secret: uuid.NewV4().String()}
}

// NewURLSignerFromSecret creates a new URSigner instance
// from a passed secret
func NewURLSignerFromSecret(secret string) *URLSigner {
	return &URLSigner{secret: secret}
}

// SignURL takes a url and securely signs it using sha256
// and the secret on the URLSigner and returns the token
func (s *URLSigner) SignURL(u *url.URL) string {
	// Create a new query with the token
	query := u.Query()
	query.Add(TokenKey, s.generateToken(u.String()))

	// Encode and return the URL
	u.RawQuery = query.Encode()
	return u.String()
}

// ValidateURL takes a url
func (s *URLSigner) ValidateURL(u *url.URL) error {
	oldToken := u.Query().Get(TokenKey)

	// Verify a token exists on the request URL
	if oldToken == "" {
		return ErrNoTokenFound
	}

	// Clean the token off the old URL
	query := u.Query()
	query.Del(TokenKey)
	u.RawQuery = query.Encode()

	// Check the old token against a newly generated one
	if oldToken != s.generateToken(u.String()) {
		return ErrInvalidToken
	}
	return nil
}

// generateToken signs the passed URL and add the token to the
// parsed url.URL
func (s *URLSigner) generateToken(rawurl string) string {
	mac := hmac.New(sha256.New, []byte(s.secret))
	mac.Write([]byte(rawurl))
	return hex.EncodeToString(mac.Sum(nil))
}
