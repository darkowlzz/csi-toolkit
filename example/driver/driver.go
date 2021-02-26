package driver

import (
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/darkowlzz/csi-toolkit/controller"
	"github.com/darkowlzz/csi-toolkit/identity"
	"github.com/darkowlzz/csi-toolkit/node"
)

var log = ctrl.Log.WithName("mydriver")

type MyDriver struct {
	Options

	NodeName string
}

type Options struct {
	Client   client.Client
	Topology map[string]string

	// Storage provider API client.
}

// New sets up a new driver.
func New(nodeName string, opts Options) *MyDriver {
	md := &MyDriver{
		NodeName: nodeName,
		Options:  opts,
	}

	// Set up some node local components.

	return md
}

func (md *MyDriver) GetIdentityService() *identity.Identity {
	return identity.New(md)
}

func (md *MyDriver) GetControllerService() *controller.Controller {
	return controller.New(md)
}

func (md *MyDriver) GetNodeService() *node.Node {
	return node.New(md)
}
