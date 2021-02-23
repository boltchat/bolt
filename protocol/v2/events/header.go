package events

var ProtocolSignature [2]byte = [2]byte{0x13, 0x37}

type Header struct {
	Version   uint8
	EventType EventType
	HasCRC    bool
}
