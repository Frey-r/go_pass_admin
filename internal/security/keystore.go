package security

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
	"runtime"
)

type KeyStore struct {
	privateKeyPath string
	publicKeyPath  string
}

func NewKeyStore() (*KeyStore, error) {
	// Obtener el directorio home del usuario
	homeDir, err := getHomeDir()
	if err != nil {
		return nil, err
	}

	// Crear directorio para las claves
	keyDir := filepath.Join(homeDir, ".passadmin", "keys")
	if err := os.MkdirAll(keyDir, 0700); err != nil {
		return nil, err
	}

	return &KeyStore{
		privateKeyPath: filepath.Join(keyDir, "private.pem"),
		publicKeyPath:  filepath.Join(keyDir, "public.pem"),
	}, nil
}

func (ks *KeyStore) SaveKeys(privateKey *rsa.PrivateKey) error {
	// Guardar llave privada
	if err := ks.savePrivateKey(privateKey); err != nil {
		return err
	}

	// Guardar llave pública
	if err := ks.savePublicKey(&privateKey.PublicKey); err != nil {
		return err
	}

	return nil
}

func (ks *KeyStore) LoadPrivateKey() (*rsa.PrivateKey, error) {
	data, err := os.ReadFile(ks.privateKeyPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func (ks *KeyStore) savePrivateKey(key *rsa.PrivateKey) error {
	bytes := x509.MarshalPKCS1PrivateKey(key)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: bytes,
	}

	// Crear archivo con permisos restrictivos
	file, err := os.OpenFile(ks.privateKeyPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(file, block)
}

func (ks *KeyStore) savePublicKey(key *rsa.PublicKey) error {
	bytes := x509.MarshalPKCS1PublicKey(key)
	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: bytes,
	}

	file, err := os.OpenFile(ks.publicKeyPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(file, block)
}

func (ks *KeyStore) KeysExist() bool {
	_, err := os.Stat(ks.privateKeyPath)
	return !os.IsNotExist(err)
}

func getHomeDir() (string, error) {
	if runtime.GOOS == "windows" {
		return os.Getenv("USERPROFILE"), nil
	}
	return os.Getenv("HOME"), nil
}
