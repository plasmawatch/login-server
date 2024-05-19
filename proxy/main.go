package main

import (
	"bnet-mock/network"
	"bnet-mock/network/client/attribute_types"
	"bnet-mock/network/client/authentication_service"
	"bnet-mock/network/client/connection_service"
	"bnet-mock/network/client/game_utilities_service"
	"bnet-mock/services"
	"bnet-mock/webserver"
	"bytes"
	"crypto/tls"
	"encoding/hex"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})

	log.SetLevel(log.DebugLevel)
}

func main() {
	cer, err := tls.LoadX509KeyPair("bnet-emu.fish.crt", "bnet-emu.fish.key")
	if err != nil {
		log.Println(err)
		return
	}

	go webserver.StartWebServer(cer)

	config := tls.Config{
		Certificates:       []tls.Certificate{cer},
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS12,
		// ServerName:         "bnet-emu.fish",
	}

	rpcServer := network.NewRpcServer(network.NewRpcServerConfig("127.0.0.1", 1119, &config), onClientAccepted)

	registerRpcServices(rpcServer)

	go func() {
		if err := rpcServer.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt)
	<-exitChan

	if err := rpcServer.Stop(); err != nil {
		log.Fatal(err)
	}
}

func registerRpcServices(server *network.RpcServer) {
	authHandler := authentication_service.NewAuthenticationServiceHandler(&services.AuthenticationServiceHandler{})
	server.AddService(authHandler.GetServiceHash(), authHandler)

	connectionService := connection_service.NewConnectionServiceHandler(&services.ConnectionServiceHandler{})
	server.AddService(connectionService.GetServiceHash(), connectionService)

	gameUtilsService := game_utilities_service.NewGameUtilitiesServiceHandler(&services.GameUtilitiesServiceHandler{})
	server.AddService(gameUtilsService.GetServiceHash(), gameUtilsService)
}

func onClientAccepted(conn *network.RpcClient, server *network.RpcServer) {
	defer conn.Close()

	buf := make([]byte, 4096)

	proxyServerAddre, err := tls.Dial("tcp", "us.actual.battle.net:1119", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Error("Error connecting to proxy", err)
		return
	}

	defer proxyServerAddre.Close()

	go func() {
		for {
			n, err := proxyServerAddre.Read(buf)
			if err != nil {
				log.Println(err)
				return
			}

			if _, err := conn.Write(buf[:n]); err != nil {
				log.Error("Error writing to client: ", err)
				return
			}

			header, err := network.NewBNetPacket(buf[:n])
			if err != nil {
				log.Error("Error parsing packet header: ", err)
				return
			}

			log.Debugf("From: Service: %s", header.Header.String())

			println(hex.Dump(buf[:n]))
		}
	}()

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		println(hex.Dump(buf[:n]))

		header, err := network.NewBNetPacket(buf[:n])
		if err != nil {
			log.Error("Error parsing packet header: ", err)
			return
		}

		log.Debug("To: Received packet:", header.GetHeader().String())

		if header.GetHeader().GetServiceHash() == game_utilities_service.GameUtilitiesServiceOriginalFullyQualifiedDescriptorNameHash {
			log.Debug("To: Handling game utilities packet size: ", header.GetHeader().GetSize())

			handleGameUtilitiesPacket(header, buf[:n])

			b, err := header.Marshal()
			if err != nil {
				log.Error("Error marshalling packet: ", err)
				return
			}

			if bytes.Compare(b, buf[:n]) == 0 {
				log.Debug("To: Packet did not change, sending original packet")

				if _, err := proxyServerAddre.Write(buf[:n]); err != nil {
					log.Error("Error writing to proxy server: ", err)
					return
				}
			}

			log.Debug("To: Handling new game utilities packet size: ", header.GetHeader().GetSize())

			if _, err := proxyServerAddre.Write(b); err != nil {
				log.Error("Error writing to proxy server: ", err)
				return
			}
		} else {
			if _, err := proxyServerAddre.Write(buf[:n]); err != nil {
				log.Error("Error writing to proxy server: ", err)
				return
			}
		}
	}
}

func handleGameUtilitiesPacket(header *network.BNetPacket, bytes []byte) {
	log.Println("Handling game utilities packet")

	if header.GetHeader().GetMethodId() == 1 {
		var gameUtilitiesRequest game_utilities_service.ClientRequest
		if err := header.UnmarshalBody(&gameUtilitiesRequest); err != nil {
			log.Error("Error unmarshalling game utilities request: ", err)
			return
		}

		log.Debug("To: Game utilities request: ", gameUtilitiesRequest.String())

		for i := range gameUtilitiesRequest.GetAttribute() {
			if gameUtilitiesRequest.GetAttribute()[i].GetName() == "build" {
				gameUtilitiesRequest.GetAttribute()[i].Value = &attribute_types.Variant{
					UintValue: proto.Uint64(114992),
				}
			}
		}

		log.Println("To: Game utilities response: ", gameUtilitiesRequest.String())

		if err := header.MarshalBody(&gameUtilitiesRequest); err != nil {
			log.Error("Error marshalling game utilities request: ", err)
			return
		}
	}
}
