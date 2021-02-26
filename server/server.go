package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/csi-lib-utils/protosanitizer"
	"google.golang.org/grpc"
	ctrl "sigs.k8s.io/controller-runtime"
)

var log = ctrl.Log.WithName("grpc-server")

var defaultEndpoint = "unix://tmp/csi.sock"

// Server is CSI server.
type Server struct {
	Options

	server *grpc.Server
}

// Options for the server.
type Options struct {
	Endpoint string
	IDS      csi.IdentityServer
	CS       csi.ControllerServer
	NS       csi.NodeServer
}

func (o *Options) setDefaults() {
	if len(o.Endpoint) == 0 {
		o.Endpoint = defaultEndpoint
	}
}

// NewServer creates a new server.
func NewServer(ops Options) *Server {
	ops.setDefaults()

	return &Server{
		Options: ops,
	}
}

// NeedLeaderElection implements the LeaderElectionRunnable interface, which
// indicates the gRPC server doesn't need leader election.
func (*Server) NeedLeaderElection() bool {
	return false
}

// Start starts the GRPC server.
func (s *Server) Start(ctx context.Context) error {
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

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(logGRPC),
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

	// Start the GRPC server.
	// TODO: Better error handling.
	go func() {
		// This is blocking.
		if err := s.server.Serve(listener); err != nil {
			log.Error(err, "failed to serve")
			return
		}
	}()

	go func() {
		for {
			<-ctx.Done()
			// TODO: Support graceful stop.
			// s.server.GracefulStop()
			s.server.Stop()
			return
		}
	}()

	return nil
}

func parseEndpoint(ep string) (string, string, error) {
	if strings.HasPrefix(strings.ToLower(ep), "unix://") || strings.HasPrefix(strings.ToLower(ep), "tcp://") {
		s := strings.SplitN(ep, "://", 2)
		if s[1] != "" {
			return s[0], s[1], nil
		}
	}
	return "", "", fmt.Errorf("Invalid endpoint: %v", ep)
}

func logGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if info.FullMethod == "/csi.v1.Identity/Probe" {
		return handler(ctx, req)
	}
	log.V(3).Info("GRPC call", "method", info.FullMethod)
	log.V(5).Info("GRPC request", "request", protosanitizer.StripSecrets(req))
	resp, err := handler(ctx, req)
	if err != nil {
		log.Error(err, "GRPC error")
	} else {
		log.V(5).Info("GRPC response", "response", protosanitizer.StripSecrets(resp))
	}
	return resp, err
}
