package node

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	ctrl "sigs.k8s.io/controller-runtime"
)

var log = ctrl.Log.WithName("node-service")

type Node struct {
	node NodeInterface
}

// Ensure Node implements the NodeServer interface.
var _ csi.NodeServer = &Node{}

func New(node NodeInterface) *Node {
	return &Node{
		node: node,
	}
}

func (n *Node) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	log.Info("node publish volume")

	// Check arguments
	if req.GetVolumeCapability() == nil {
		return nil, status.Error(codes.InvalidArgument, "Volume capability missing in request")
	}
	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID missing in request")
	}
	if len(req.GetTargetPath()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Target path missing in request")
	}

	if req.GetVolumeCapability().GetBlock() != nil &&
		req.GetVolumeCapability().GetMount() != nil {
		return nil, status.Error(codes.InvalidArgument, "cannot have both block and mount access type")
	}

	resp := &csi.NodePublishVolumeResponse{}

	if err := n.node.NodePublishVolumeHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}

func (n *Node) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	log.Info("node unpublish volume")

	// Check arguments
	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID missing in request")
	}
	if len(req.GetTargetPath()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Target path missing in request")
	}

	resp := &csi.NodeUnpublishVolumeResponse{}

	if err := n.node.NodeUnpublishVolumeHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}

func (n *Node) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	log.Info("node stage volume")

	// Check arguments
	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID missing in request")
	}
	if len(req.GetStagingTargetPath()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Target path missing in request")
	}
	if req.GetVolumeCapability() == nil {
		return nil, status.Error(codes.InvalidArgument, "Volume Capability missing in request")
	}

	resp := &csi.NodeStageVolumeResponse{}

	if err := n.node.NodeStageVolumeHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}

func (n *Node) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	log.Info("node unstage volume")

	// Check arguments
	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID missing in request")
	}
	if len(req.GetStagingTargetPath()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Target path missing in request")
	}

	resp := &csi.NodeUnstageVolumeResponse{}

	if err := n.node.NodeUnstageVolumeHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}

func (n *Node) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	log.Info("node get info")

	resp := &csi.NodeGetInfoResponse{}

	if err := n.node.NodeGetInfoHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}

func (n *Node) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	log.Info("node get capabilities")

	resp := &csi.NodeGetCapabilitiesResponse{}

	if err := n.node.NodeGetCapabilitiesHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}

func (n *Node) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	log.Info("node get volume stats")

	resp := &csi.NodeGetVolumeStatsResponse{}

	if err := n.node.NodeGetVolumeStatsHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}

func (n *Node) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	log.Info("node expand volume")

	volID := req.GetVolumeId()
	if len(volID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID not provided")
	}

	volPath := req.GetVolumePath()
	if len(volPath) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume path not provided")
	}

	resp := &csi.NodeExpandVolumeResponse{}

	if err := n.node.NodeExpandVolume(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}
