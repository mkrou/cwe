package wallets

import (
	"fmt"
)

type Wallet interface {
	GeneratePK() (string, error)
	GenerateAddress(pkString string) (string, error)
	Sign(pkString string, msg string) (string, error)
}

const (
	PhantomLabel  = "phantom"
	MetamaskLabel = "metamask"
)

func NewWallet(wallet string) (Wallet, error) {
	switch wallet {
	case PhantomLabel:
		return &Phantom{}, nil
	case MetamaskLabel:
		return &Metamask{}, nil
	default:
		return nil, fmt.Errorf("wallet %s in unknown", wallet)
	}
}
