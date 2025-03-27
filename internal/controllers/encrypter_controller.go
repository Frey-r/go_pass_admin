package controllers

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"passcript/internal/utils"
)

func generateKeys() (rsa.PrivateKey, error) {
	utils.Log().Info("Generating keys")
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return rsa.PrivateKey{}, err
	}
	return *privateKey, nil
}

func Encoder() {
	privateKey, err := generateKeys()
	if err != nil {
		return
	}
	publicKey := privateKey.PublicKey
	fmt.Println(publicKey)

}
