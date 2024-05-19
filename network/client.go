package network

import (
	"bnet-mock/network/client/rpc_types"
	"encoding/binary"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
)

type RpcClient struct {
	net.Conn
	ClientId uint32
	Server   *RpcServer
}

func NewRpcClient(con net.Conn, clientId uint32, server *RpcServer) *RpcClient {
	return &RpcClient{
		Conn:     con,
		ClientId: clientId,
		Server:   server,
	}
}

func (c *RpcClient) Send(packet *BNetPacket) error {
	log.Print("Sent packet: ", packet.Header.String())

	return c.Server.SendPacket(c, packet)
}

func (c *RpcClient) SendResponse(serviceHash uint32, serviceId uint32, methodId uint32, status uint32, token uint32, body interface{}) error {
	header := new(rpc_types.Header)
	header.ServiceHash = proto.Uint32(serviceHash)
	header.ServiceId = proto.Uint32(serviceId)
	header.MethodId = proto.Uint32(methodId)
	header.Status = proto.Uint32(status)
	header.Token = proto.Uint32(token)

	packet := new(BNetPacket)
	packet.Header = header

	if body != nil {
		if err := packet.MarshalBody(body.(proto.Message)); err != nil {
			return err
		}
	} else {
		packet.Body = make([]byte, 0)
	}

	if err := c.Send(packet); err != nil {
		return err
	}

	return nil
}

func (c *RpcClient) SendRaw(packet []byte) error {
	headerSize := binary.BigEndian.Uint16(packet[0:2])

	var header rpc_types.Header
	if err := proto.Unmarshal(packet[2:2+headerSize], &header); err != nil {
		return err
	}

	log.Println("Sent Raw packet: ", header.String())

	_, err := c.Write(packet)
	return err
}
