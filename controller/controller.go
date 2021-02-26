package controller

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/pborman/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	ctrl "sigs.k8s.io/controller-runtime"
)

var log = ctrl.Log.WithName("controller-service")

type Controller struct {
	ctrlr ControllerInterface
}

func New(ctrlr ControllerInterface) *Controller {
	return &Controller{
		ctrlr: ctrlr,
	}
}

func (c *Controller) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (resp *csi.CreateVolumeResponse, err error) {
	log.Info("create volume request")

	// Check arguments
	if len(req.GetName()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Name missing in request")
	}
	caps := req.GetVolumeCapabilities()
	if caps == nil {
		return nil, status.Error(codes.InvalidArgument, "Volume Capabilities missing in request")
	}

	volumeID := uuid.NewUUID().String()

	resp = &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			VolumeId:      volumeID,
			CapacityBytes: req.GetCapacityRange().GetRequiredBytes(),
			VolumeContext: req.GetParameters(),
			// ContentSource:      req.GetVolumeContentSource(),
			// AccessibleTopology: topologies,
		},
	}

	if err := c.ctrlr.CreateVolumeHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Controller) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	// Check arguments
	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID missing in request")
	}

	resp := &csi.DeleteVolumeResponse{}

	if err := c.ctrlr.DeleteVolumeHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Controller) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	resp := &csi.ControllerGetCapabilitiesResponse{}

	if err := c.ctrlr.GetCapabilitiesHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Controller) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	// Check arguments
	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID cannot be empty")
	}
	if len(req.VolumeCapabilities) == 0 {
		return nil, status.Error(codes.InvalidArgument, req.VolumeId)
	}

	for _, cap := range req.GetVolumeCapabilities() {
		if cap.GetMount() == nil && cap.GetBlock() == nil {
			return nil, status.Error(codes.InvalidArgument, "cannot have both mount and block access type be undefined")
		}

		// A real driver would check the capabilities of the given volume with
		// the set of requested capabilities.
	}

	resp := &csi.ValidateVolumeCapabilitiesResponse{
		Confirmed: &csi.ValidateVolumeCapabilitiesResponse_Confirmed{
			VolumeContext:      req.GetVolumeContext(),
			VolumeCapabilities: req.GetVolumeCapabilities(),
			Parameters:         req.GetParameters(),
		},
	}

	if err := c.ctrlr.ValidateVolumeCapabilitiesHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Controller) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	resp := &csi.ControllerPublishVolumeResponse{}

	if err := c.ctrlr.ControllerPublishVolumeHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Controller) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	resp := &csi.ControllerUnpublishVolumeResponse{}

	if err := c.ctrlr.ControllerUnpublishVolumeHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Controller) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	resp := &csi.GetCapacityResponse{}

	if err := c.ctrlr.GetCapacityHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Controller) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	resp := &csi.ListVolumesResponse{}

	if err := c.ctrlr.ListVolumesHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Controller) ControllerGetVolume(ctx context.Context, req *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	resp := &csi.ControllerGetVolumeResponse{}

	if err := c.ctrlr.ControllerGetVolumeHandler(ctx, resp, req); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Controller) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (c *Controller) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (c *Controller) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (c *Controller) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
