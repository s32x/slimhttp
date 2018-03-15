package slimhttp

import (
	"net/url"
	"testing"
)

func TestURLSigner(t *testing.T) {
	s := NewURLSigner()
	s.secret = "STATIC_SECRET"
	u, err := url.Parse("/playlist-url.m3u8?key1=val1&key2=val2")
	equal(t, err, nil)
	signedURL := s.SignURL(u)
	equal(t, signedURL, "/playlist-url.m3u8?key1=val1&key2=val2&token=9192539f2c3969c05699c3a163f83c642880a77597cfce8f19bfd3f3924bfd06")
	equal(t, err, nil)
}

func TestValidateURLSuccess(t *testing.T) {
	s := NewURLSigner()
	s.secret = "STATIC_SECRET"
	u, err := url.Parse("/playlist-url.m3u8?key1=val1&key2=val2&token=9192539f2c3969c05699c3a163f83c642880a77597cfce8f19bfd3f3924bfd06")
	equal(t, err, nil)
	err = s.ValidateURL(u)
	equal(t, err, nil)
}

func TestValidateURLNoToken(t *testing.T) {
	s := NewURLSigner()
	s.secret = "STATIC_SECRET"
	u, err := url.Parse("/playlist-url.m3u8?key1=val1&key2=val2")
	equal(t, err, nil)
	err = s.ValidateURL(u)
	equal(t, err, ErrNoTokenFound)
}

func TestValidateURLFailure(t *testing.T) {
	s := NewURLSigner()
	s.secret = "STATIC_SECRET"
	u, err := url.Parse("/playlist-url.m3u8?key1=val1&key2=val2&token=DEFINITELYINVALID")
	equal(t, err, nil)
	err = s.ValidateURL(u)
	equal(t, err, ErrInvalidToken)
}
