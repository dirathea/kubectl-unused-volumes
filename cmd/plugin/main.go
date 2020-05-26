package main

import (
	"github.com/dirathea/kubectl-unused-volumes/cmd/plugin/cli"
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure" // required for Azure
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"   // required for GKE
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	_ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {
	cli.InitAndExecute()
}
