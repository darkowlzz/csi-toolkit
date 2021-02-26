package node

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
)

// GetServiceCapabilities returns a collection of all the Node Service
// Capabilities.
func GetServiceCapabilities(caps ...csi.NodeServiceCapability_RPC_Type) []*csi.NodeServiceCapability {
	var nsc []*csi.NodeServiceCapability

	for _, cap := range caps {
		nsc = append(nsc, &csi.NodeServiceCapability{
			Type: &csi.NodeServiceCapability_Rpc{
				Rpc: &csi.NodeServiceCapability_RPC{
					Type: cap,
				},
			},
		})
	}

	return nsc
}
