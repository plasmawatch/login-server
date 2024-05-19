package main

import (
	"bnet-mock/network"
	"bnet-mock/network/client/authentication_service"
	"bnet-mock/network/client/connection_service"
	"bnet-mock/network/client/game_utilities_service"
	"bnet-mock/services"
	"bnet-mock/webserver"
	"context"
	"crypto/tls"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
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

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		header, err := network.NewBNetPacket(buf[:n])
		if err != nil {
			log.Error("Error parsing packet header: ", err)
			return
		}

		serviceHandler := server.GetService(header.GetHeader().GetServiceHash())
		if serviceHandler == nil {
			log.Warnf("Service handler not found for service ID: %d", header.GetHeader().GetServiceHash())
			continue
		}

		log.Debug("Received packet:", header.GetHeader().String())
		log.Debugf("Service: %s, Service ID: %d, Method ID: %d", serviceHandler.GetServiceFullName(), header.GetHeader().GetServiceHash(), header.GetHeader().GetMethodId())

		if err := serviceHandler.Call(context.Background(), header.GetHeader().GetMethodId(), header.Body, conn); err != nil {
			log.Error("Error calling service handler: ", err)
			return
		}

		println("")
	}
}
