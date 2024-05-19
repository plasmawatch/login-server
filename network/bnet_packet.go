package network

import (
	"bnet-mock/network/client/rpc_types"
	"encoding/binary"
	"errors"
	"google.golang.org/protobuf/proto"
	"log"
)

func HeaderFromPacket(packet []byte) (*rpc_types.Header, uint16, error) {
	if len(packet) < 2 {
		return nil, 0, errors.New("packet too small")
	}

	headerSize := binary.BigEndian.Uint16(packet[0:2])

	var header rpc_types.Header
	if err := proto.Unmarshal(packet[2:2+headerSize], &header); err != nil {
		return nil, 0, err
	}

	return &header, headerSize, nil
}

type BNetPacket struct {
	Header *rpc_types.Header
	Body   []byte
}

func NewBNetPacket(buf []byte) (*BNetPacket, error) {
	header, headerSize, err := HeaderFromPacket(buf)
	if err != nil {
		return nil, err
	}

	// read the body after the header
	body := buf[2+headerSize:]

	return &BNetPacket{
		Header: header,
		Body:   body,
	}, nil
}

func (b *BNetPacket) GetHeader() *rpc_types.Header {
	return b.Header
}

func (b *BNetPacket) GetBody() []byte {
	return b.Body
}

func (b *BNetPacket) GetBodySize() uint32 {
	return b.Header.GetSize()
}

// Marshal the packet into a byte slice
func (b *BNetPacket) Marshal() ([]byte, error) {
	log.Print("Marshalling packet: ", b.Header.String())

	header, err := proto.Marshal(b.Header)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 2+len(header)+len(b.Body))
	binary.BigEndian.PutUint16(buf[0:2], uint16(len(header)))
	copy(buf[2:2+len(header)], header)
	copy(buf[2+len(header):], b.Body)

	return buf, nil
}

// UnmarshalBody the body into the given message
func (b *BNetPacket) UnmarshalBody(msg proto.Message) error {
	return proto.Unmarshal(b.Body, msg)
}

// MarshalBody the given message into the body
func (b *BNetPacket) MarshalBody(msg proto.Message) error {
	body, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	b.Body = body
	if b.Header != nil {
		b.Header.Size = proto.Uint32(uint32(len(b.Body)))
	}
	
	return nil
}
