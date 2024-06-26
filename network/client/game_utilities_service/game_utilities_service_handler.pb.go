// Code generated by protoc-gen-go-bnet. DO NOT EDIT.

package game_utilities_service

import (
	"context"
	"fmt"
	"strings"
	"google.golang.org/protobuf/proto"
)

// GameUtilitiesServiceHandler is the handler interface for bgs.protocol.game_utilities.v1.GameUtilitiesService
type GameUtilitiesServiceHandler struct {
	serviceHash uint32
	service     GameUtilitiesService
}

// NewGameUtilitiesServiceHandler creates a new GameUtilitiesServiceHandler
func NewGameUtilitiesServiceHandler(service GameUtilitiesService) *GameUtilitiesServiceHandler {
	return &GameUtilitiesServiceHandler{
		serviceHash: GameUtilitiesServiceOriginalFullyQualifiedDescriptorNameHash,
		service:     service,
	}
}

// GetServiceHash returns the service hash
func (h *GameUtilitiesServiceHandler) GetServiceHash() uint32 {
	return h.serviceHash
}

// GetServiceFullName returns the service
func (h *GameUtilitiesServiceHandler) GetServiceFullName() string {
	return GameUtilitiesServiceOriginalFullyQualifiedDescriptorName
}

// GetServiceFullName returns the service
func (h *GameUtilitiesServiceHandler) GetServiceName() string {
	return GameUtilitiesServiceOriginalFullyQualifiedDescriptorName[strings.LastIndex(GameUtilitiesServiceOriginalFullyQualifiedDescriptorName, ".")+1:]
}

// GetService returns the service
func (h *GameUtilitiesServiceHandler) GetService() GameUtilitiesService {
	return h.service
}

// Call calls the service by hash and method id on the handlers service
func (h *GameUtilitiesServiceHandler) Call(ctx context.Context, methodId uint32, body []byte, conn interface{}) error {
	switch methodId {
	case 1:
		var req ClientRequest
		if err := proto.Unmarshal(body, &req); err != nil {
			return err
		}
		return h.service.ProcessClientRequest(ctx, &req, conn)
	case 2:
		var req PresenceChannelCreatedRequest
		if err := proto.Unmarshal(body, &req); err != nil {
			return err
		}
		return h.service.PresenceChannelCreated(ctx, &req, conn)
	case 3:
		var req GetPlayerVariablesRequest
		if err := proto.Unmarshal(body, &req); err != nil {
			return err
		}
		return h.service.GetPlayerVariables(ctx, &req, conn)
	case 6:
		var req ServerRequest
		if err := proto.Unmarshal(body, &req); err != nil {
			return err
		}
		return h.service.ProcessServerRequest(ctx, &req, conn)
	case 7:
		var req GameAccountOnlineNotification
		if err := proto.Unmarshal(body, &req); err != nil {
			return err
		}
		return h.service.OnGameAccountOnline(ctx, &req, conn)
	case 8:
		var req GameAccountOfflineNotification
		if err := proto.Unmarshal(body, &req); err != nil {
			return err
		}
		return h.service.OnGameAccountOffline(ctx, &req, conn)
	case 9:
		var req GetAchievementsFileRequest
		if err := proto.Unmarshal(body, &req); err != nil {
			return err
		}
		return h.service.GetAchievementsFile(ctx, &req, conn)
	case 10:
		var req GetAllValuesForAttributeRequest
		if err := proto.Unmarshal(body, &req); err != nil {
			return err
		}
		return h.service.GetAllValuesForAttribute(ctx, &req, conn)
	default:
		return fmt.Errorf("unknown method id %d", methodId)
	}
}
