package main

import (
	"fmt"
	"os"

	"github.com/nbd-wtf/go-nostr/nip19"
)

func main() {

	// get command line arg
	if len(os.Args) < 2 {
		fmt.Println("usage: ", os.Args[0], "npub2hex <npub pubkey>")
		os.Exit(1)
	}

	prefix, value, err := nip19.Decode(os.Args[1])

	fmt.Println("npub: ", os.Args[1])
	if err == nil {
		fmt.Println("found prefix", prefix)
		fmt.Println("hex pubkey:", value)
	} else {
		fmt.Println(err)
	}
}
