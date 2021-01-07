package gflr

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jhlq/go-flare/erc20"

	"context"
	"fmt"
	"math/big"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

var host = "https://costone.flare.network/ext/bc/C/rpc" //old: "http://coston.flare.network:9650/ext/bc/C/rpc"

func InputHidden(prompt string) (string, error) {
	fmt.Print(prompt)
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
func From18zToFloat(amount *big.Int) *big.Float {
	b, _ := new(big.Int).SetString("1000000000000000000", 10)
	f := new(big.Float).SetInt(amount)
	g := new(big.Float).SetInt(b)
	z := new(big.Float).Quo(f, g)
	return z
}
func Send(secret, address string, amount float64) (string, error) {
	client, err := ethclient.Dial(host)
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
	value := FloatTo18z(amount)
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

	client.Close()
	return signedTx.Hash().Hex(), nil
}
func SendERC20(secret, tokenContract, address string, amount float64) (string, error) {
	client, err := ethclient.Dial(host)
	if err != nil {
		return "", err
	}

	tcaddress := common.HexToAddress(tokenContract)
	instance, err := erc20.NewErc20(tcaddress, client)
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
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	addre := common.HexToAddress(address)
	a := FloatTo18z(amount)
	tx, err := instance.Transfer(auth, addre, a)
	if err != nil {
		return "", err
	}
	client.Close()
	return tx.Hash().Hex(), nil
}

func Balance(address string) (*big.Int, error) {
	client, err := ethclient.Dial(host)
	if err != nil {
		return nil, err
	}
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
func BalanceERC20(tokenContract, address string) (*big.Int, error) {
	client, err := ethclient.Dial(host)
	if err != nil {
		return nil, err
	}
	tcaddress := common.HexToAddress(tokenContract)
	instance, err := erc20.NewErc20(tcaddress, client)
	if err != nil {
		return nil, err
	}
	account := common.HexToAddress(address)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
		return nil, err
	}
	return bal, nil
}
