package driver

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/protobuf/ptypes/wrappers"

	"github.com/darkowlzz/csi-toolkit/identity"
)

func (md *MyDriver) GetPluginInfoHandler(ctx context.Context, resp *csi.GetPluginInfoResponse, req *csi.GetPluginInfoRequest) error {
	log.Info("get plugin info handler")

	resp.Name = "example.driver"
	resp.VendorVersion = "unknown"

	return nil
}

func (md *MyDriver) ProbeHandler(ctx context.Context, resp *csi.ProbeResponse, req *csi.ProbeRequest) error {
	log.Info("get plugin info handler")

	resp.Ready = &wrappers.BoolValue{Value: true}

	return nil
}

func (md *MyDriver) GetPluginCapabilitiesHandler(ctx context.Context, resp *csi.GetPluginCapabilitiesResponse, req *csi.GetPluginCapabilitiesRequest) error {
	log.Info("get plugin capabilities handler")

	resp.Capabilities = identity.GetPluginCapabilities(csi.PluginCapability_Service_CONTROLLER_SERVICE)

	return nil
}
