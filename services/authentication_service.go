package services

import (
	"bnet-mock/network"
	services "bnet-mock/network/account_manager"
	"bnet-mock/network/client/authentication_service"
	"bnet-mock/network/client/challenge_service"
	"bnet-mock/network/client/entity_types"
	"context"

	log "github.com/sirupsen/logrus"

	"google.golang.org/protobuf/proto"
)

type AuthenticationServiceHandler struct {
	accMgr services.AccountMgr
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

	// Hardcoded login data, should still come from dynamic source
	userDataResp, err := a.accMgr.GetAccountByEmail("test@test.com")
	if err != nil {
		return err
	}
	log.Debug("Verify Request: ", request.String())

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
