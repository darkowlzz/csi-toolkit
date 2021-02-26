package driver

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"

	"github.com/darkowlzz/csi-toolkit/node"
)

func (md *MyDriver) NodePublishVolumeHandler(context.Context, *csi.NodePublishVolumeResponse, *csi.NodePublishVolumeRequest) error {
	log.Info("node publish volume handler")

	return nil
}

func (md *MyDriver) NodeUnpublishVolumeHandler(context.Context, *csi.NodeUnpublishVolumeResponse, *csi.NodeUnpublishVolumeRequest) error {
	log.Info("node unpublish volume handler")

	return nil
}

func (md *MyDriver) NodeStageVolumeHandler(context.Context, *csi.NodeStageVolumeResponse, *csi.NodeStageVolumeRequest) error {
	log.Info("node stage volume handler")

	return nil
}

func (md *MyDriver) NodeUnstageVolumeHandler(context.Context, *csi.NodeUnstageVolumeResponse, *csi.NodeUnstageVolumeRequest) error {
	log.Info("node unstage volume handler")

	return nil
}

func (md *MyDriver) NodeGetInfoHandler(ctx context.Context, resp *csi.NodeGetInfoResponse, req *csi.NodeGetInfoRequest) error {
	log.Info("node get info handler")

	resp.NodeId = md.NodeName
	resp.MaxVolumesPerNode = int64(10)
	resp.AccessibleTopology = &csi.Topology{
		Segments: md.Topology,
	}

	return nil
}

func (md *MyDriver) NodeGetCapabilitiesHandler(ctx context.Context, resp *csi.NodeGetCapabilitiesResponse, req *csi.NodeGetCapabilitiesRequest) error {
	log.Info("node get capabilities handler")

	resp.Capabilities = node.GetServiceCapabilities(
		csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME,
		// csi.NodeServiceCapability_RPC_EXPAND_VOLUME,
		// csi.NodeServiceCapability_RPC_VOLUME_CONDITION,
	)

	return nil
}

func (md *MyDriver) NodeGetVolumeStatsHandler(context.Context, *csi.NodeGetVolumeStatsResponse, *csi.NodeGetVolumeStatsRequest) error {
	log.Info("node get volume stats handler")

	return nil
}

func (md *MyDriver) NodeExpandVolume(context.Context, *csi.NodeExpandVolumeResponse, *csi.NodeExpandVolumeRequest) error {
	log.Info("node expand volume handler")

	return nil
}
