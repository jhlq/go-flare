package gflr

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"context"
	"fmt"
	"math/big"
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

func FloatTo18z(amount float64) *big.Int {
	s := fmt.Sprintf("%f", amount)
	a := strings.Split(s, ".")
	for i := 0; i < 18; i++ {
		if len(a) > 1 && len(a[1]) > i {
			a[0] += string(a[1][i])
		} else {
			a[0] += "0"
		}
	}
	b, ok := new(big.Int).SetString(a[0], 10)
	if !ok {
		panic("Could not set big.Int string for value " + s)
	}
	return b
}
func Send(secret, address string, amount float64) (string, error) {
	client, err := ethclient.Dial("http://coston.flare.network:9650/ext/bc/C/rpc")
	if err != nil {
		return "", err
	}

	privateKey, err := crypto.HexToECDSA(secret)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}
	/*s := fmt.Sprintf("%f", amount*math.Pow(10, 18))
	s = strings.Split(s, ".")[0]
	value, ok := new(big.Int).SetString(s, 10)*/
	value := FloatTo18z(amount)
	if !ok {
		panic("Couldn't set big.Int string.")
	}
	gasLimit := uint64(210000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	toAddress := common.HexToAddress(address)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	//fmt.Printf("tx sent: %s, to address: %s\n", signedTx.Hash().Hex(), address)

	client.Close()
	return signedTx.Hash().Hex(), nil
}
