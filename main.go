package main

import (
	"os"
	"passcript/internal/utils"

	"go.uber.org/zap"
)

func init() {

}

func main() {
	utils.Log().Info("Starting application")
	utils.Log().Info("PUBLIC_KEY set", zap.String("public_key", os.Getenv("PUBLIC_KEY")))
}
