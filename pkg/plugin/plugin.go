package plugin

import (
	"context"
	"strings"

	"github.com/dirathea/kubectl-unused-volumes/pkg/api"
	"github.com/dirathea/kubectl-unused-volumes/pkg/workload"
	"github.com/gosuri/uitable"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	v1 "k8s.io/api/core/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

func RunPlugin(configFlags *genericclioptions.ConfigFlags) (output string, err error) {
	// log := logger.NewLogger()
	config, err := configFlags.ToRESTConfig()
	if err != nil {
		err = errors.Wrap(err, "failed to read kubeconfig")
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		err = errors.Wrap(err, "failed to create clientset")
		return
	}

	volumes, err := GetVolumes(clientset, *configFlags.Namespace)
	if err != nil {
		err = errors.Wrap(err, "failed to get all pvc in namespaces")
		return
	}

	allResources := []func(client *kubernetes.Clientset, namespace string) ([]api.Workload, error){
		workload.GetAllDaemonSet,
		workload.GetAllDeployment,
		workload.GetAllJobs,
		workload.GetAllStatefulset,
	}

	workloads := []api.Workload{}
	fetchGroup, _ := errgroup.WithContext(context.Background())

	for _, f := range allResources {
		fetcherFunction := f
		fetchGroup.Go(func() error {
			lists, err := fetcherFunction(clientset, *configFlags.Namespace)
			if err == nil {
				workloads = append(workloads, lists...)
			}
			return err
		})
	}

	// Wait for all HTTP fetches to complete.
	if err = fetchGroup.Wait(); err != nil {
		err = errors.Wrap(err, "failed to get all pvc in namespaces")
		return
	}

	for _, wk := range workloads {
		if !wk.IsEmpty() {
			removeVolume(volumes, wk)
		} else {
			markVolumeAsZeroReplica(volumes, wk)
		}
	}

	table := uitable.New()
	table.AddRow("Name", "Volume Name", "Size", "Reason", "Used By")

	for _, p := range volumes {
		if p != nil {
			storageSize := p.Spec.Resources.Requests[v1.ResourceStorage]
			table.AddRow(p.Name, p.Spec.VolumeName, storageSize.String(), p.Reason, workload.Join(p.Workloads, ","))
		}
	}
	output = table.String()

	return
}

func removeVolume(volumes []*api.Volume, workload api.Workload) {
	for _, claim := range workload.GetVolumeNames() {
		for idx, vol := range volumes {
			if vol != nil {
				if strings.HasPrefix(vol.GetName(), claim) {
					volumes[idx] = nil
				}
			}
		}
	}
}

func markVolumeAsZeroReplica(volumes []*api.Volume, workload api.Workload) {
	for _, claim := range workload.GetVolumeNames() {
		for idx, vol := range volumes {
			if vol != nil {
				if strings.HasPrefix(vol.GetName(), claim) {
					volumes[idx].Reason = api.WORKLOAD_HAS_ZERO_REPLICAS
					volumes[idx].Workloads = append(volumes[idx].Workloads, workload)
				}
			}
		}
	}
}
