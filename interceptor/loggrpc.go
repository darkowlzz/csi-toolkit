// This LogGRPC UnaryServerInterceptor is taken from
// https://github.com/kubernetes-csi/csi-driver-host-path/blob/v1.6.0/pkg/hostpath/server.go#L116.
package interceptor

import (
	"context"

	"github.com/kubernetes-csi/csi-lib-utils/protosanitizer"
	"google.golang.org/grpc"
	ctrl "sigs.k8s.io/controller-runtime"
)

// LogGRPC handles logging of the GRPC requests. The log messages are sanitized
// to remove any sensitive data.
func LogGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
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

// LogGRPCWithName allows setting the logger name and values in LogGRPC logs.
func LogGRPCWithNameAndValues(name string, tags ...interface{}) grpc.UnaryServerInterceptor {
	log = ctrl.Log.WithName(name).WithValues(tags...)
	return LogGRPC
}
