package wallets

import (
	"fmt"
	"github.com/gagliardetto/solana-go"
)

type Phantom struct {
}

func (p *Phantom) GeneratePK() (string, error) {
	account := solana.NewWallet()

	return account.PrivateKey.String(), nil
}

func (p *Phantom) GenerateAddress(pkString string) (string, error) {
	pk, err := solana.PrivateKeyFromBase58(pkString)
	if err != nil {
		return "", err
	}

	return pk.PublicKey().String(), nil
}

func (p *Phantom) Sign(pkString string, msg string) (string, error) {
	pk, err := solana.PrivateKeyFromBase58(pkString)
	if err != nil {
		return "", err
	}

	sign, err := pk.Sign([]byte(msg))
	if err != nil {
		return "", fmt.Errorf("can't sign message: %w", err)
	}

	return sign.String(), nil
}
