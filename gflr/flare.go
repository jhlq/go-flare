package gflr

import (
	"crypto/ecdsa"
	"encoding/csv"
	"errors"
	"io/ioutil"
	"os"
	"os/user"
	"regexp"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

var host string

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
func GenerateWallet() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	key := hexutil.Encode(privateKeyBytes)[2:] // fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address, key, nil
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
func Unlock(ks, passphrase string) (string, error) {
	d, err := GetDir()
	if err != nil {
		return "", err
	}
	jsonBytes, err := ioutil.ReadFile(d + "/" + ks)
	if err != nil {
		return "", err
	}
	key, err := keystore.DecryptKey(jsonBytes, passphrase)
	privateKeyBytes := crypto.FromECDSA(key.PrivateKey)
	secret := hexutil.Encode(privateKeyBytes)[2:]
	return string(secret), nil
}
func GetDir() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	homeDirectory := user.HomeDir
	return homeDirectory + "/go-flare-config", nil
}
func Addresses(tag string) (string, error) {
	if tag[0] != "@"[0] {
		return tag, nil
	}
	dir, err := GetDir()
	if err != nil {
		return "", err
	}
	f, err := os.Open(dir + "/addresses.csv")
	if err != nil {
		return "", err
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return "", err
	}
	address := ""
	for i := 0; i < len(records); i++ {
		if records[i][0] == tag {
			address = records[i][1]
			break
		}
	}
	return address, nil
}
func GetHost() (string, error) {
	dir, err := GetDir()
	if err != nil {
		return "", err
	}
	h, err := ioutil.ReadFile(dir + "/host.txt")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(h)), nil
}
func SetHost(h string) error {
	var err error
	if h == "" {
		host, err = GetHost()
		return err
	}
	host = h
	return nil
}
func ValidateAddress(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}
func Float2Int(amount float64, decimals int) *big.Int {
	s := strconv.FormatFloat(amount, 'f', -1, 64)
	a := strings.Split(s, ".")
	for i := 0; i < decimals; i++ {
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
func Int2Float(amount *big.Int, decimals int) *big.Float {
	s := "1"
	for i := 0; i < decimals; i++ {
		s += "0"
	}
	b, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("Could not set big.Int string for value " + s)
	}
	f := new(big.Float).SetInt(amount)
	g := new(big.Float).SetInt(b)
	z := new(big.Float).Quo(f, g)
	return z
}
func Block(number int) (*types.Block, error) {
	client, err := ethclient.Dial(host)
	if err != nil {
		return nil, err
	}
	var block *types.Block
	if number < 0 {
		block, err = client.BlockByNumber(context.Background(), nil)
	} else {
		block, err = client.BlockByNumber(context.Background(), big.NewInt(int64(number)))
	}
	if err != nil {
		return nil, err
	}
	return block, nil
}
func Send(secret, address string, amount float64) (string, error) {
	address, err := Addresses(address)
	if err != nil {
		return "", err
	}
	valid := ValidateAddress(address)
	if !valid {
		return "", errors.New("Invalid address")
	}
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
	value := Float2Int(amount, 18)
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
	address, err := Addresses(address)
	if err != nil {
		return "", err
	}
	tokenContract, err = Addresses(tokenContract)
	if err != nil {
		return "", err
	}
	valid := ValidateAddress(address)
	if !valid {
		return "", errors.New("Invalid address")
	}
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
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		return "", err
	}
	a := Float2Int(amount, int(decimals))
	tx, err := instance.Transfer(auth, addre, a)
	if err != nil {
		return "", err
	}
	client.Close()
	return tx.Hash().Hex(), nil
}

func Balance(address string) (*big.Int, error) {
	address, err := Addresses(address)
	if err != nil {
		return nil, err
	}
	valid := ValidateAddress(address)
	if !valid {
		return nil, errors.New("Invalid address")
	}
	client, err := ethclient.Dial(host)
	if err != nil {
		return nil, err
	}
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}
	client.Close()
	return balance, nil
}
func BalanceERC20(tokenContract, address string) (*big.Float, error) {
	address, err := Addresses(address)
	if err != nil {
		return nil, err
	}
	tokenContract, err = Addresses(tokenContract)
	if err != nil {
		return nil, err
	}
	valid := ValidateAddress(address)
	if !valid {
		return nil, errors.New("Invalid address")
	}
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
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	balance := Int2Float(bal, int(decimals))
	client.Close()
	return balance, nil
}
