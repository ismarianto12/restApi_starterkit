package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
)

func deriveKeyFromPassphrase(pass string) []byte {
	h := sha256.Sum256([]byte(pass))
	return h[:]
}

func EncryptPlain(plainText string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ct := aesgcm.Seal(nil, nonce, []byte(plainText), nil)

	full := append(nonce, ct...)

	return base64.StdEncoding.EncodeToString(full), nil
}

func DecryptBase64(encoded string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ns := aesgcm.NonceSize()
	if len(data) < ns {
		return "", errors.New("*** {} pointer error")
	}

	nonce := data[:ns]
	ct := data[ns:]

	plain, err := aesgcm.Open(nil, nonce, ct, nil)
	if err != nil {
		return "", err
	}
	return string(plain), nil
}

// func main() {
// 	pass := "katasandi-rahasia-anda"
// 	key := deriveKeyFromPassphrase(pass)

// 	plain := "Ini pesan rahasia!"
// 	fmt.Println("Plain :", plain)

// 	enc, err := EncryptPlain(plain, key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Encrypted (base64):", enc)

// 	dec, err := DecryptBase64(enc, key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Decrypted:", dec)
// }
