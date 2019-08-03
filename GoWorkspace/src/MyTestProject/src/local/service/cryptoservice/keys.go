package service

import (
	"crypto/rsa"
)

type Keys struct {
	publickey  *rsa.PublicKey
	privatekey *rsa.PrivateKey
}
