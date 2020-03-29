package signer

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rc4"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

func decryptKeyFile(path, login, password string) ([]byte, []byte, error) {
	h := sha1.New()
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	bb := bytes.Split(b, []byte("\n"))
	salt := bb[0]
	keypassplain := string(salt) + login + password

	_, err = h.Write([]byte(keypassplain))
	if err != nil {
		return nil, nil, err
	}

	RC4Cipher, err := rc4.NewCipher(h.Sum(nil))
	if err != nil {
		return nil, nil, err
	}
	rawKey, err := base64.StdEncoding.DecodeString(string(bb[1])) // временно
	if err != nil {
		return nil, nil, err
	}
	RC4Cipher.XORKeyStream(rawKey, rawKey)

	return rawKey, h.Sum(nil), nil
}

func Sign(pathToKey, login, password string, CSVRaw []byte) ([]byte, error) {
	k, _, err := decryptKeyFile(pathToKey, login, password)
	if err != nil {
		return nil, err
	}
	hashed := sha1.Sum(CSVRaw)
	block, _ := pem.Decode(k)
	privKey, err := parsePrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	SignedBytes, err := rsa.SignPKCS1v15(rand.Reader, privKey.(*rsa.PrivateKey), crypto.SHA1, hashed[:])
	if err != nil {
		return nil, err
	}

	return SignedBytes, nil
}

// Attempt to parse the given private key DER block. OpenSSL 0.9.8 generates
// PKCS#1 private keys by default, while OpenSSL 1.0.0 generates PKCS#8 keys.
// OpenSSL ecparam generates SEC1 EC private keys for ECDSA. We try all three.
// Copy-n-pasted from here https://groups.google.com/forum/#!msg/Golang-nuts/ZZ1Wt9d268Q/gQGh6e-wGAAJ
func parsePrivateKey(der []byte) (crypto.PrivateKey, error) {
	if key, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		return key, nil
	}
	if key, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		switch key := key.(type) {
		case *rsa.PrivateKey, *ecdsa.PrivateKey:
			return key, nil
		default:
			return nil, errors.New("crypto/tls: found unknown private key type in PKCS#8 wrapping")
		}
	}
	if key, err := x509.ParseECPrivateKey(der); err == nil {
		return key, nil
	}

	return nil, errors.New("crypto/tls: failed to parse private key")
}
