package controller

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

type ControllerInterface interface {
	CreateVolumeHandler(context.Context, *csi.CreateVolumeResponse, *csi.CreateVolumeRequest) error
	DeleteVolumeHandler(context.Context, *csi.DeleteVolumeResponse, *csi.DeleteVolumeRequest) error
	GetCapabilitiesHandler(context.Context, *csi.ControllerGetCapabilitiesResponse, *csi.ControllerGetCapabilitiesRequest) error
	ValidateVolumeCapabilitiesHandler(context.Context, *csi.ValidateVolumeCapabilitiesResponse, *csi.ValidateVolumeCapabilitiesRequest) error
	ControllerPublishVolumeHandler(context.Context, *csi.ControllerPublishVolumeResponse, *csi.ControllerPublishVolumeRequest) error
	ControllerUnpublishVolumeHandler(context.Context, *csi.ControllerUnpublishVolumeResponse, *csi.ControllerUnpublishVolumeRequest) error
	GetCapacityHandler(context.Context, *csi.GetCapacityResponse, *csi.GetCapacityRequest) error
	ListVolumesHandler(context.Context, *csi.ListVolumesResponse, *csi.ListVolumesRequest) error
	ControllerGetVolumeHandler(context.Context, *csi.ControllerGetVolumeResponse, *csi.ControllerGetVolumeRequest) error
}
