package gflr

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
	"fmt"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)


func InputSecret() (string, error) {
    fmt.Print("Enter secret: ")
    byteSecret, err := terminal.ReadPassword(int(syscall.Stdin))
    fmt.Print("\n")
    if err != nil {
        return "", err
    }

    secret := string(byteSecret)
    return strings.TrimSpace(secret), nil
}


func ToAddress(secret string) (string, error) {
	key, err := crypto.HexToECDSA(secret)
	if err != nil {
		return "", err
	}
	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address, nil
}

