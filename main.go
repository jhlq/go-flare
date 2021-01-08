package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/jhlq/go-flare/cmd"
)

func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Println("Could not create config directory: " + err.Error())
	} else {
		homeDirectory := user.HomeDir
		if _, err := os.Stat(homeDirectory + "/go-flare-config"); os.IsNotExist(err) {
			err = os.Mkdir(homeDirectory+"/go-flare-config", 0755)
			if err != nil {
				fmt.Println("Could not create config directory: " + err.Error())
			} else {
				file, err := os.Create(homeDirectory + "/go-flare-config/addresses.csv")
				defer file.Close()
				if err != nil {
					fmt.Println("Could not create config file: " + err.Error())
				} else {
					_, err = file.WriteString(`@FXRP,0x2d0Eadb55df10A9f83AcD9C30EA59213aad13134
@YFIN,0x09842e3495B5cABd8dcE69Fd058Ba76D1b09B2d2
@YFLR,0x4d62C39BA6AC75ecBfeE1261F7599ccaD6bd1198
@xUSD,0xDCd9e0250c315e79Ad2164205a8AeF408B0b230f
@faucet,0x1C22a9dC4bfe93c49bBa31c36A887C01b5eba265`)
					if err != nil {
						fmt.Println("Could not create config file: " + err.Error())
					}
				}
			}
		}
	}
	cmd.Execute()
}
