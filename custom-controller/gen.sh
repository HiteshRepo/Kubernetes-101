# ROOT_PACKAGE :: the package (relative to $GOPATH/src) that is the target for code generation
ROOT_PACKAGE=/home/hitesh/Documents/personal/codebases/Kubernetes-101/custom-controller
# CUSTOM_RESOURCE_NAME :: the name of the custom resource that we're generating client code for
CUSTOM_RESOURCE_NAME="flightticket"
# CUSTOM_RESOURCE_VERSION :: the version of the resource
CUSTOM_RESOURCE_VERSION="v1"

# retrieve the code-generator scripts and bins
go get -u sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.5
cd $GOPATH/pkg/mod/github.com/kubernetes-sigs/controller-tools@v0.12.0/


# run the code-generator entrypoint script
./test.sh all "$ROOT_PACKAGE/pkg/client" "$ROOT_PACKAGE/pkg/apis" "$CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION"
