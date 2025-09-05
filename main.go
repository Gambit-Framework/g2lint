package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Gambit-Framework/g2lint/sprofile"
	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("g2lint", "parse and validate server and evasion profile")

	sprofilePath := parser.String("s", "server-profile", &argparse.Options{Required: false, Help: "path to server profile",
		Validate: func(args []string) error {
			if _, err := os.Stat(args[0]); errors.Is(err, os.ErrNotExist) {
				return errors.New("invalid server profile path")
			}

			return nil
		},
	})

	eprofilePath := parser.String("e", "evasion-profile", &argparse.Options{Required: false, Help: "path to evasion profile",
		Validate: func(args []string) error {
			if _, err := os.Stat(args[0]); errors.Is(err, os.ErrNotExist) {
				return errors.New("invalid server profile path")
			}

			return nil
		},
	})

	verbose := parser.Flag("v", "verbose", &argparse.Options{Required: false, Default: false, Help: "show parsed data"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	fmt.Print(`
 ██████╗ ██████╗ ██╗     ██╗███╗   ██╗████████╗
██╔════╝ ╚════██╗██║     ██║████╗  ██║╚══██╔══╝
██║  ███╗ █████╔╝██║     ██║██╔██╗ ██║   ██║   
██║   ██║██╔═══╝ ██║     ██║██║╚██╗██║   ██║   
╚██████╔╝███████╗███████╗██║██║ ╚████║   ██║   
 ╚═════╝ ╚══════╝╚══════╝╚═╝╚═╝  ╚═══╝   ╚═╝   

`)

	if *sprofilePath == "" && *eprofilePath == "" {
		fmt.Print(parser.Usage(err))
		os.Exit(0)
	}

	if *sprofilePath != "" {
		sp, err := sprofile.ParseServerProfile(*sprofilePath)
		if err != nil {
			log.Fatal(err)
		}

		err = sp.Validate(*verbose)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("[+] Server Profile Parsed Successfully")
	}
}
