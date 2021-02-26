package node

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

type NodeInterface interface {
	NodePublishVolumeHandler(context.Context, *csi.NodePublishVolumeResponse, *csi.NodePublishVolumeRequest) error
	NodeUnpublishVolumeHandler(context.Context, *csi.NodeUnpublishVolumeResponse, *csi.NodeUnpublishVolumeRequest) error
	NodeStageVolumeHandler(context.Context, *csi.NodeStageVolumeResponse, *csi.NodeStageVolumeRequest) error
	NodeUnstageVolumeHandler(context.Context, *csi.NodeUnstageVolumeResponse, *csi.NodeUnstageVolumeRequest) error
	NodeGetInfoHandler(context.Context, *csi.NodeGetInfoResponse, *csi.NodeGetInfoRequest) error
	NodeGetCapabilitiesHandler(context.Context, *csi.NodeGetCapabilitiesResponse, *csi.NodeGetCapabilitiesRequest) error
	NodeGetVolumeStatsHandler(context.Context, *csi.NodeGetVolumeStatsResponse, *csi.NodeGetVolumeStatsRequest) error
	NodeExpandVolume(context.Context, *csi.NodeExpandVolumeResponse, *csi.NodeExpandVolumeRequest) error
}
