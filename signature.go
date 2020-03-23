package permata

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"unicode"

	"github.com/juju/errors"
)

func canonicalize(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

// GenerateSignature generate SHA-256 HMAC signature
func GenerateSignature(staticKey, apiKeyOrToken, timestamp, requestBody string) (signature string, strToSign string, err error) {
	canonicalReqBody := canonicalize(requestBody)

	strToSign =
		apiKeyOrToken + ":" +
			timestamp + ":" +
			canonicalReqBody

	mac := hmac.New(sha256.New, []byte(staticKey))
	if _, err = mac.Write([]byte(strToSign)); err != nil {
		return "", strToSign, errors.Trace(err)
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), strToSign, nil
}
