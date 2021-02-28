package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/darkowlzz/operator-toolkit/runnable"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/darkowlzz/csi-toolkit/interceptor"
)

var log = ctrl.Log.WithName("grpc-server")

var defaultEndpoint = "unix://tmp/csi.sock"

// Server is a CSI server.
type Server struct {
	Options

	// Graceful runnable helps implement Runnable interface for starting the
	// server with graceful shutdown.
	// NOTE: Runnables are components that controller manager can manage.
	*runnable.Graceful

	// server is the CSI grpc server.
	server *grpc.Server
}

// Options for the server.
type Options struct {
	// Endpoint is the server endpoint.
	Endpoint string
	// IDS is the IdentityServer.
	IDS csi.IdentityServer
	// CS is the ControllerServer.
	CS csi.ControllerServer
	// NS is the NodeServer.
	NS csi.NodeServer
	// RequireLeaderElection can be set to start the server after a leader
	// election.
	RequireLeaderElection bool
	// UnaryInterceptors are the interceptors to be added in the grpc server.
	// The order of the interceptors form a chain of unary interceptors.
	// By default a LogGRPC interceptor is used if no interceptor is provided.
	// To use LogGRPC interceptor with other interceptors, it should be
	// explicitly included when setting the value.
	UnaryInterceptors []grpc.UnaryServerInterceptor
}

// setDefaults sets the defaul options for the Server.
func (o *Options) setDefaults() {
	if len(o.Endpoint) == 0 {
		o.Endpoint = defaultEndpoint
	}

	// If no unary interceptor is set, add the default LogGRPC interceptor.
	if len(o.UnaryInterceptors) == 0 {
		o.UnaryInterceptors = append(o.UnaryInterceptors, interceptor.LogGRPC)
	}
}

// NewServer creates a new server with graceful shutdown support.
func NewServer(ops Options, wg *sync.WaitGroup) *Server {
	ops.setDefaults()

	s := &Server{
		Options: ops,
	}

	// Create a graceful runnable with run and stop of the server.
	s.Graceful = runnable.NewGraceful(s.Run, s.Stop, ops.RequireLeaderElection, wg, log)

	return s
}

// Run starts the GRPC server. It is a blocking function.
func (s *Server) Run(ctx context.Context) error {
	proto, addr, err := parseEndpoint(s.Endpoint)
	if err != nil {
		return err
	}

	if proto == "unix" {
		addr = "/" + addr
		if err := os.Remove(addr); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to remove %s, error: %w", addr, err)
		}
	}

	listener, err := net.Listen(proto, addr)
	if err != nil {
		return fmt.Errorf("failed to listen, error: %w", err)
	}

	// Chain the unary interceptors.
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(s.Options.UnaryInterceptors...),
		),
	}
	server := grpc.NewServer(opts...)
	s.server = server

	if s.IDS != nil {
		csi.RegisterIdentityServer(server, s.IDS)
	}
	if s.CS != nil {
		csi.RegisterControllerServer(server, s.CS)
	}
	if s.NS != nil {
		csi.RegisterNodeServer(server, s.NS)
	}

	log.Info("listening for connections", "address", listener.Addr())

	// This is a blocking call.
	return s.server.Serve(listener)
}

// Stop stops the GRPC server gracefully.
func (s *Server) Stop() error {
	s.server.GracefulStop()
	return nil
}

// parseEndpoint parses a given endpoint and returns the protocol and path.
func parseEndpoint(ep string) (string, string, error) {
	if strings.HasPrefix(strings.ToLower(ep), "unix://") || strings.HasPrefix(strings.ToLower(ep), "tcp://") {
		s := strings.SplitN(ep, "://", 2)
		if s[1] != "" {
			return s[0], s[1], nil
		}
	}
	return "", "", fmt.Errorf("Invalid endpoint: %v", ep)
}
