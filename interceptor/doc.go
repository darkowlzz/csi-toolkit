// Package interceptor provides GRPC UnaryInterceptors to use with a GRPC
// server.
package interceptor

import (
	ctrl "sigs.k8s.io/controller-runtime"
)

var log = ctrl.Log.WithName("grpc-server")
