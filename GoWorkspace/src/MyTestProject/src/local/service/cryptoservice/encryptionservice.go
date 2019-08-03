package service

import (
	//"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"

	//"crypto/sha256"
	"fmt"
	//"os"
)

func StartProcessing(messageToBeEncrypted string) {
	keys := generateKeys()
	ciphertext := encryptMessage(keys, []byte(messageToBeEncrypted))
	plainText := decryptMessage(keys, ciphertext)

	fmt.Println(string(plainText))

	GenerateCertificate()

}

func decryptMessage(keys Keys, ciphertext []byte) []byte {
	label := []byte("")
	hash := sha256.New()
	plainText, err := rsa.DecryptOAEP(
		hash,
		rand.Reader,
		keys.privatekey,
		ciphertext,
		label,
	)
	if err != nil {
		fmt.Println(err)

	}
	return plainText
}

func encryptMessage(keys Keys, message []byte) []byte {
	label := []byte("")
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(
		hash,
		rand.Reader,
		keys.publickey,
		message,
		label,
	)
	if err != nil {
		fmt.Println(err)
	}
	return ciphertext
}

func generateKeys() Keys {
	fmt.Println("Private Key : ")
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Private Key : ", privateKey)
	publicKey := &privateKey.PublicKey

	keys := Keys{publickey: publicKey, privatekey: privateKey}

	return keys
	// raulPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	// if err != nil {
	// 	fmt.Println(err.Error)
	// 	//os.Exit(1)
	// }
	// raulPublicKey := &raulPrivateKey.PublicKey
}
