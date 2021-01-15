package main

import (
	"log"
	"os"

	"github.com/jhlq/go-flare/cmd"
	"github.com/jhlq/go-flare/gflr"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
func main() {
	dir, err := gflr.GetDir()
	if err != nil {
		log.Fatal("Could not create config directory: " + err.Error())
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			log.Fatal("Could not create config directory: " + err.Error())
		}
		file, err := os.Create(dir + "/addresses.csv")
		defer file.Close()
		if err != nil {
			log.Fatal("Could not create config file: " + err.Error())
		}
		_, err = file.WriteString(`@FXRP,0x2d0Eadb55df10A9f83AcD9C30EA59213aad13134
@YFIN,0x09842e3495B5cABd8dcE69Fd058Ba76D1b09B2d2
@YFLR,0x4d62C39BA6AC75ecBfeE1261F7599ccaD6bd1198
@xUSD,0xDCd9e0250c315e79Ad2164205a8AeF408B0b230f
@faucet,0x1C22a9dC4bfe93c49bBa31c36A887C01b5eba265`)
		if err != nil {
			log.Fatal("Could not create config file: " + err.Error())
		}
	}
	if !fileExists(dir + "/host.txt") {
		file, err := os.Create(dir + "/host.txt")
		defer file.Close()
		if err != nil {
			log.Fatal("Could not create host file: " + err.Error())
		}
		_, err = file.WriteString("https://costone.flare.network/ext/bc/C/rpc")
		if err != nil {
			log.Fatal("Could not create host file: " + err.Error())
		}
	}
	err = gflr.SetHost("")
	if err != nil {
		log.Fatal("Could not set host: " + err.Error())
	}
	cmd.Execute()
}
