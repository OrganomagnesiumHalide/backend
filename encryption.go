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

func signData(key string, msg string) []byte {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		panic("Unable to decode")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	message := []byte(msg)
	hashed := sha256.Sum256(message)
	returnVal, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}
	return returnVal
}
