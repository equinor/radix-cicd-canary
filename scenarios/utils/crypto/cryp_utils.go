package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

// SHA256HMAC computes the GitHub SHA1 HMAC.
func SHA256HMAC(salt, message []byte) string {
	// GitHub creates a SHA256 HMAC, where the key is the GitHub secret and the
	// message is the JSON body.
	digest := hmac.New(sha256.New, salt)
	digest.Write(message)
	sum := digest.Sum(nil)
	return fmt.Sprintf("sha256=%x", sum)
}
