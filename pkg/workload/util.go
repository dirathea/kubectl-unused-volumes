package workload

import (
	"strings"

	"github.com/dirathea/kubectl-unused-volumes/pkg/api"
)

func Join(workloads []api.Workload, sep string) string {
	allNames := []string{}

	for _, wk := range workloads {
		allNames = append(allNames, wk.GetName())
	}
	return strings.Join(allNames, sep)
}
