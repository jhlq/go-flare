# go-flare
Tools for the Flare network

If you are on Windows/Mac check out the release.

If you are on Linux running should be as easy as
```
./go-flare
./go-flare send --help
```

Else you need gcc installed in order to go build or go run main.go

Edit the file addresses.csv in go-flare-config at your home dir to add addresses to your addressbook, note that names have to start with @. Then it is easy to query for example the balance of [the faucet](https://weilianqi.laerande.org/flarefaucet.php):
```
./go-flare balance @faucet
```

The addressbook is prefilled with the four primary tokens of Flare Finance, so to send YFLR simply:
```
./go-flare sendERC20 @YFLR 0x213E269a503AD47Db5fa115905CbE3bE1aF490E3 10
```

Edit the file hosts.txt to change which node you connect to.

Available commands and their usage can be garnered from running go-flare and utilizing the help flag.

FLR-20 tokens operate with the same interface as ERC-20 ones.

If you want to use the commands directly in Go code import github.com/jhlq/go-flare/gflr and then look at the various files in the cmd directory, it should be pretty self explanatory. Note that you have to set the host with gflr.SetHost("") (the "" tells gflr to use the host specified in the config folder)

This code has not been audited, use at your own risk.
