package identity

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	ctrl "sigs.k8s.io/controller-runtime"
)

var log = ctrl.Log.WithName("identity-service")

type Identity struct {
	ident IdentityInterface
}

// Ensure Identity implements the IdentityServer interface.
var _ csi.IdentityServer = &Identity{}

func New(ident IdentityInterface) *Identity {
	return &Identity{
		ident: ident,
	}
}

func (i *Identity) GetPluginInfo(ctx context.Context, req *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	log.Info("get plugin info request")

	resp := &csi.GetPluginInfoResponse{}

	if err := i.ident.GetPluginInfoHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}

func (i *Identity) Probe(ctx context.Context, req *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	log.Info("probe request")

	resp := &csi.ProbeResponse{}

	if err := i.ident.ProbeHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}

func (i *Identity) GetPluginCapabilities(ctx context.Context, req *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	log.Info("get plugin capabilities")

	resp := &csi.GetPluginCapabilitiesResponse{}

	if err := i.ident.GetPluginCapabilitiesHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}
