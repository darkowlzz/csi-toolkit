package identity

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
)

// GetServiceCapabilities returns a collection of all the Plugin Capabilities.
func GetPluginCapabilities(caps ...csi.PluginCapability_Service_Type) []*csi.PluginCapability {
	var pc []*csi.PluginCapability

	for _, cap := range caps {
		pc = append(pc, &csi.PluginCapability{
			Type: &csi.PluginCapability_Service_{
				Service: &csi.PluginCapability_Service{
					Type: cap,
				},
			},
		})
	}

	return pc
}
