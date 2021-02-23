package events

type Event struct {
	Header    *Header
	CRC32     uint32
	Signature *[]byte
	Payload   *[]byte
}
