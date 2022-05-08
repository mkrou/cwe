package main

import (
	"fmt"
	"github.com/mkrou/cwe/pkg/wallets"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const (
	FlagWallet     = "wallet"
	FlagPrivateKey = "pk"
	FlagMessage    = "message"
)

func main() {
	app := &cli.App{
		Name: "Crypto Wallet Emulator",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  FlagWallet,
				Usage: "set wallet name [phantom, metamask]",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "generate",
				Usage: "generate pk & address",
				Action: func(ctx *cli.Context) error {
					wallet, err := wallets.NewWallet(ctx.String(FlagWallet))
					if err != nil {
						return err
					}

					pk, err := wallet.GeneratePK()
					if err != nil {
						return err
					}

					address, err := wallet.GenerateAddress(pk)
					if err != nil {
						return err
					}

					fmt.Printf("PK: %s \n", pk)
					fmt.Printf("Address: %s \n", address)
					return nil
				},
			},
			{
				Name:  "sign",
				Usage: "sign message",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: FlagMessage,
					},
					&cli.StringFlag{
						Name: FlagPrivateKey,
					},
				},
				Action: func(ctx *cli.Context) error {
					wallet, err := wallets.NewWallet(ctx.String(FlagWallet))
					if err != nil {
						return err
					}

					pk := ctx.String(FlagPrivateKey)
					msg := ctx.String(FlagMessage)

					signedMsg, err := wallet.Sign(pk, msg)
					if err != nil {
						return err
					}

					fmt.Printf("Signed message: %s \n", signedMsg)

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
