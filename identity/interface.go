package identity

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

type IdentityInterface interface {
	GetPluginInfoHandler(context.Context, *csi.GetPluginInfoResponse, *csi.GetPluginInfoRequest) error
	ProbeHandler(context.Context, *csi.ProbeResponse, *csi.ProbeRequest) error
	GetPluginCapabilitiesHandler(context.Context, *csi.GetPluginCapabilitiesResponse, *csi.GetPluginCapabilitiesRequest) error
}
