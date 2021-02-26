package driver

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"

	"github.com/darkowlzz/csi-toolkit/controller"
)

func (md *MyDriver) CreateVolumeHandler(ctx context.Context, resp *csi.CreateVolumeResponse, req *csi.CreateVolumeRequest) error {
	log.Info("create volume handler")
	// resp.Volume.VolumeId = "foo"

	return nil
}

func (md *MyDriver) DeleteVolumeHandler(ctx context.Context, resp *csi.DeleteVolumeResponse, req *csi.DeleteVolumeRequest) error {
	return nil
}

func (md *MyDriver) GetCapabilitiesHandler(ctx context.Context, resp *csi.ControllerGetCapabilitiesResponse, req *csi.ControllerGetCapabilitiesRequest) error {
	log.Info("get capabilities handler")
	resp.Capabilities = controller.GetServiceCapabilities(
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_GET_VOLUME,
		// csi.ControllerServiceCapability_RPC_GET_CAPACITY,
		// csi.ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT,
		// csi.ControllerServiceCapability_RPC_LIST_SNAPSHOTS,
		// csi.ControllerServiceCapability_RPC_LIST_VOLUMES,
		// csi.ControllerServiceCapability_RPC_CLONE_VOLUME,
		// csi.ControllerServiceCapability_RPC_EXPAND_VOLUME,
		// csi.ControllerServiceCapability_RPC_VOLUME_CONDITION,
	)

	return nil
}

func (md *MyDriver) ValidateVolumeCapabilitiesHandler(context.Context, *csi.ValidateVolumeCapabilitiesResponse, *csi.ValidateVolumeCapabilitiesRequest) error {
	return nil
}

func (md *MyDriver) ControllerPublishVolumeHandler(context.Context, *csi.ControllerPublishVolumeResponse, *csi.ControllerPublishVolumeRequest) error {
	return nil
}

func (md *MyDriver) ControllerUnpublishVolumeHandler(context.Context, *csi.ControllerUnpublishVolumeResponse, *csi.ControllerUnpublishVolumeRequest) error {
	return nil
}

func (md *MyDriver) GetCapacityHandler(context.Context, *csi.GetCapacityResponse, *csi.GetCapacityRequest) error {
	return nil
}

func (md *MyDriver) ListVolumesHandler(context.Context, *csi.ListVolumesResponse, *csi.ListVolumesRequest) error {
	return nil
}

func (md *MyDriver) ControllerGetVolumeHandler(context.Context, *csi.ControllerGetVolumeResponse, *csi.ControllerGetVolumeRequest) error {
	return nil
}
