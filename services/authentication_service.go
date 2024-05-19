package services

import (
	"bnet-mock/network"
	"bnet-mock/network/client/authentication_service"
	"bnet-mock/network/client/challenge_service"
	"bnet-mock/network/client/entity_types"
	"context"

	log "github.com/sirupsen/logrus"

	"google.golang.org/protobuf/proto"
)

type AuthenticationServiceHandler struct {
}

func (a AuthenticationServiceHandler) Logon(ctx context.Context, request *authentication_service.LogonRequest, i interface{}) error {
	client := i.(*network.RpcClient)

	log.Debug("Logon Request: ", request.String())

	logonResp := new(challenge_service.ChallengeExternalRequest)
	logonResp.PayloadType = proto.String("web_auth_url")
	logonResp.Payload = []byte("http://bnet-emu.fish:6969/battlenet/login?externalChallenge=login&app=pro")

	return client.SendResponse(3151632159, 0, 3, 0, 1, logonResp)
}

func (a AuthenticationServiceHandler) ModuleNotify(ctx context.Context, notification *authentication_service.ModuleNotification, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthenticationServiceHandler) ModuleMessage(ctx context.Context, request *authentication_service.ModuleMessageRequest, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthenticationServiceHandler) SelectGameAccount_DEPRECATED(ctx context.Context, id *entity_types.EntityId, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthenticationServiceHandler) GenerateSSOToken(ctx context.Context, request *authentication_service.GenerateSSOTokenRequest, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthenticationServiceHandler) SelectGameAccount(ctx context.Context, request *authentication_service.SelectGameAccountRequest, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthenticationServiceHandler) VerifyWebCredentials(ctx context.Context, request *authentication_service.VerifyWebCredentialsRequest, i interface{}) error {
	client := i.(*network.RpcClient)

	high1 := uint64(72057594037927936)
	low1 := uint64(379775058)
	high2 := uint64(144115192376095343)
	low2 := uint64(559865145)

	log.Debug("Verify Request: ", request.String())

	userDataResp := new(authentication_service.LogonResult)
	userDataResp.ErrorCode = proto.Uint32(0)
	userDataResp.AccountId = &entity_types.EntityId{
		High: &high1,
		Low:  &low1,
	}
	userDataResp.GameAccountId = []*entity_types.EntityId{
		&entity_types.EntityId{
			High: &high2,
			Low:  &low2,
		},
	}
	userDataResp.Email = proto.String("")
	userDataResp.AvailableRegion = []uint32{1}
	userDataResp.ConnectedRegion = proto.Uint32(1)
	userDataResp.BattleTag = proto.String("Test#0001")
	userDataResp.GeoipCountry = proto.String("US")
	userDataResp.SessionKey = []byte{0xAB, 0xA8, 0xF2, 0x94, 0x99, 0x77, 0x65, 0x55, 0x58, 0x6B, 0xEB, 0xC1, 0xEE, 0xF0, 0xB4, 0xCE, 0x24, 0xCC, 0xAF, 0xDD, 0x53, 0x37, 0xF7, 0x78, 0x62, 0xC1, 0xA0, 0x54, 0xF9, 0xA5, 0x36, 0x15, 0x0D, 0xE6, 0x78, 0xFD, 0x16, 0x49, 0x39, 0xDC, 0xE7, 0xB4, 0x48, 0x5B, 0xA2, 0xC4, 0x73, 0x05, 0x44, 0xB2, 0x00, 0xAD, 0xA8, 0x90, 0x6E, 0x74, 0x15, 0xCC, 0xF2, 0x65, 0x73, 0x21, 0x23, 0x99}
	userDataResp.RestrictedMode = proto.Bool(false)

	return client.SendResponse(
		authentication_service.AuthenticationListenerOriginalFullyQualifiedDescriptorNameHash,
		0,
		authentication_service.AuthenticationListenerMethod_OnLogonComplete.ToUint32(),
		0,
		3,
		userDataResp,
	)
}

func (a AuthenticationServiceHandler) GenerateWebCredentials(ctx context.Context, request *authentication_service.GenerateWebCredentialsRequest, i interface{}) error {
	client := i.(*network.RpcClient)

	log.Debug("Logon Request: ", request.String())

	logonResp := new(challenge_service.ChallengeExternalRequest)
	logonResp.PayloadType = proto.String("web_auth_url")
	logonResp.Payload = []byte("http://bnet-emu.fish:6969/battlenet/login?externalChallenge=login&app=pro")

	return client.SendResponse(
		challenge_service.ChallengeListenerOriginalFullyQualifiedDescriptorNameHash,
		0,
		challenge_service.ChallengeListenerMethod_OnExternalChallenge.ToUint32(),
		0,
		1,
		logonResp,
	)
}
