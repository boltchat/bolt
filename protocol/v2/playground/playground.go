package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/boltchat/lib/pgp"
	"github.com/boltchat/protocol/v2/encoder"
	"github.com/boltchat/protocol/v2/events"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
	"os"
	"strings"
)

// This is very poorly written, but it's just
// for debugging purposes. ;)
func printResult(res []byte) {
	fmt.Printf("bin: ")
	for _, b := range res {
		fmt.Printf("%08b ", b)
	}

	fmt.Println()

	fmt.Printf("dec: ")
	for _, b := range res {
		fmt.Printf("%d ", b)
	}
	fmt.Println()

	fmt.Printf("hex: %s", hex.EncodeToString(res))

	fmt.Println()
}

func sign(content string) *[]byte {
	r := strings.NewReader(content)
	buff := new(bytes.Buffer)

	entity, entityErr := pgp.LoadPGPEntity(
		"/home/kees/.config/boltchat/entities/default.pgp",
	)

	if entityErr != nil {
		panic(entityErr)
	}

	err := openpgp.DetachSignText(buff, entity, r, &packet.Config{})
	if err != nil {
		panic(err)
	}

	b := buff.Bytes()
	return &b
}

func main() {
	d := encoder.NewEncoder(nil)
	res := d.Encode(&events.Event{
		Header: &events.Header{
			Version:   1,
			EventType: events.JoinEvent,
			HasCRC:    false,
		},
		CRC32:     0xCBF43926,
		Signature: sign("Hello, world!"),
	})

	os.Stdout.Write(res)

	// printResult(res)
}
