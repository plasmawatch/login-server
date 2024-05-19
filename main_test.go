package main

import (
	"bnet-mock/packets/proto/bgs/low/pb/client/authentication_service"
	"bnet-mock/packets/proto/bgs/low/pb/client/challenge_service"
	"bnet-mock/packets/proto/bgs/low/pb/client/rpc_types"
	"encoding/binary"
	"encoding/hex"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	hexx := "000a08fe01180028a80130000a0c08aac0a2c50f10aea09ca506120a0895df0310edf8aca50630a5f0a3d6dd828003425e0a5c0d53550000157274656d1a203488f45185ae7e8501db89583de0cc291589347588777fd3bcb299e24df8c5e1222e68747470733a2f2f70726f642e6465706f742e626174746c652e6e65742f247b686173687d2e247b75736167657d4a21463841384130324136344137313032452d303030303030303030303030454639355001"

	payload, err := hex.DecodeString(hexx)
	if err != nil {
		t.Error(err)
	}

	t.Log(hex.Dump(payload))
	log.Printf("Total Bytes: %d", len(payload))

	headerLen := binary.BigEndian.Uint16(payload[0:2])
	log.Printf("headerLen: %d", headerLen)

	headerBytes := payload[2 : 2+headerLen]

	var serviceHeader rpc_types.Header
	if err := proto.Unmarshal(headerBytes, &serviceHeader); err != nil {
		t.Error(err)
	}

	payloadSize := serviceHeader.GetSize()
	payloadBytes := payload[2+headerLen : uint32(headerLen)+payloadSize]

	var moduleLoadRequest authentication_service.ModuleLoadRequest
	if err := proto.Unmarshal(payloadBytes, &moduleLoadRequest); err != nil {
		t.Error(err)
	}

	t.Log(serviceHeader.String())
	t.Log(moduleLoadRequest.String())
	t.Log(moduleLoadRequest.GetModuleHandle().String())
}

func TestDecodeHeader(t *testing.T) {
	hexx := "000d080010071802282f5d01fcec0d0a2d55532d31383565346439333364323633343338626532393031656434343037366432612d333739373735303538"

	payload, err := hex.DecodeString(hexx)
	if err != nil {
		t.Error(err)
	}

	t.Log("payload size: ", len(payload))

	t.Log(hex.Dump(payload))
	log.Printf("Total Bytes: %d", len(payload))

	headerLen := binary.BigEndian.Uint16(payload[0:2])
	log.Printf("headerLen: %d", headerLen)

	headerBytes := payload[2 : 2+headerLen]

	var serviceHeader rpc_types.Header
	if err := proto.Unmarshal(headerBytes, &serviceHeader); err != nil {
		t.Error(err)
	}

	t.Log(serviceHeader.String())

	os.WriteFile("logon_res_clean.bin", payload, 0644)
}

func TestCreateLogonPayload(t *testing.T) {
	var logonResponse challenge_service.ChallengeExternalRequest
	logonResponse.PayloadType = proto.String("web_auth_url")
	logonResponse.Payload = []byte("https://us.battle.net/login/?externalChallenge=login&app=pro")
	// logonResponse.RequestToken = proto.String("F8A8A02A64A7102E-000000000000EF95")

	logonResponseBytes, err := proto.Marshal(&logonResponse)
	if err != nil {
		t.Error(err)
	}

	logonResponseHeader := new(rpc_types.Header)
	logonResponseHeader.ServiceId = proto.Uint32(0)
	logonResponseHeader.MethodId = proto.Uint32(3)
	logonResponseHeader.Token = proto.Uint32(1)
	logonResponseHeader.Size = proto.Uint32(uint32(len(logonResponseBytes)))
	logonResponseHeader.ServiceHash = proto.Uint32(3151632159)
	logonResponseHeader.ClientId = proto.String("F8A8A02A64A7102E-000000000000EF95")

	logonResponseHeaderBytes, err := proto.Marshal(logonResponseHeader)
	if err != nil {
		t.Error(err)
	}

	logonResponseHeaderLen := make([]byte, 2)
	binary.BigEndian.PutUint16(logonResponseHeaderLen, uint16(len(logonResponseHeaderBytes)))

	logonResponsePayload := append(logonResponseHeaderLen, logonResponseHeaderBytes...)
	logonResponsePayload = append(logonResponsePayload, logonResponseBytes...)

	t.Log(logonResponseHeader.String())
	t.Log(hex.Dump(logonResponsePayload))

	os.WriteFile("logon_response.bin", logonResponsePayload, 0644)
}
