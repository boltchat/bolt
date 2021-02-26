package main

import (
	"bytes"
	"fmt"
	"github.com/boltchat/lib/pgp"
	"github.com/boltchat/protocol/v2/decoder"
	"github.com/boltchat/protocol/v2/encoder"
	"github.com/boltchat/protocol/v2/events"
	"github.com/vmihailenco/msgpack/v5"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
	"os"
	"strings"
)

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

func payload(d string) *[]byte {
	payload := struct{ msg string }{
		msg: d,
	}

	b, err := msgpack.Marshal(&payload)
	if err != nil {
		panic(err)
	}

	return &b
}

func main() {
	enc := encoder.NewEncoder()
	header := &events.Header{
		Version:        1,
		EventType:      events.JoinEvent,
		HasCRC:         true,
		HasCompression: false,
	}

	encResult := enc.Encode(&events.Event{
		Header:    header,
		CRC32:     0xCBF43926,
		Signature: sign("Hello, world!"),
		Payload:   payload("Hi there! This is an event."),
	})

	dec := decoder.NewDecoder()
	decResult, err := dec.Decode(encResult)
	if err != nil {
		panic(err)
	}

	os.Stderr.WriteString(
		fmt.Sprintf("%v\n%v\n", *decResult.Header, *decResult),
	)

	os.Stdout.Write(encResult)
}
