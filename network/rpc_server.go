package network

import (
	"context"
	"crypto/tls"
	"errors"
	log "github.com/sirupsen/logrus"
	"net"
	"strconv"
	"time"
)

type RpcServiceHandler interface {
	GetServiceName() string
	GetServiceHash() uint32
	GetServiceFullName() string
	Call(ctx context.Context, methodId uint32, body []byte, i interface{}) error
}

type RpcServerConfig struct {
	Host      string
	Port      int
	tlsConfig *tls.Config
}

func NewRpcServerConfig(host string, port int, tlsConfig *tls.Config) *RpcServerConfig {
	return &RpcServerConfig{
		Host:      host,
		Port:      port,
		tlsConfig: tlsConfig,
	}
}

func (c *RpcServerConfig) GetHost() string {
	return c.Host
}

func (c *RpcServerConfig) GetPort() int {
	return c.Port
}

func (c *RpcServerConfig) GetTlsConfig() *tls.Config {
	return c.tlsConfig
}

type RpcServer struct {
	config   *RpcServerConfig
	listener net.Listener
	onAccept func(client *RpcClient, server *RpcServer)
	services map[uint32]RpcServiceHandler
}

func NewRpcServer(config *RpcServerConfig, onAccept func(client *RpcClient, server *RpcServer)) *RpcServer {
	return &RpcServer{
		config:   config,
		onAccept: onAccept,
		services: make(map[uint32]RpcServiceHandler),
	}
}

func (s *RpcServer) AddService(serviceId uint32, service RpcServiceHandler) {
	s.services[serviceId] = service
}

func (s *RpcServer) GetService(serviceId uint32) RpcServiceHandler {
	return s.services[serviceId]
}

// Start the server
func (s *RpcServer) Start() error {
	if len(s.services) == 0 || s.onAccept == nil || s.config == nil {
		return errors.New("RpcServer: Start: services, onAccept or config is nil")
	}

	for _, service := range s.services {
		log.Debugf("Registered service: %s (%d)", service.GetServiceFullName(), service.GetServiceHash())
	}

	// listen on a port
	listener, err := tls.Listen("tcp", s.config.GetHost()+":"+strconv.Itoa(s.config.GetPort()), s.config.GetTlsConfig())
	if err != nil {
		return err
	}

	s.listener = listener

	// close the listener when the application closes
	defer listener.Close()

	log.Infof("Listening on %s:%d", s.config.GetHost(), s.config.GetPort())

	for {
		// accept a connection
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		client := NewRpcClient(conn, uint32(time.Now().Unix()), s)

		// handle the connection
		go s.onAccept(client, s)
	}
}

// Stop the server
func (s *RpcServer) Stop() error {
	if s.listener != nil {
		return s.listener.Close()
	}

	return nil
}

func (s *RpcServer) SendPacket(conn net.Conn, packet *BNetPacket) error {
	if packet == nil {
		return errors.New("RpcServer: SendPacket: packet is nil")
	}

	if conn == nil {
		return errors.New("RpcServer: SendPacket: conn is nil")
	}

	buf, err := packet.Marshal()
	if err != nil {
		return err
	}

	_, err = conn.Write(buf)
	if err != nil {
		return err
	}

	return nil
}
