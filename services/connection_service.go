package services

import (
	"bnet-mock/network"
	"bnet-mock/network/client/connection_service"
	"bnet-mock/network/client/rpc_types"
	"context"
	"encoding/hex"
	"google.golang.org/protobuf/proto"
	"time"
)

type ConnectionServiceHandler struct{}

func (c ConnectionServiceHandler) Connect(ctx context.Context, request *connection_service.ConnectRequest, i interface{}) error {
	client := i.(*network.RpcClient)

	res := new(connection_service.ConnectResponse)
	res.ServerId = &rpc_types.ProcessId{
		Label: proto.Uint32(0),
		Epoch: proto.Uint32(uint32(time.Now().Unix())),
	}
	res.ServerTime = proto.Uint64(uint64(time.Now().Unix()))
	res.ClientId = &rpc_types.ProcessId{
		Label: proto.Uint32(1),
		Epoch: proto.Uint32(client.ClientId),
	}

	if err := client.SendResponse(0, 254, 0, 0, 3, res); err != nil {
		return err
	}

	hexx := "000a08fe01180028a80130000a0c08aac0a2c50f10aea09ca506120a0895df0310edf8aca50630a5f0a3d6dd828003425e0a5c0d53550000157274656d1a203488f45185ae7e8501db89583de0cc291589347588777fd3bcb299e24df8c5e1222e68747470733a2f2f70726f642e6465706f742e626174746c652e6e65742f247b686173687d2e247b75736167657d4a21463841384130324136344137313032452d303030303030303030303030454639355001"
	hardCodedPacket, _ := hex.DecodeString(hexx)

	return client.SendRaw(hardCodedPacket)
}

func (c ConnectionServiceHandler) Bind(ctx context.Context, request *connection_service.BindRequest, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c ConnectionServiceHandler) Echo(ctx context.Context, request *connection_service.EchoRequest, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c ConnectionServiceHandler) ForceDisconnect(ctx context.Context, notification *connection_service.DisconnectNotification, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c ConnectionServiceHandler) KeepAlive(ctx context.Context, data *rpc_types.NoData, i interface{}) error {
	client := i.(*network.RpcClient)
	res := new(connection_service.ConnectResponse)
	res.ServerId = &rpc_types.ProcessId{
		Label: proto.Uint32(0),
		Epoch: proto.Uint32(uint32(time.Now().Unix())),
	}
	res.ServerTime = proto.Uint64(uint64(time.Now().Unix()))
	res.ClientId = &rpc_types.ProcessId{
		Label: proto.Uint32(1),
		Epoch: proto.Uint32(client.ClientId),
	}

	return client.SendResponse(0, 254, 0, 0, 3, res)
}

func (c ConnectionServiceHandler) Encrypt(ctx context.Context, request *connection_service.EncryptRequest, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c ConnectionServiceHandler) RequestDisconnect(ctx context.Context, request *connection_service.DisconnectRequest, i interface{}) error {
	//TODO implement me
	panic("implement me")
}
