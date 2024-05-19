package services

import (
	"bnet-mock/network"
	"bnet-mock/network/client/attribute_types"
	"bnet-mock/network/client/game_utilities_service"
	"context"
	"encoding/hex"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type GameUtilitiesServiceHandler struct {
}

func (g GameUtilitiesServiceHandler) ProcessClientRequest(ctx context.Context, request *game_utilities_service.ClientRequest, i interface{}) error {
	client := i.(*network.RpcClient)
	log.Debugf("GameUtilitiesServiceHandler.ProcessClientRequest: %v", request.String())

	s1 := "95423E62008C29B0D53E1759C5A863DF9F8E814BBA0B99D425E49B00C08FC8A8BBEF7FA424031A3B2ACBE51628BE25D115BC3FD077B128E85ACBDC690C65CB89"
	decoded1, err := hex.DecodeString(s1)
	if err != nil {
		log.Fatal(err)
	}
	s2 := "474F1B1A5C4ACC23F2C8319E424791F66DA541B85B2B3C2D0C0E6E5E3A7C9FD233FD748BA4903B2DE4D072D34D040B7BFA00CEC65D02524A4E274C7956F2C60C"
	decoded2, err := hex.DecodeString(s2)
	if err != nil {
		log.Fatal(err)
	}
	s3 := "78296ADEB7A7E78DD901FA93646032426A33DB73BEBBC1DBD2FC31DCDF1CE2D15A69F291B0F3AF31D7BB597356900BBE9EF3FE21FE0BF64E5D7F4218A4502173"
	decoded3, err := hex.DecodeString(s3)
	if err != nil {
		log.Fatal(err)
	}
	s4 := "3852D9D47C598BB5C140A1F0188FF42C1C08AB6915FDF916E456A1995B5F07E00EDE2DD4B79E83F222C8147F207B101987717AC1E65C3B9CB119B3D35B4E8BAF"
	decoded4, err := hex.DecodeString(s4)
	if err != nil {
		log.Fatal(err)
	}

	host := &attribute_types.Attribute{
		Name: proto.String("hostv4"),
		Value: &attribute_types.Variant{
			StringValue: proto.String("us.game.bwattle.uwu:1185"),
		},
	}

	cid := &attribute_types.Attribute{
		Name: proto.String("cid"),
		Value: &attribute_types.Variant{
			UintValue: proto.Uint64(379775058),
		},
	}

	rtype := &attribute_types.Attribute{
		Name: proto.String("response_type"),
		Value: &attribute_types.Variant{
			StringValue: proto.String("ReferralInfo"),
		},
	}

	k0 := &attribute_types.Attribute{
		Name: proto.String("k0"),
		Value: &attribute_types.Variant{
			BlobValue: decoded1,
		},
	}

	k1 := &attribute_types.Attribute{
		Name: proto.String("k1"),
		Value: &attribute_types.Variant{
			BlobValue: decoded2,
		},
	}

	k2 := &attribute_types.Attribute{
		Name: proto.String("k2"),
		Value: &attribute_types.Variant{
			BlobValue: decoded3,
		},
	}

	k3 := &attribute_types.Attribute{
		Name: proto.String("k3"),
		Value: &attribute_types.Variant{
			BlobValue: decoded4,
		},
	}

	res := new(game_utilities_service.ClientResponse)
	res.Attribute = make([]*attribute_types.Attribute, 0)
	res.Attribute = append(res.Attribute, rtype)
	res.Attribute = append(res.Attribute, k0)
	res.Attribute = append(res.Attribute, k1)
	res.Attribute = append(res.Attribute, k2)
	res.Attribute = append(res.Attribute, k3)
	res.Attribute = append(res.Attribute, cid)
	res.Attribute = append(res.Attribute, host)

	return client.SendResponse(
		0,
		254,
		0,
		0,
		3,
		res,
	)
}

func (g GameUtilitiesServiceHandler) PresenceChannelCreated(ctx context.Context, request *game_utilities_service.PresenceChannelCreatedRequest, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (g GameUtilitiesServiceHandler) GetPlayerVariables(ctx context.Context, request *game_utilities_service.GetPlayerVariablesRequest, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (g GameUtilitiesServiceHandler) ProcessServerRequest(ctx context.Context, request *game_utilities_service.ServerRequest, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (g GameUtilitiesServiceHandler) OnGameAccountOnline(ctx context.Context, notification *game_utilities_service.GameAccountOnlineNotification, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (g GameUtilitiesServiceHandler) OnGameAccountOffline(ctx context.Context, notification *game_utilities_service.GameAccountOfflineNotification, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (g GameUtilitiesServiceHandler) GetAchievementsFile(ctx context.Context, request *game_utilities_service.GetAchievementsFileRequest, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (g GameUtilitiesServiceHandler) GetAllValuesForAttribute(ctx context.Context, request *game_utilities_service.GetAllValuesForAttributeRequest, i interface{}) error {
	//TODO implement me
	panic("implement me")
}
