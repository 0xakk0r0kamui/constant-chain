package wallet

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

type AES struct {
}

// create a salt
func (aesAl AES) deriveKey(passPhrase string, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 8)
		rand.Read(salt)
	}
	return pbkdf2.Key([]byte(passPhrase), salt, 1000, 32, sha256.New), salt
}

// encrypt with AES
func (aesAl AES) Encrypt(passphrase string, plaintext []byte) (string, error) {
	key, salt := aesAl.deriveKey(passphrase, nil)
	iv := make([]byte, 12)
	// http://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-38d.pdf
	// Section 8.2
	rand.Read(iv)
	b, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(b)
	data := aesgcm.Seal(nil, iv, plaintext, nil)
	return hex.EncodeToString(salt) + "-" + hex.EncodeToString(iv) + "-" + hex.EncodeToString(data), err
}

// decrypt with AES
func (aesAl AES) Decrypt(passPhrase, cipherText string) ([]byte, error) {
	arr := strings.Split(cipherText, "-")
	salt, err := hex.DecodeString(arr[0])
	if err != nil {
		return nil, err
	}
	iv, err := hex.DecodeString(arr[1])
	if err != nil {
		return nil, err
	}
	data, err := hex.DecodeString(arr[2])
	if err != nil {
		return nil, err
	}
	key, _ := aesAl.deriveKey(passPhrase, salt)
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(b)
	if err != nil {
		return nil, err
	}
	data, err = aesgcm.Open(nil, iv, data, nil)
	if err != nil {
		return nil, err
	}
	return data, err
}
