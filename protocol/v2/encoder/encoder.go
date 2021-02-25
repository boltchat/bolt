package encoder

import (
	"encoding/binary"
	"github.com/boltchat/protocol/v2/events"
	"io"
)

type Encoder struct {
}

func (e *Encoder) encodeHeader(h *events.Header) []byte {
	var header []byte

	evtType := byte(h.EventType) << 2

	crc := byte(0x00)
	if h.HasCRC {
		crc = byte(0x01) << 1
	}

	compression := byte(0x00)
	if h.HasCompression {
		compression = byte(0x01)
	}

	header = events.ProtocolSignature[:]
	header = append(header, byte(h.Version))
	header = append(header, evtType|crc|compression)

	return header
}

func (e *Encoder) Encode(evt *events.Event) []byte {
	// The final result
	var res []byte

	// The CRC-32 checksum
	var crc [4]byte

	// The PGP signature length
	var sigLen [2]byte

	// The event payload length
	var payloadLen [2]byte

	binary.BigEndian.PutUint32(crc[:], evt.CRC32)
	binary.BigEndian.PutUint16(sigLen[:], uint16(len(*evt.Signature)))
	binary.BigEndian.PutUint16(payloadLen[:], uint16(len(*evt.Payload)))

	res = e.encodeHeader(evt.Header)
	res = append(res, crc[:]...)

	res = append(res, sigLen[:]...)
	res = append(res, *evt.Signature...)

	res = append(res, payloadLen[:]...)
	res = append(res, *evt.Payload...)

	return res
}

func NewEncoder(r io.Reader) *Encoder {
	return &Encoder{}
}
