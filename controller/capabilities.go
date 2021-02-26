package controller

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
)

// GetServiceCapabilities returns a collection of all the Controller Service
// Capabilities.
func GetServiceCapabilities(caps ...csi.ControllerServiceCapability_RPC_Type) []*csi.ControllerServiceCapability {
	var csc []*csi.ControllerServiceCapability

	for _, cap := range caps {
		csc = append(csc, &csi.ControllerServiceCapability{
			Type: &csi.ControllerServiceCapability_Rpc{
				Rpc: &csi.ControllerServiceCapability_RPC{
					Type: cap,
				},
			},
		})
	}

	return csc
}
