module github.com/darkowlzz/csi-toolkit/example

go 1.15

require (
	github.com/container-storage-interface/spec v1.3.0
	github.com/darkowlzz/csi-toolkit v0.0.0
	github.com/golang/protobuf v1.4.3
	google.golang.org/grpc v1.36.0
	sigs.k8s.io/controller-runtime v0.8.2
)

replace github.com/darkowlzz/csi-toolkit => ../
