package decnet

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
)

type Key struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func GenerateKey() (*Key, error) {
	k := new(Key)

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return k, err
	}

	k.publicKey = &privateKey.PublicKey
	k.privateKey = privateKey

	return k, nil
}

func (k Key) PublicKeyToPemString() string {
	return string(
		pem.EncodeToMemory(
			&pem.Block{
				Type:  "RSA PUBLIC KEY",
				Bytes: x509.MarshalPKCS1PublicKey(k.publicKey),
			},
		),
	)
}

func (k Key) PrivateKeyToPemString() string {
	return string(
		pem.EncodeToMemory(
			&pem.Block{
				Type:  "RSA PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(k.privateKey),
			},
		),
	)
}

func (k *Key) Decrypt(encryptedMessage []byte) ([]byte, error) {

	return rsa.DecryptOAEP(
		sha512.New(),
		rand.Reader,
		k.privateKey,
		pemStringToCipher(encryptedMessage),
		nil,
	)

}

func Encrypt(publicKey *rsa.PublicKey, plainText []byte) ([]byte, error) {
	if publicKey == nil {
		return plainText, nil
	}
	cipher, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, publicKey, plainText, nil)
	if err != nil {
		return nil, err
	}
	return cipherToPemString(cipher), nil
}

func pemStringToCipher(encryptedMessage []byte) []byte {
	b, _ := pem.Decode(encryptedMessage)
	return b.Bytes
}

func convertBytesToPublicKey(keyBytes []byte) (*rsa.PublicKey, error) {
	var err error

	block, _ := pem.Decode(keyBytes)
	blockBytes := block.Bytes
	ok := x509.IsEncryptedPEMBlock(block)

	if ok {
		blockBytes, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}

	publicKey, err := x509.ParsePKCS1PublicKey(blockBytes)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

func cipherToPemString(cipher []byte) []byte {
	return pem.EncodeToMemory(
		&pem.Block{
			Type:  "MESSAGE",
			Bytes: cipher,
		},
	)
}
