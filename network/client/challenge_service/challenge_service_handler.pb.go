// Code generated by protoc-gen-go-bnet. DO NOT EDIT.

package challenge_service

import (
	"context"
	"fmt"
	"strings"
	"google.golang.org/protobuf/proto"
)

// ChallengeServiceHandler is the handler interface for bgs.protocol.challenge.v1.ChallengeService
type ChallengeServiceHandler struct {
	serviceHash uint32
	service     ChallengeService
}

// NewChallengeServiceHandler creates a new ChallengeServiceHandler
func NewChallengeServiceHandler(service ChallengeService) *ChallengeServiceHandler {
	return &ChallengeServiceHandler{
		serviceHash: ChallengeServiceOriginalFullyQualifiedDescriptorNameHash,
		service:     service,
	}
}

// GetServiceHash returns the service hash
func (h *ChallengeServiceHandler) GetServiceHash() uint32 {
	return h.serviceHash
}

// GetServiceFullName returns the service
func (h *ChallengeServiceHandler) GetServiceFullName() string {
	return ChallengeServiceOriginalFullyQualifiedDescriptorName
}

// GetServiceFullName returns the service
func (h *ChallengeServiceHandler) GetServiceName() string {
	return ChallengeServiceOriginalFullyQualifiedDescriptorName[strings.LastIndex(ChallengeServiceOriginalFullyQualifiedDescriptorName, ".")+1:]
}

// GetService returns the service
func (h *ChallengeServiceHandler) GetService() ChallengeService {
	return h.service
}

// Call calls the service by hash and method id on the handlers service
func (h *ChallengeServiceHandler) Call(ctx context.Context, methodId uint32, body []byte, conn interface{}) error {
	switch methodId {
	case 1:
		var req ChallengePickedRequest
		if err := proto.Unmarshal(body, &req); err != nil {
			return err
		}
		return h.service.ChallengePicked(ctx, &req, conn)
	case 2:
		var req ChallengeAnsweredRequest
		if err := proto.Unmarshal(body, &req); err != nil {
			return err
		}
		return h.service.ChallengeAnswered(ctx, &req, conn)
	case 3:
		var req ChallengeCancelledRequest
		if err := proto.Unmarshal(body, &req); err != nil {
			return err
		}
		return h.service.ChallengeCancelled(ctx, &req, conn)
	case 4:
		var req SendChallengeToUserRequest
		if err := proto.Unmarshal(body, &req); err != nil {
			return err
		}
		return h.service.SendChallengeToUser(ctx, &req, conn)
	default:
		return fmt.Errorf("unknown method id %d", methodId)
	}
}

// ChallengeListenerHandler is the handler interface for bgs.protocol.challenge.v1.ChallengeListener
type ChallengeListenerHandler struct {
	serviceHash uint32
	service     ChallengeListener
}

// NewChallengeListenerHandler creates a new ChallengeListenerHandler
func NewChallengeListenerHandler(service ChallengeListener) *ChallengeListenerHandler {
	return &ChallengeListenerHandler{
		serviceHash: ChallengeListenerOriginalFullyQualifiedDescriptorNameHash,
		service:     service,
	}
}

// GetServiceHash returns the service hash
func (h *ChallengeListenerHandler) GetServiceHash() uint32 {
	return h.serviceHash
}

// GetServiceFullName returns the service
func (h *ChallengeListenerHandler) GetServiceFullName() string {
	return ChallengeListenerOriginalFullyQualifiedDescriptorName
}

// GetServiceFullName returns the service
func (h *ChallengeListenerHandler) GetServiceName() string {
	return ChallengeListenerOriginalFullyQualifiedDescriptorName[strings.LastIndex(ChallengeListenerOriginalFullyQualifiedDescriptorName, ".")+1:]
}

// GetService returns the service
func (h *ChallengeListenerHandler) GetService() ChallengeListener {
	return h.service
}

// Call calls the service by hash and method id on the handlers service
func (h *ChallengeListenerHandler) Call(ctx context.Context, methodId uint32, body []byte, conn interface{}) error {
	switch methodId {
	case 1:
		var req ChallengeUserRequest
		if err := proto.Unmarshal(body, &req); err != nil {
			return err
		}
		return h.service.OnChallengeUser(ctx, &req, conn)
	case 2:
		var req ChallengeResultRequest
		if err := proto.Unmarshal(body, &req); err != nil {
			return err
		}
		return h.service.OnChallengeResult(ctx, &req, conn)
	case 3:
		var req ChallengeExternalRequest
		if err := proto.Unmarshal(body, &req); err != nil {
			return err
		}
		return h.service.OnExternalChallenge(ctx, &req, conn)
	case 4:
		var req ChallengeExternalResult
		if err := proto.Unmarshal(body, &req); err != nil {
			return err
		}
		return h.service.OnExternalChallengeResult(ctx, &req, conn)
	default:
		return fmt.Errorf("unknown method id %d", methodId)
	}
}