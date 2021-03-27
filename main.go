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
		_, err = file.WriteString(`@faucet,0x1C22a9dC4bfe93c49bBa31c36A887C01b5eba265`)
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
		_, err = file.WriteString("https://coston.flare.network/ext/bc/C/rpc")
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
