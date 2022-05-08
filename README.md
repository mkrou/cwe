# Crypto Wallet Emulator
This tool is designed to help test the backend of crypto applications. 
It allows you to create private keys and sign messages.


## Supported Wallets
* [Solana](https://solana.com/)
  * [Phantom](https://phantom.app/)
* [Ethereum](https://ethereum.org/)
  * [Metamask](https://metamask.io/)

## Installation
```shell
go install github.com/mkrou/cwe/cmd/cwe
```

## Usage 
```shell
cwe --wallet=<wallet_name> generate
# output
# PK: xxx
# Address: xxx

cwe --wallet=<wallet_name> sign --pk=<private_key> --message=<message>
# output
# Signed Message: xxx
```

## Roadmap:
- [ ] http server for postman/insomnia automation
- [ ] docker image with http server
- [ ] transaction encoding/decoding