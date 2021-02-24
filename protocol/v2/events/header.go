package events

var ProtocolSignature [2]byte = [2]byte{0x10, 0xE6}

type Header struct {
	Version        uint8
	EventType      EventType
	HasCRC         bool
	HasCompression bool
}
