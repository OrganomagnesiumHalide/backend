package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
)

func getPrivateKeyString() string {
	privatekey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privatekey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	return string(privkey_pem)
}

func getPrivateKeyFromString(key string) *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		panic("Unable to decode")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	return privateKey
}
func verifyData(publicKey *rsa.PublicKey, msg string, signature []byte) bool {
	message := []byte(msg)
	hashed := sha256.Sum256(message)
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	return err == nil

}
func signData(key *rsa.PrivateKey, msg string) []byte {
	message := []byte(msg)
	hashed := sha256.Sum256(message)
	returnVal, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashed[:])
	rsa.VerifyPKCS1v15(&key.PublicKey, crypto.SHA256, hashed[:], returnVal)
	if err != nil {
		panic(err)
	}
	return returnVal
}
