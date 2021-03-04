package decoder

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/boltchat/protocol/v2/events"
)

type Decoder struct{}

var ErrHeaderTooShort = errors.New("header too short")
var ErrInvalidSignature = errors.New("invalid protocol signature")

func (d *Decoder) DecodeHeader(h []byte) (*events.Header, error) {
	header := &events.Header{}

	if len(h) < 4 {
		return nil, ErrHeaderTooShort
	}

	if !bytes.HasPrefix(h, events.ProtocolSignature[:]) {
		return nil, ErrInvalidSignature
	}

	evtType := events.EventType(h[3] >> 2)
	crc := (h[3] & 2) == 2
	compression := (h[3] & 1) == 1

	header.Version = uint8(h[2])
	header.EventType = evtType
	header.HasCRC = crc
	header.HasCompression = compression

	return header, nil
}

func (d *Decoder) Decode(b []byte) (*events.Event, error) {
	event := &events.Event{}

	header, err := d.DecodeHeader(b)
	if err != nil {
		return nil, err
	}

	event.Header = header
	event.CRC32 = binary.BigEndian.Uint32(b[4:8])

	sigLen := binary.BigEndian.Uint16(b[8:10])

	sig := b[10 : sigLen+10]
	payload := b[10+sigLen:]

	event.Signature = &sig
	event.Payload = &payload

	return event, nil
}

func NewDecoder() *Decoder {
	return &Decoder{}
}
