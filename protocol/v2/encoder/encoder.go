package encoder

import (
	"encoding/binary"

	"github.com/boltchat/protocol/v2/events"
	"github.com/vmihailenco/msgpack/v5"
)

type Encoder struct{}

func (e *Encoder) EncodeHeader(h *events.Header) []byte {
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

func (e *Encoder) EncodePayload(p interface{}) ([]byte, error) {
	return msgpack.Marshal(&p)
}

func (e *Encoder) Encode(evt *events.Event) ([]byte, error) {
	var res []byte
	var crc [4]byte
	var sigLen [2]byte
	var payloadLen [2]byte

	if evt.RawPayload == nil {
		b, err := e.EncodePayload(evt.Payload)
		if err != nil {
			return nil, err
		}

		evt.RawPayload = &b
	}

	binary.BigEndian.PutUint32(crc[:], evt.CRC32)
	binary.BigEndian.PutUint16(sigLen[:], uint16(len(*evt.Signature)))
	binary.BigEndian.PutUint16(payloadLen[:], uint16(len(*evt.RawPayload)))

	res = e.EncodeHeader(evt.Header)
	res = append(res, crc[:]...)

	res = append(res, sigLen[:]...)
	res = append(res, *evt.Signature...)

	res = append(res, payloadLen[:]...)
	res = append(res, *evt.RawPayload...)

	return res
}

func NewEncoder() *Encoder {
	return &Encoder{}
}
