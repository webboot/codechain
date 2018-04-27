package command

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"

	"github.com/frankbraun/codechain/hashchain"
)

// RemKey implements the 'remkey' command.
func RemKey(argv0 string, args ...string) error {
	fs := flag.NewFlagSet(argv0, flag.ContinueOnError)
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s pubkey\n", argv0)
		fmt.Fprintf(os.Stderr, "Remove existing signer from hashchain.\n")
		fs.PrintDefaults()
	}
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() != 1 {
		fs.Usage()
		return flag.ErrHelp
	}
	pubkey := fs.Arg(0)
	_, err := base64.URLEncoding.DecodeString(pubkey)
	if err != nil {
		return fmt.Errorf("cannot decode pubkey: %s", err)
	}
	c, err := hashchain.Read(hashchainFile)
	if err != nil {
		return err
	}
	line, err := c.RemKey(hashchainFile, pubkey)
	if err != nil {
		return err
	}
	fmt.Println(line)
	return nil
}