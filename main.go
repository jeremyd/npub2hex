package main

import (
	"fmt"
	"os"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

func main() {

	// get command line arg
	if len(os.Args) < 2 {
		fmt.Println("usage: ", os.Args[0], "npub2hex [--generate] <npub pubkey>")
		os.Exit(1)
	}

	if os.Args[1] == "--generate" {
		privk := nostr.GeneratePrivateKey()
		pubk, _ := nostr.GetPublicKey(privk)
		npub, _ := nip19.EncodePublicKey(pubk)
		nsec, _ := nip19.EncodePrivateKey(privk)
		fmt.Printf("NOSTR_PRIVATE_KEY=%s\nNOSTR_PUBLIC_KEY=%s\nNOSTR_NPUB=%s\nNOSTR_NSEC=%s", privk, pubk, npub, nsec)
		os.Exit(0)
	}

	if os.Args[1][0:4] != "npub" {
		np, e := nip19.EncodePublicKey(os.Args[1])
		if e != nil {
			fmt.Println("wtf, got error", e)
		}
		fmt.Println("npub: ", np)
	} else {
		prefix, value, err := nip19.Decode(os.Args[1])
		fmt.Println("npub: ", os.Args[1])
		if err == nil {
			fmt.Println("found prefix", prefix)
			fmt.Println("hex pubkey:", value)
		} else {
			fmt.Println(err)
		}
	}

}
