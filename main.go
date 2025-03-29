package main

import (
	"passcript/internal/controllers"
	"passcript/internal/utils"

	"go.uber.org/zap"
)

var log *zap.Logger = utils.Log()

func init() {
	log.Info("Starting application")
	controllers.MigrateTables()
	if err := controllers.InitializeRSAKeys(); err != nil {
		log.Fatal("Failed to initialize RSA keys:" + err.Error())
	}
	publicKey := controllers.GetPublicKey()
	privateKey := controllers.GetPrivateKey()

	mensaje := "Llaves cargadas con éxito"
	encrypted := controllers.Encrypter(publicKey, mensaje)
	decrypted := controllers.Decrypter(privateKey, encrypted)
	log.Info("Decrypted message: " + decrypted)
}

func main() {
	controllers.CreateUser("eduardo", "123456")
}
