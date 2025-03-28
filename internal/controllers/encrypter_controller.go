package controllers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"passcript/internal/utils"

	"go.uber.org/zap"
)

func GenerateKeys() (rsa.PrivateKey, error) {
	utils.Log().Info("Generating keys")
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return rsa.PrivateKey{}, err
	}
	return *privateKey, nil
}

func Encrypter(publicKey *rsa.PublicKey, object string) []byte {
	utils.Log().Info("Encrypting object")
	encrypted, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(object), nil)
	if err != nil {
		utils.Log().Error("Error encrypting object", zap.Error(err))
		return nil
	}
	return encrypted
}

func Decrypter(privateKey *rsa.PrivateKey, object []byte) string {
	utils.Log().Info("Decrypting object")
	decrypted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, object, nil)
	if err != nil {
		utils.Log().Error("Error decrypting object", zap.Error(err))
		return ""
	}
	return string(decrypted)
}

/*
USE
	privateKey, err := controllers.GenerateKeys()
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.Public().(*rsa.PublicKey)
	encrypted := controllers.Encrypter(publicKey, OBJECT)
	decrypted := controllers.Decrypter(&privateKey, encrypted)
*/
