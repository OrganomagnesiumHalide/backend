package main

import (
	"crypto/rand"
	"crypto/rsa"
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
