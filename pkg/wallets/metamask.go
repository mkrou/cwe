package wallets

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Metamask struct {
}

func (m *Metamask) GeneratePK() (string, error) {
	pk, err := crypto.GenerateKey()
	if err != nil {
		return "", fmt.Errorf("can't generate private key for ethereum: %w", err)
	}

	return common.Bytes2Hex(crypto.FromECDSA(pk)), nil
}

func (m *Metamask) GenerateAddress(pkString string) (string, error) {
	pk, err := crypto.HexToECDSA(pkString)
	if err != nil {
		return "", fmt.Errorf("can't cast ethereum private key to ecdsa: %w", err)
	}

	return crypto.PubkeyToAddress(pk.PublicKey).String(), nil
}

func (m *Metamask) Sign(pkString string, msg string) (string, error) {
	pk, err := crypto.HexToECDSA(pkString)
	if err != nil {
		return "", fmt.Errorf("can't cast ethereum private key to ecdsa: %w", err)
	}

	msgWithPrefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)
	msgHash := crypto.Keccak256([]byte(msgWithPrefix))
	sign, err := crypto.Sign(msgHash, pk)
	if err != nil {
		return "", fmt.Errorf("can't sign message: %w", err)
	}

	return common.Bytes2Hex(sign), nil
}
