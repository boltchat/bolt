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

	header = events.ProtocolSignature[:]
	header = append(header, byte(h.Version))
	header = append(header, evtType|crc)

	return header
}

func (e *Encoder) Encode(evt *events.Event) []byte {
	var res []byte

	var crc [4]byte

	binary.BigEndian.PutUint32(crc[:], evt.CRC32)

	res = e.encodeHeader(evt.Header)
	res = append(res, crc[:]...)

	return res
}

func NewEncoder(r io.Reader) *Encoder {
	return &Encoder{}
}
