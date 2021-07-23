package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"hash"
)

// ShaType is a type to store sha encryption method
type ShaType int

const (
	// Sha1 type
	Sha1 ShaType = 1
	// Sha256 type
	Sha256 ShaType = 256
)

// GetMD5Hash func to get md5 hash
func GetMD5Hash(message string) (value string) {
	res := md5.Sum([]byte(message))
	return hex.EncodeToString(res[:])
}

// EncryptSha util function to encrpty sha
func EncryptSha(message string, shaType ShaType) string {
	var sha hash.Hash
	switch shaType {
	case Sha256:
		sha = sha256.New()
		break
	default:
		sha = sha1.New()
		break
	}
	sha.Write([]byte(message))
	// The argument to `Sum` can be used to append
	digest := sha.Sum(nil)

	return hex.EncodeToString(digest)
}

// EncryptAes util function to encrypt aes
func EncryptAes(message, salt string) (value string, err error) {
	text := []byte(message)
	key := []byte(salt)

	c, err := aes.NewCipher(key)
	if err != nil {
		return value, err
	}

	// Galois/Counter Mode, is a mode of operation for symmetric key cryptographic block ciphers
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return value, err
	}

	// creates a new byte array the size of the nonce which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())

	// populates the nonce with a cryptographically random sequence (intentionally turned of)
	/* if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	} */

	value = hex.EncodeToString(gcm.Seal(nonce, nonce, text, nil))
	return
}

// DecryptAes util function to encrypt aes
func DecryptAes(cipherMessage, salt string) (value string, err error) {
	ciphertext, _ := hex.DecodeString(cipherMessage)
	key := []byte(salt)

	c, err := aes.NewCipher(key)
	if err != nil {
		return value, err
	}

	// Galois/Counter Mode, is a mode of operation for symmetric key cryptographic block ciphers
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return value, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return value, errors.New(ErrMessageCipherLessThanNonce)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	res, err := gcm.Open(nil, nonce, ciphertext, nil)

	value = string(res)

	return value, err
}
